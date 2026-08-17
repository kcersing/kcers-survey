package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"kcers-survey/biz/dal/db/ent"
	"kcers-survey/biz/dal/db/ent/survey"
	surveyquestion "kcers-survey/biz/dal/db/ent/surveyquestion"
	"kcers-survey/idl_gen/model/service"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func openPg(databaseUrl string) *entsql.Driver {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return entsql.OpenDB(dialect.Postgres, db)
}

func main() {
	dsn := "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"

	drv := openPg(dsn)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	ctx := context.Background()
	fmt.Println("Connected to RDS via ent ORM")

	// Create tables via ent auto-migration (fresh DB, no conflicts)
	if err := client.Schema.Create(ctx); err != nil {
		fmt.Println("Schema create failed:", err)
		os.Exit(1)
	}
	fmt.Println("Schema ready")

	// Read survey JSON
	data, err := os.ReadFile("../_survey_data.json")
	if err != nil {
		fmt.Println("Failed to read JSON:", err)
		os.Exit(1)
	}

	var d struct {
		Title    string `json:"title"`
		Sections []struct {
			Code      string `json:"code"`
			Name      string `json:"name"`
			Target    string `json:"target"`
			Questions []struct {
				ID       int    `json:"id"`
				Serial   string `json:"serial"`
				Content  string `json:"content"`
				Type     string `json:"type"`
				Required int    `json:"required"`
				Options  []struct {
					Content string `json:"content"`
					Serial  int    `json:"serial"`
				} `json:"options"`
			} `json:"questions"`
		} `json:"sections"`
	}
	json.Unmarshal(data, &d)

	if d.Title == "" || len(d.Title) < 10 {
		d.Title = "农村养老专项调查问卷"
	}

	groups := []struct{ Code, Suffix, Title string }{
		{"A", "A", d.Title + " — A卷 农村老年人"},
		{"B", "B", d.Title + " — B卷 家庭照护者"},
		{"C", "C", d.Title + " — C卷 村庄"},
	}

	for _, g := range groups {
		qc := 0
		for _, s := range d.Sections {
			if strings.HasPrefix(s.Code, g.Suffix) {
				qc += len(s.Questions)
			}
		}
		if qc == 0 {
			continue
		}

		sv, err := client.Survey.Create().
			SetTitle(g.Title).
			SetCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			fmt.Printf("FAIL survey %s: %v\n", g.Code, err)
			continue
		}
		fmt.Printf("Survey #%d: %s (%d qs)\n", sv.ID, g.Title, qc)

		sortOrder := 0
		for _, sec := range d.Sections {
			if !strings.HasPrefix(sec.Code, g.Suffix) {
				continue
			}
			// Section header as h2
			hdr := fmt.Sprintf("[%s] %s", sec.Code, sec.Name)
			client.SurveyQuestion.Create().
				SetSurveyID(sv.ID).SetParentID(0).SetSerial(sec.Code).
				SetContent(hdr).SetType("h2").SetRequired(1).
				SetSort(int64(sortOrder)).
				Save(ctx)
			sortOrder++

			for i, q := range sec.Questions {
				ops := make([]*service.Options, len(q.Options))
				for j, o := range q.Options {
					ops[j] = &service.Options{Content: o.Content, Serial: int64(o.Serial)}
				}
				client.SurveyQuestion.Create().
					SetSurveyID(sv.ID).SetParentID(0).SetSerial(q.Serial).
					SetContent(q.Content).SetType(q.Type).SetOptions(ops).
					SetRequired(int64(q.Required)).SetSort(int64(sortOrder*100+i)).
					Save(ctx)
			}
			fmt.Printf("  [%s] %d qs\n", sec.Code, len(sec.Questions))
		}
	}

	// Add A2-6/7/8
	addA2(ctx, client)

	// Update B3-3/4
	updateB3(ctx, client)

	// Summary
	fmt.Println("\n====================")
	svs, _ := client.Survey.Query().All(ctx)
	for _, s := range svs {
		cnt, _ := client.SurveyQuestion.Query().
			Where(surveyquestion.SurveyID(s.ID), surveyquestion.TypeNEQ("h2")).
			Count(ctx)
		fmt.Printf("Survey #%d: %s — %d questions\n", s.ID, s.Title[:min(60, len(s.Title))], cnt)
	}
	fmt.Println("Done!")
}

func addA2(ctx context.Context, client *ent.Client) {
	sv, err := client.Survey.Query().Where(survey.TitleContains("A卷")).Order(ent.Desc(survey.FieldID)).First(ctx)
	if err != nil {
		fmt.Println("A卷 not found:", err)
		return
	}

	for oldNum := 12; oldNum >= 6; oldNum-- {
		oldS := fmt.Sprintf("A2-%d", oldNum)
		newS := fmt.Sprintf("A2-%d", oldNum+3)
		client.SurveyQuestion.Update().
			Where(surveyquestion.SurveyID(sv.ID), surveyquestion.Serial(oldS)).
			SetSerial(newS).Save(ctx)
	}

	insertQ(ctx, client, sv.ID, "A2-6", "您是否签约家庭医生？", "single_choice",
		`[{"content":"是","serial":1},{"content":"否（跳转至A2-9）","serial":2},{"content":"不清楚（跳转至A2-9）","serial":3}]`)
	insertQ(ctx, client, sv.ID, "A2-7", "过去一年内，家庭医生是否主动联系过您并开展过相应服务？", "single_choice",
		`[{"content":"是，___次","serial":1},{"content":"否（跳转至A2-9）","serial":2}]`)
	insertQ(ctx, client, sv.ID, "A2-8", "家庭医生为您提供过哪些服务？（可多选）", "multiple_choice",
		`[{"content":"测量血压血糖等基础检查","serial":1},{"content":"慢性病随访","serial":2},{"content":"健康咨询","serial":3},{"content":"康复护理","serial":4},{"content":"用药指导","serial":5},{"content":"转诊服务","serial":6},{"content":"中医药服务","serial":7},{"content":"上门诊疗","serial":8},{"content":"其他","serial":9}]`)
	fmt.Println("  A2-6/7/8 inserted")
}

func updateB3(ctx context.Context, client *ent.Client) {
	sv, err := client.Survey.Query().Where(survey.TitleContains("B卷")).Order(ent.Desc(survey.FieldID)).First(ctx)
	if err != nil {
		fmt.Println("B卷 not found:", err)
		return
	}

	var opts1, opts2 []*service.Options
	json.Unmarshal([]byte(`[{"content":"线上学习","serial":1},{"content":"村里培训","serial":2},{"content":"医院/机构培训","serial":3},{"content":"入户培训","serial":4},{"content":"其他培训","serial":5}]`), &opts1)
	json.Unmarshal([]byte(`[{"content":"基础生活照护","serial":1},{"content":"失能护理技能","serial":2},{"content":"康复训练方法","serial":3},{"content":"慢病管理","serial":4},{"content":"应急处理知识","serial":5},{"content":"营养与膳食管理","serial":6},{"content":"心理与情绪疏导","serial":7},{"content":"用药管理与安全指导","serial":8},{"content":"安宁疗护与临终关怀","serial":9},{"content":"其他","serial":10}]`), &opts2)

	client.SurveyQuestion.Update().
		Where(surveyquestion.SurveyID(sv.ID), surveyquestion.Serial("B3-3")).
		SetContent("您接受的培训类型是？").SetType("single_choice").SetOptions(opts1).Save(ctx)
	client.SurveyQuestion.Update().
		Where(surveyquestion.SurveyID(sv.ID), surveyquestion.Serial("B3-4")).
		SetContent("您接受的培训内容是？（可多选）").SetType("multiple_choice").SetOptions(opts2).Save(ctx)
	fmt.Println("  B3-3/4 updated")
}

func insertQ(ctx context.Context, client *ent.Client, sid int64, serial, content, qtype, optsJSON string) {
	var ops []*service.Options
	json.Unmarshal([]byte(optsJSON), &ops)
	_, err := client.SurveyQuestion.Create().
		SetSurveyID(sid).SetParentID(0).SetSerial(serial).
		SetContent(content).SetType(qtype).SetOptions(ops).
		SetRequired(1).SetSort(0).Save(ctx)
	if err != nil {
		fmt.Printf("  FAIL %s: %v\n", serial, err)
	} else {
		fmt.Printf("  OK %s\n", serial)
	}
}

func min(a, b int) int { if a < b { return a }; return b }
