package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()
	db.Ping()
	fmt.Println("Connected")

	// === Fix B3: replace all B3 records cleanly ===
	// Delete existing B3 records in survey #2
	db.Exec("DELETE FROM survey_question WHERE survey_id=2 AND serial LIKE 'B3-%'")
	fmt.Println("Deleted old B3 records")

	// Re-insert B3 with correct data
	b3Questions := []Question{
		{Serial: "B3-1", Content: "您掌握基本照护知识和技能的程度如何？", Type: "single_choice", Options: []Opt{
			{"非常熟悉", 1}, {"比较熟悉", 2}, {"一般", 3}, {"不太熟悉", 4}, {"完全不熟", 5},
		}},
		{Serial: "B3-2", Content: "您参加过照护知识和技能培训吗？", Type: "single_choice", Options: []Opt{
			{"经常参加", 1}, {"偶尔参加", 2}, {"从未参加（跳转至B3-4）", 3},
		}},
		{Serial: "B3-3", Content: "您接受的培训类型是？", Type: "single_choice", Options: []Opt{
			{"线上学习", 1}, {"村里培训", 2}, {"医院/机构培训", 3}, {"入户培训", 4}, {"其他培训", 5},
		}},
		{Serial: "B3-4", Content: "您接受的培训内容是？（可多选）", Type: "multiple_choice", Options: []Opt{
			{"基础生活照护", 1}, {"失能护理技能", 2}, {"康复训练方法", 3},
			{"慢病管理", 4}, {"应急处理知识", 5}, {"营养与膳食管理", 6},
			{"心理与情绪疏导", 7}, {"用药管理与安全指导", 8},
			{"安宁疗护与临终关怀", 9}, {"其他", 10},
		}},
		{Serial: "B3-5", Content: "当照护过程中遇到困难时，您通常向谁寻求帮助？", Type: "single_choice", Options: []Opt{
			{"亲属", 1}, {"村干部", 2}, {"医生", 3}, {"邻居", 4}, {"朋友", 5}, {"其他", 6},
		}},
		{Serial: "B3-6", Content: "您愿意参加照护技能培训吗？", Type: "single_choice", Options: []Opt{
			{"非常愿意", 1}, {"比较愿意", 2}, {"一般", 3}, {"不太愿意", 4}, {"不愿意", 5},
		}},
	}

	for _, q := range b3Questions {
		opts, _ := json.Marshal(q.Options)
		_, err := db.Exec(
			"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, options, required, sort) VALUES (2, 0, $1, $2, $3, $4, 1, 0)",
			q.Serial, q.Content, q.Type, string(opts),
		)
		if err != nil {
			fmt.Printf("FAIL %s: %v\n", q.Serial, err)
		} else {
			fmt.Printf("OK %s: %s (%s)\n", q.Serial, q.Content[:min(45, len(q.Content))], q.Type)
		}
	}

	fmt.Println("\n--- Final State ---")
	for _, sid := range []int{1, 2, 3} {
		var cnt int
		db.QueryRow("SELECT count(*) FROM survey_question WHERE survey_id=$1 AND type!='h2'", sid).Scan(&cnt)
		fmt.Printf("Survey #%d: %d questions\n", sid, cnt)
	}

	fmt.Println("\nA2 section:")
	printSection(db, 1, "A2")
	fmt.Println("\nB3 section:")
	printSection(db, 2, "B3")
	fmt.Println("Done!")
}

type Question struct {
	Serial  string
	Content string
	Type    string
	Options []Opt
}
type Opt struct {
	Content string
	Serial  int
}

func printSection(db *sql.DB, surveyID int, section string) {
	rows, _ := db.Query(
		"SELECT serial, content, type FROM survey_question WHERE survey_id=$1 AND serial LIKE $2 ORDER BY serial",
		surveyID, section+"-%",
	)
	defer rows.Close()
	for rows.Next() {
		var s, c, t string
		rows.Scan(&s, &c, &t)
		fmt.Printf("  %s: %s (%s)\n", s, c[:min(60, len(c))], t)
	}
}

func min(a, b int) int {
	if a < b { return a }
	return b
}