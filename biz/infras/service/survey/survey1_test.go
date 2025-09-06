package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	service2 "kcers-survey/biz/infras/service"
	"testing"
	"time"
)

func SName(id int64, dbs *ent.Client) string {

	sq, _ := dbs.SurveyQuestion.Query().
		Where(
			surveyquestion2.ID(id),
		).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		First(context.Background())

	if sq != nil {
		if sq.ParentID != 0 {
			return SName(sq.ParentID, dbs) + "-" + sq.Content
		} else {
			return sq.Content
		}
	}

	return ""
}

func TestSurvey1(t *testing.T) {

	dbs := db.InItDB("root:kcer-913639@tcp(101.126.9.226:3306)/survey?charset=utf8mb4&parseTime=True&loc=Local", true)

	rd := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})

	ctx := context.Background()
	var name string
	sq, err := dbs.SurveyQuestion.Query().
		Where(
			surveyquestion2.SurveyID(2),
			surveyquestion2.Delete(0),
			surveyquestion2.Type("single_choice"),
		).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		First(ctx)

	if err != nil {
		return
	}

	//sqArr := make(map[int64]*ent.SurveyQuestion)

	//for _, item := range sq {
	//	sqArr[item.ID] = item
	//
	//}
	//
	//resp := findTreeQuestionChildren1(sq, 0)
	//
	//treeMap := treeToMap(resp)

	var tale []interface{}
	var list []map[int]interface{}
	tale = append(tale, "编号")
	tale = append(tale, "受访人")
	tale = append(tale, "受访人联系电话")
	tale = append(tale, "调研员")
	tale = append(tale, "调研员联系电话")
	tale = append(tale, "填写问卷时间")
	tale = append(tale, "完成度")
	tale = append(tale, "省")
	tale = append(tale, "市（州）")
	tale = append(tale, "县（区、旗）")
	tale = append(tale, "乡（镇）")
	tale = append(tale, "详细地址")
	tale = append(tale, "题目")
	//for _, s := range treeMap {
	//	if s.Type == "single_choice" {
	//		for _, o := range s.Options {
	//			tale = append(tale, o.Content)
	//		}
	//	} else if s.Type == "multiple_choice" {
	//		for _, o := range s.Options {
	//			tale = append(tale, o.Content)
	//		}
	//	} else {
	//		tale = append(tale, s.Title)
	//	}
	//
	//}

	for _, o := range sq.Options {
		tale = append(tale, o.Content)
	}
	tale = append(tale, "其他补充")

	sr, err := dbs.SurveyResponse.Query().
		Where(surveyresponse2.SurveyID(2), surveyresponse2.Delete(0),
			//surveyresponse2.AnswersCountGTE(100),
			surveyresponse2.Or(surveyresponse2.ResearcherNEQ(""),
				surveyresponse2.ResearcherPhoneNEQ(""),
			),
		).
		Order(ent.Asc(surveyresponse2.FieldID)).
		All(ctx)
	if err != nil {
		return
	}

	//var datas []*Data
	for _, item := range sr {

		area, _ := rd.Get(ctx, "area"+item.Area).Result()
		city, _ := rd.Get(ctx, "area"+item.City).Result()
		district, _ := rd.Get(ctx, "area"+item.District).Result()
		village, _ := rd.Get(ctx, "area"+item.Village).Result()

		li := map[int]interface{}{
			1:  item.Sn,
			2:  item.Respondent,
			3:  item.RespondentPhone,
			4:  item.Researcher,
			5:  item.ResearcherPhone,
			6:  item.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
			7:  item.AnswersCount,
			8:  area,
			9:  city,
			10: district,
			11: village,
			12: item.Address,
		}
		name = sq.Serial + " " + SName(sq.ParentID, dbs) + "-" + sq.Content
		li[13] = name
		sra, err := dbs.SurveyResponseAnswers.
			Query().
			Where(surveyresponseanswers2.SurveyResponseID(item.ID),
				surveyresponseanswers2.Delete(0),
				surveyresponseanswers2.SurveyQuestionID(sq.ID),
			).
			Order(ent.Asc(surveyresponse2.FieldID)).
			All(ctx)
		if err != nil {
			return
		}

		for ib, b := range sq.Options {

			li[ib+1+13] = ""
			for _, s := range sra {
				for _, s1 := range s.Answer {
					if b.Content == s1 {
						li[ib+1+13] = 1
					} else {
						li[ib+1+13] = 0
					}
				}
				if s.AnswerText != "" {
					li[len(tale)] = s.AnswerText
				}
			}
		}

		//for _, s := range data.Ree {
		//	hlog.Info(s.Question)
		//}
		list = append(list, li)
	}

	//for _, row := range datas {

	//for i, s1 := range treeMap {
	//	for _, s2 := range row.Question {
	//		if s1.Id == s2.Id {
	//			li[i+12] = s2.QuestionContent
	//			//ans := append(s2.Answer, s2.AnswerText)
	//
	//		}
	//
	//	}
	//
	//}
	//list = append(list, li)

	//for _, row := range resp {
	//	list = append(list, map[int]interface{}{
	//		1:  row.Sn,
	//		2:  row.Respondent,
	//		3:  row.RespondentPhone,
	//		4:  row.Researcher,
	//		5:  row.ResearcherPhone,
	//		6:  row.CreatedAt,
	//		7:  row.AnswerCount,
	//		8:  row.Area,
	//		9:  row.City,
	//		10: row.District,
	//		11: row.Village,
	//		12: row.Address,
	//	})
	//
	//}

	domain, err := service2.Export(tale, list, name)
	//hlog.Info(err)
	hlog.Info(domain)
}
