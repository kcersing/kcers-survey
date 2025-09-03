package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"kcers-survey/idl_gen/model/service"
	"strconv"
	"testing"
)

type Ree struct {
	Title    string
	Id       string
	parentID int64
	Options  []*service.Options
}

func treeToMap(tree []*Tree) []Ree {
	var result []Ree

	for _, item := range tree {
		result = append(result, Ree{
			Title:    item.Title,
			Id:       item.Key,
			parentID: item.parentID,
			Options:  item.Options,
		})
		if item.Children != nil {
			result = append(result, treeToMap(item.Children)...)
		}
	}

	return result
}

type Tree struct {
	Title    string
	Value    string
	Key      string
	Method   string
	parentID int64
	Options  []*service.Options
	Children []*Tree
}

func findTreeQuestionChildren1(data []*ent.SurveyQuestion, parentID int64) []*Tree {
	if data == nil {
		return nil
	}
	var result []*Tree
	for _, v := range data {
		if v.ParentID == parentID && v.ID != parentID {
			var m = new(Tree)
			m.Title = v.Content
			m.Value = strconv.FormatInt(v.ID, 10)
			m.Key = strconv.FormatInt(v.ID, 10)
			m.parentID = v.ParentID
			m.Options = v.Options
			m.Children = findTreeQuestionChildren1(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
func TestSurvey(t *testing.T) {

	dbs := db.InItDB("root:kcer-913639@tcp(101.126.9.226:3306)/survey?charset=utf8mb4&parseTime=True&loc=Local", true)
	ctx := context.Background()

	//areas, err := dbs.Area.Query().All(ctx)
	//if err != nil {
	//	return
	//}
	//areaArr := make(map[int64]string)
	//for _, item := range areas {
	//	areaArr[item.ID] = item.Name
	//
	//}
	//
	//sq, err := dbs.SurveyQuestion.Query().
	//	Where(surveyquestion2.SurveyID(1), surveyquestion2.Delete(0)).
	//	Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
	//	All(ctx)
	//if err != nil {
	//	return
	//}

	//sqArr := make(map[int64]*ent.SurveyQuestion)
	//for _, item := range sq {
	//	sqArr[item.ID] = item
	//
	//}

	//resp := findTreeQuestionChildren1(sq, 0)
	//
	//treeMap := treeToMap(resp)
	//
	//for i, item := range treeMap {
	//	hlog.Info(i)
	//	hlog.Info(item)
	//}
	//var tale []interface{}
	//for _, item := range treeMap {
	//	tale = append(tale, item.Title)
	//}

	sr, err := dbs.SurveyResponse.Query().
		Where(surveyresponse2.SurveyID(1), surveyresponse2.Delete(0),
			surveyresponse2.AnswersCountGTE(100),
			surveyresponse2.Or(surveyresponse2.ResearcherNEQ(""),
				surveyresponse2.ResearcherPhoneNEQ(""),
			),
		).
		Order(ent.Asc(surveyresponse2.FieldID)).
		All(ctx)
	if err != nil {
		return
	}

	for i, item := range sr {

		sra, err := dbs.SurveyResponseAnswers.
			Query().
			Where(surveyresponseanswers2.SurveyResponseID(item.ID), surveyresponseanswers2.Delete(0)).
			Order(ent.Asc(surveyresponse2.FieldID)).
			All(ctx)
		if err != nil {
			return
		}

		hlog.Info(sra)
		panic(i)
	}

	//sra, err := dbs.SurveyResponseAnswers.Query().
	//	Where(surveyresponseanswers2.SurveyID(1), surveyresponseanswers2.Delete(0)).
	//	Order(ent.Asc(surveyresponseanswers2.FieldID)).
	//	All(ctx)
	//if err != nil {
	//	return
	//}

	//
	//tale := []interface{}{
	//	"编号",
	//	"受访人",
	//	"受访人联系电话",
	//	"调研员",
	//	"调研员",
	//	"填写问卷时间",
	//	"完成度",
	//	"省",
	//	"市（州）",
	//	"县（区、旗）",
	//	"乡（镇）",
	//	"村",
	//}
	//var list []map[int]interface{}
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
	//domain, err := service2.Export(tale, list)
	//hlog.Info(err)
	//hlog.Info(domain)

}
