package survey

import (
	"context"
	service2 "kcers-survey/biz/infras/service"
	"math"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"

	"testing"
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
		}
		return sq.Content

	}

	return ""
}

type Basic struct {
	Name string

	Value       int64
	Proportion  float64
	Proportion2 float64
}
type StatisticsBasic struct {
	Count2 int
	Count  int64
	Name   string
	Data   []Basic
}

func TestSurvey1(t *testing.T) {

	dbs := db.InItDB("root:kcer-913639@tcp(101.126.9.226:3306)/survey?charset=utf8mb4&parseTime=True&loc=Local", true)

	//rd := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   1,
	//})

	ctx := context.Background()

	sr, err := dbs.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(1), surveyresponse2.Delete(0),
			surveyresponse2.AnswersCountGTE(100),
			surveyresponse2.Or(surveyresponse2.ResearcherNEQ(""),
				surveyresponse2.ResearcherPhoneNEQ(""),
			),
		).
		Order(ent.Asc(surveyresponse2.FieldID)).
		IDs(ctx)
	if err != nil {
		return
	}

	sqarr, err := dbs.SurveyQuestion.Query().
		Where(
			surveyquestion2.SurveyID(1),
			surveyquestion2.Delete(0),
			surveyquestion2.TypeIn("single_choice", "multiple_choice"),
		).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		All(ctx)

	if err != nil {
		return
	}
	var tale []interface{}
	tale = append(tale, "问题")
	tale = append(tale, "频次")
	tale = append(tale, "合计")
	tale = append(tale, "合计2")

	var list []map[int]interface{}

	for _, sq := range sqarr {

		resp := answerCount1(sq, sr, dbs, ctx)

		hlog.Info(resp)
		li := map[int]interface{}{}
		li[1] = resp.Name
		li[2] = "频次"
		li[3] = "合计"
		li[4] = "合计2"
		list = append(list, li)

		for _, v := range resp.Data {
			li := map[int]interface{}{}
			li[1] = v.Name
			li[2] = v.Value
			li[3] = resp.Count
			li[4] = resp.Count2
			list = append(list, li)
		}

		list = append(list, map[int]interface{}{
			1: "",
		})

	}

	service2.Export(tale, list, "")
}

func answerCount1(sq *ent.SurveyQuestion, ids []int64, db *ent.Client, ctx context.Context) (resp *StatisticsBasic) {
	count, err := db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(sq.ID),
		surveyresponseanswers2.Delete(0),
		surveyresponseanswers2.SurveyResponseIDIn(ids...),
	).Count(ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}

	var bas []Basic
	for _, o := range sq.Options {
		bas = append(bas, Basic{
			Name:  o.Content,
			Value: 0,
		})
	}

	all, err := db.SurveyResponseAnswers.
		Query().
		Where(
			surveyresponseanswers2.SurveyQuestionID(sq.ID),
			surveyresponseanswers2.Delete(0),
			surveyresponseanswers2.SurveyResponseIDIn(ids...),
		).All(ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	var str []string
	for _, v := range all {

		if len(v.Answer) > 0 {
			for _, an := range v.Answer {

				str = append(str, an)

			}
		}
	}
	resp = &StatisticsBasic{
		Count:  int64(count),
		Count2: len(str),
		Name:   SName(sq.ParentID, db) + sq.Content,
	}
	for _, o := range str {
		for i, b := range bas {
			if o == b.Name {
				bas[i].Value = bas[i].Value + 1
			}
		}
	}

	resp.Data = bas
	for i, v := range resp.Data {
		resp.Data[i].Proportion = math.Round(float64(v.Value)/float64(resp.Count)*10000) / 100
		resp.Data[i].Proportion2 = math.Round(float64(v.Value)/float64(resp.Count2)*10000) / 100
	}
	return
}
