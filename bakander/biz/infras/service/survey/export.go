package survey

import (
	"github.com/pkg/errors"
	surveyquestion2 "kcers-survey/biz/dal/db/ent/surveyquestion"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/ent/surveyresponseanswers"
	service2 "kcers-survey/biz/infras/service"
	"kcers-survey/idl_gen/model/service"
)

func (s Survey) ListResponseExport(req *service.ResponseListReq) (string, error) {
	req.Page = 1
	req.PageSize = 100000
	resp, total, _ := s.ListResponse(req)

	if total == 0 {
		return "", errors.New("暂无数据")
	}

	// 查询该问卷的所有一级问题
	questions, _ := s.db.SurveyQuestion.Query().Where(
		surveyquestion2.SurveyIDEQ(req.SurveyId),
		surveyquestion2.Delete(0),
		surveyquestion2.ParentID(0),
	).Order(surveyquestion2.BySort()).All(s.ctx)

	// 构建表头
	tale := []interface{}{
		"编号", "受访人", "受访人联系电话",
		"调研员", "调研员联系电话", "填写时间", "完成度",
		"省", "市（州）", "县（区、旗）", "乡（镇）", "村",
	}
	questionOrder := make([]int64, 0, len(questions))
	for _, q := range questions {
		tale = append(tale, q.Content)
		questionOrder = append(questionOrder, q.ID)
	}

	// 批量查询所有答案 — 一次性加载，避免 N+1
	responseIDs := make([]int64, 0, len(resp))
	for _, row := range resp {
		responseIDs = append(responseIDs, row.ID)
	}

	allAnswers, _ := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyResponseIDIn(responseIDs...),
		surveyresponseanswers2.Delete(0),
	).All(s.ctx)

	// 按 (responseId, questionId) 建索引
	answerMap := make(map[int64]map[int64]string)
	for _, a := range allAnswers {
		if answerMap[a.SurveyResponseID] == nil {
			answerMap[a.SurveyResponseID] = make(map[int64]string)
		}
		val := a.AnswerText
		if val == "" && len(a.Answer) > 0 {
			for i, part := range a.Answer {
				if i > 0 {
					val += ", "
				}
				val += part
			}
		}
		answerMap[a.SurveyResponseID][a.SurveyQuestionID] = val
	}

	// 组装导出数据
	var list []map[int]interface{}
	for _, row := range resp {
		item := map[int]interface{}{
			1:  row.Sn,
			2:  row.Respondent,
			3:  row.RespondentPhone,
			4:  row.Researcher,
			5:  row.ResearcherPhone,
			6:  row.CreatedAt,
			7:  row.AnswerCount,
			8:  row.Area,
			9:  row.City,
			10: row.District,
			11: row.Village,
			12: row.Address,
		}
		for i, qid := range questionOrder {
			col := 13 + i
			if m, ok := answerMap[row.ID]; ok {
				item[col] = m[qid]
			} else {
				item[col] = ""
			}
		}
		list = append(list, item)
	}

	domain, err := service2.Export(tale, list, "")
	if err != nil {
		return "", err
	}
	return domain, nil
}
