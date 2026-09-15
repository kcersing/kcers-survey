package survey

import (
	"context"
	"strconv"
	"strings"
	"testing"
	"time"

	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/db/ent"
	"kcers-survey/biz/dal/db/ent/predicate"
	surveyquestion2 "kcers-survey/biz/dal/db/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/ent/surveyresponseanswers"
	service2 "kcers-survey/biz/infras/service"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// surveyScope 一次导出任务:某份问卷 + 是否只取"做完的"答卷。
type surveyScope struct {
	SurveyID   int64
	MinAnswers int64  // 0 = 全量;>0 = 只取 AnswersCount >= MinAnswers 的答卷
	Name       string // 导出文件名前缀,如 "问卷1-已完成"
}

// surveyScopes 6 种导出:问卷 1/2/3 × 全量 / 已完成(阈值沿用原来的注释:50/20/30)。
// survey5_test.go(答卷×题目明细)与 survey3_test.go(问题×省份交叉统计)共用。
var surveyScopes = []surveyScope{
	{SurveyID: 1, MinAnswers: 0, Name: "问卷1-全量"},
	{SurveyID: 1, MinAnswers: 50, Name: "问卷1-已完成"},
	{SurveyID: 2, MinAnswers: 0, Name: "问卷2-全量"},
	{SurveyID: 2, MinAnswers: 20, Name: "问卷2-已完成"},
	{SurveyID: 3, MinAnswers: 0, Name: "问卷3-全量"},
	{SurveyID: 3, MinAnswers: 30, Name: "问卷3-已完成"},
}

// surveyResponsePredicates 答卷筛选条件:问卷 + 未删除 + 有调研员信息,
// MinAnswers > 0 时再叠加"完成度"阈值。
func surveyResponsePredicates(sc surveyScope) []predicate.SurveyResponse {
	ps := []predicate.SurveyResponse{
		surveyresponse2.SurveyID(sc.SurveyID),
		surveyresponse2.Delete(0),
		surveyresponse2.Or(surveyresponse2.ResearcherNEQ(""),
			surveyresponse2.ResearcherPhoneNEQ("")),
	}
	if sc.MinAnswers > 0 {
		ps = append(ps, surveyresponse2.AnswersCountGTE(sc.MinAnswers))
	}
	return ps
}

// TestSurvey5 导出问卷问题详情(答卷 × 题目 明细),一次运行导出 6 种范围:
// 问卷 1/2/3 × 全量 / 已完成。
// 复用 survey_test.go 中已定义的 Ree / Tree / Data / Question 类型与
// treeToMap / findTreeQuestionChildren1 函数(同包内不可重复声明)。
// 相比 TestSurvey，额外把 survey_response 的省/市/县/乡/详细地址各占一列。
func TestSurvey5(t *testing.T) {

	dbs := db.InItDB("user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey port=5432 sslmode=disable TimeZone=Asia/Shanghai", true)

	ctx := context.Background()

	for _, sc := range surveyScopes {
		if err := exportSurveyDetail(dbs, ctx, sc); err != nil {
			t.Errorf("%s 导出失败: %v", sc.Name, err)
		}
	}
}

// exportSurveyDetail 导出某份问卷的"答卷 × 题目"明细表
func exportSurveyDetail(dbs *ent.Client, ctx context.Context, sc surveyScope) error {
	sq, err := dbs.SurveyQuestion.Query().
		Where(surveyquestion2.SurveyID(sc.SurveyID), surveyquestion2.Delete(0)).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		All(ctx)
	if err != nil {
		return err
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
	tale = append(tale, "省")
	tale = append(tale, "市（州）")
	tale = append(tale, "县（区、旗）")
	tale = append(tale, "乡（镇）")
	tale = append(tale, "详细地址")
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
		mun[it] = i
	}

	sr, err := dbs.SurveyResponse.Query().
		Where(surveyResponsePredicates(sc)...).
		Order(ent.Asc(surveyresponse2.FieldID)).
		All(ctx)
	if err != nil {
		return err
	}

	for _, item := range sr {

		li := map[int]interface{}{}

		li = map[int]interface{}{
			1:  item.Sn,
			2:  item.Respondent,
			3:  item.RespondentPhone,
			4:  item.Researcher,
			5:  item.ResearcherPhone,
			6:  item.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
			7:  item.AnswersCount,
			8:  item.Area,
			9:  item.City,
			10: item.District,
			11: item.Village,
			12: item.Address,
		}
		sra, err := dbs.SurveyResponseAnswers.
			Query().
			Where(
				surveyresponseanswers2.SurveyResponseID(item.ID),
				surveyresponseanswers2.Delete(0),
			).
			Order(ent.Asc(surveyresponseanswers2.FieldID)).
			All(ctx)
		if err != nil {
			return err
		}

		for _, s := range sra {
			// 答案指向的题目可能已删除或不属于本问卷,跳过,避免整个导出 panic
			q, ok := sqArr[s.SurveyQuestionID]
			if !ok {
				continue
			}
			key := strconv.FormatInt(s.SurveyQuestionID, 10) + "-" + q.Content
			idx, ok := mun[key].(int)
			if !ok {
				continue
			}
			bian := idx + 1

			ans := append(s.Answer, s.AnswerText)
			li[bian] = strings.Join(ans, " ")

			if q.Type == "multiple_choice" {
				li[bian] = strings.Join(ans, ",")
				for _, b := range q.Options {
					idx1, ok := mun[key+"-"+b.Content].(int)
					if !ok {
						continue
					}
					bian1 := idx1 + 1
					li[bian1] = 0

					for _, s1 := range s.Answer {
						if b.Content == s1 {
							li[bian1] = 1
						}
					}
					if s.AnswerText != "" {
						if idx2, ok := mun[key+"-其他补充"].(int); ok {
							li[idx2+1] = s.AnswerText
						}
					}

				}

			}

		}

		list = append(list, li)

	}

	domain, err := service2.Export(tale, list, sc.Name)
	if err != nil {
		return err
	}
	hlog.Infof("%s 导出成功: %s", sc.Name, domain)
	return nil
}
