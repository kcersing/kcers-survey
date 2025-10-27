package survey

import (
	"context"
	service2 "kcers-survey/biz/infras/service"
	"math"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"

	"testing"
)

func TestSurvey2(t *testing.T) {

	dbs := db.InItDB("root:kcer-913639@tcp(101.126.9.226:3306)/survey?charset=utf8mb4&parseTime=True&loc=Local", true)

	//rd := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   1,
	//})

	ctx := context.Background()

	sr, err := dbs.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(3), surveyresponse2.Delete(0),
			surveyresponse2.AnswersCountGTE(60),
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
			surveyquestion2.SurveyID(3),
			surveyquestion2.Delete(0),
			surveyquestion2.TypeIn("rate"),
		).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		All(ctx)
	hlog.Info(sqarr)

	if err != nil {
		return
	}
	var tale []interface{}
	tale = append(tale, "问题")
	tale = append(tale, "频次")
	tale = append(tale, "回答总数")
	tale = append(tale, "选项总数")

	var list []map[int]interface{}

	for _, sq := range sqarr {

		resp := answerCount2(sq, sr, dbs, ctx)

		li := map[int]interface{}{}
		li[1] = resp.Name
		li[2] = "频次"
		li[3] = "合计"
		li[4] = "合计2"
		list = append(list, li)
		for _, v := range resp.Data {
			li1 := map[int]interface{}{}
			li1[1] = v.Name
			li1[2] = v.Value
			li1[3] = resp.Count
			li1[4] = resp.Count2
			list = append(list, li1)
		}

		list = append(list, map[int]interface{}{
			1: "",
		})

	}

	service2.Export(tale, list, "")
}

func answerCount2(sq *ent.SurveyQuestion, ids []int64, db *ent.Client, ctx context.Context) (resp *StatisticsBasic) {
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
	for _, o := range []int64{0, 1, 2, 3, 4, 5, 6} {
		bas = append(bas, Basic{
			Name:  strconv.FormatInt(o, 10),
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
		hlog.Info(v.Answer)
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
