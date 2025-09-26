package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	service2 "kcers-survey/biz/infras/service"
	"kcers-survey/idl_gen/model/service"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Ree struct {
	Mu       int
	Title    string
	Id       string
	parentID int64
	Type     string
	Options  []*service.Options
	Serial   string
	Question Question
}

func treeToMap(tree []*Tree) []Ree {
	var result []Ree

	for _, item := range tree {
		result = append(result, Ree{
			Title:    item.Title,
			Id:       item.Key,
			parentID: item.parentID,
			Options:  item.Options,
			Serial:   item.Serial,
			Type:     item.Type,
		})
		if item.Children != nil {
			result = append(result, treeToMap(item.Children)...)
		}
	}
	var result2 []Ree
	for i, r := range result {
		r.Mu = i
		result2 = append(result2, r)
	}

	return result2
}

type Tree struct {
	Title    string
	Value    string
	Key      string
	Method   string
	Type     string
	parentID int64
	Options  []*service.Options
	Children []*Tree
	Serial   string
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
			m.Serial = v.Serial
			m.Type = v.Type
			result = append(result, m)
		}
	}
	return result
}

type Data struct {
	Sn       string
	Area     string
	City     string
	District string
	Village  string
	//Question        []*Question
	Address         string
	Respondent      string
	RespondentPhone string
	Researcher      string
	ResearcherPhone string
	At              string
	AnswersCount    int64
	Ree             []Ree
}
type Question struct {
	Id              string
	Serial          string
	QuestionContent string
	Options         []*service.Options
	Answer          []string
	AnswerText      string
	Type            string
}

func TestSurvey(t *testing.T) {

	dbs := db.InItDB("root:kcer-913639@tcp(101.126.9.226:3306)/survey?charset=utf8mb4&parseTime=True&loc=Local", true)

	//rd := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   1,
	//})

	ctx := context.Background()

	//areas, err := dbs.Area.Query().All(ctx)
	//if err != nil {
	//	return
	//}
	////areaArr := make(map[string]string)
	//
	//for _, item := range areas {
	//	//areaArr[strconv.FormatInt(item.ID, 10)] = item.Name
	//	//tx.Append(ctx, "area"+strconv.FormatInt(item.ID, 10), item.Name)
	//	//tx.Exec(ctx)
	//	rd.Del(ctx, "area"+strconv.FormatInt(item.ID, 10))
	//	rd.Set(ctx, "area"+strconv.FormatInt(item.ID, 10), item.Name, 0)
	//	result, err := rd.Get(ctx, "area"+strconv.FormatInt(item.ID, 10)).Result()
	//	if err != nil {
	//		return
	//	}
	//
	//	hlog.Info(result)
	//
	//}
	//print("")
	//
	sq, err := dbs.SurveyQuestion.Query().
		Where(surveyquestion2.SurveyID(3), surveyquestion2.Delete(0)).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		All(ctx)
	if err != nil {
		return
	}

	sqArr := make(map[int64]*ent.SurveyQuestion)
	for _, item := range sq {
		sqArr[item.ID] = item

	}

	resp := findTreeQuestionChildren1(sq, 0)

	treeMap := treeToMap(resp)

	var tale []interface{}
	var list []map[int]interface{}

	tale = append(tale, "编号")
	tale = append(tale, "受访人")
	tale = append(tale, "受访人联系电话")
	tale = append(tale, "调研员")
	tale = append(tale, "调研员联系电话")
	tale = append(tale, "填写问卷时间")
	tale = append(tale, "完成度")
	//tale = append(tale, "省")
	//tale = append(tale, "市（州）")
	//tale = append(tale, "县（区、旗）")
	//tale = append(tale, "乡（镇）")
	//tale = append(tale, "详细地址")
	for _, s := range treeMap {
		tale = append(tale, s.Id+"-"+s.Title)
		if s.Type == "multiple_choice" {
			for _, o := range s.Options {
				tale = append(tale, s.Id+"-"+s.Title+"-"+o.Content)
			}
			tale = append(tale, s.Id+"-"+s.Title+"-其他补充")
		}
	}

	mun := make(map[string]interface{})

	for i, item := range tale {
		it := item.(string)
		//hlog.Info(item)
		mun[it] = i
	}

	sr, err := dbs.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(3), surveyresponse2.Delete(0),
			surveyresponse2.AnswersCountGTE(60),
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

		//area, _ := rd.Get(ctx, "area"+item.Area).Result()
		//city, _ := rd.Get(ctx, "area"+item.City).Result()
		//district, _ := rd.Get(ctx, "area"+item.District).Result()
		//village, _ := rd.Get(ctx, "area"+item.Village).Result()
		li := map[int]interface{}{}

		li = map[int]interface{}{
			1: item.Sn,
			2: item.Respondent,
			3: item.RespondentPhone,
			4: item.Researcher,
			5: item.ResearcherPhone,
			6: item.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
			7: item.AnswersCount,
			//8:  area,
			//9:  city,
			//10: district,
			//11: village,
			//8: item.Address,
		}
		sra, err := dbs.SurveyResponseAnswers.
			Query().
			Where(
				surveyresponseanswers2.SurveyResponseID(item.ID),
				surveyresponseanswers2.Delete(0),
			).
			Order(ent.Asc(surveyresponse2.FieldID)).
			All(ctx)
		if err != nil {
			return
		}

		//for ib, b := range treeMap {

		for _, s := range sra {
			//if b.Id == strconv.FormatInt(s.SurveyQuestionID, 10) {
			bian := mun[strconv.FormatInt(s.SurveyQuestionID, 10)+"-"+sqArr[s.SurveyQuestionID].Content].(int) + 1

			ans := append(s.Answer, s.AnswerText)
			li[bian] = strings.Join(ans, " ")

			if sqArr[s.SurveyQuestionID].Type == "multiple_choice" {
				li[bian] = strings.Join(ans, ",")
				for _, b := range sqArr[s.SurveyQuestionID].Options {

					bian1 := mun[strconv.FormatInt(s.SurveyQuestionID, 10)+"-"+sqArr[s.SurveyQuestionID].Content+"-"+b.Content].(int) + 1
					li[bian1] = 0

					for _, s1 := range s.Answer {
						if b.Content == s1 {
							li[bian1] = 1
						}
					}
					if s.AnswerText != "" {
						bian2 := mun[strconv.FormatInt(s.SurveyQuestionID, 10)+"-"+sqArr[s.SurveyQuestionID].Content+"-其他补充"].(int) + 1

						li[bian2] = s.AnswerText
					}

				}

			}

		}

		list = append(list, li)

	}

	domain, err := service2.Export(tale, list, "")
	//hlog.Info(err)
	hlog.Info(domain)
}
