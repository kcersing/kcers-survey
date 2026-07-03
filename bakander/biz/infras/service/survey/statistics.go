package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/dal/db/ent/predicate"
	surveyquestion2 "kcers-survey/biz/dal/db/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/ent/surveyresponseanswers"
	"kcers-survey/idl_gen/model/service"
)

func (s Survey) GetQuestionStatisticsBasic(id int64) (resp *service.StatisticsBasic, err error) {

	first, err := s.db.SurveyQuestion.Query().Where(
		surveyquestion2.IDEQ(id),
	).First(s.ctx)
	if err != nil {
		return nil, err
	}

	switch first.Type {
	case "multiple_choice":
		resp = s.answerCountMulti(id)
	case "matrix":
		resp = s.answerCountMatrix(id)
	case "ranking":
		resp = s.answerCountRanking(id)
	default:
		resp = s.answerCount(id)
	}

	return resp, err
}

type ScAll struct {
	Count  int64    `json:"count"`
	Answer []string `json:"answer"`
}

func (s Survey) answerCountMulti(id int64) (resp *service.StatisticsBasic) {
	auestion, err := s.db.SurveyQuestion.Query().Where(surveyquestion2.IDEQ(id)).First(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	var bas []*service.Basic
	for _, o := range auestion.Options {
		bas = append(bas, &service.Basic{
			Type:  o.Content,
			Value: 1,
		})
	}
	all, err := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).All(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	for _, v := range all {
		if len(v.Answer) > 0 {
			for _, an := range v.Answer {
				for _, o := range bas {
					if an == o.Type {
						o.Value = o.Value + 1
					}
				}
			}
		}
	}

	count, err := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).Count(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	resp = &service.StatisticsBasic{
		Count:      int64(count),
		Data:       bas,
		QuestionId: id,
	}
	return
}

// answerCountMatrix 矩阵题统计 — 按子问题 x 选项组合统计
func (s Survey) answerCountMatrix(id int64) (resp *service.StatisticsBasic) {
	question, err := s.db.SurveyQuestion.Query().Where(surveyquestion2.IDEQ(id)).First(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}

	// 查询子问题（行）
	children, _ := s.db.SurveyQuestion.Query().Where(surveyquestion2.ParentID(id)).All(s.ctx)

	basMap := make(map[string]*service.Basic)
	for _, row := range children {
		for _, col := range question.Options {
			key := row.Content + " - " + col.Content
			basMap[key] = &service.Basic{Type: key, Value: 0}
		}
	}

	answers, _ := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).All(s.ctx)

	for _, a := range answers {
		for _, ansVal := range a.Answer {
			if b, ok := basMap[ansVal]; ok {
				b.Value++
			}
		}
	}

	var bas []*service.Basic
	for _, b := range basMap {
		bas = append(bas, b)
	}

	count, _ := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).Count(s.ctx)

	return &service.StatisticsBasic{Count: int64(count), Data: bas, QuestionId: id}
}

// answerCountRanking 排序题统计 — 统计每个选项被选中的总频次
func (s Survey) answerCountRanking(id int64) (resp *service.StatisticsBasic) {
	question, err := s.db.SurveyQuestion.Query().Where(surveyquestion2.IDEQ(id)).First(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}

	basMap := make(map[string]*service.Basic)
	var bas []*service.Basic
	for _, o := range question.Options {
		b := &service.Basic{Type: o.Content, Value: 0}
		bas = append(bas, b)
		basMap[o.Content] = b
	}

	answers, _ := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).All(s.ctx)

	for _, a := range answers {
		for _, ansVal := range a.Answer {
			// 排序答案格式: "A > B > C"
			parts := splitRanking(ansVal)
			for _, part := range parts {
				if b, ok := basMap[part]; ok {
					b.Value++
				}
			}
		}
	}

	count, _ := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).Count(s.ctx)

	return &service.StatisticsBasic{Count: int64(count), Data: bas, QuestionId: id}
}

func splitRanking(s string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '>' {
			result = append(result, trimSpace(s[start:i]))
			start = i + 1
		}
	}
	result = append(result, trimSpace(s[start:]))
	return result
}

func trimSpace(s string) string {
	start, end := 0, len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

func (s Survey) answerCount(id int64) (resp *service.StatisticsBasic) {
	var scAll []*ScAll
	err := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(sql.Count("*"), "count"),
				sql.As(surveyresponseanswers2.FieldAnswer, "answer"),
			).
				GroupBy(
					surveyresponseanswers2.FieldSurveyQuestionID,
					surveyresponseanswers2.FieldAnswer,
				)
		}).
		Scan(context.Background(), &scAll)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	var bas []*service.Basic
	for _, v := range scAll {
		if len(v.Answer) > 0 {
			bas = append(bas, &service.Basic{
				Type:  v.Answer[0],
				Value: v.Count,
			})
		}

	}

	count, err := s.db.SurveyResponseAnswers.Query().Where(
		surveyresponseanswers2.SurveyQuestionID(id),
		surveyresponseanswers2.Delete(0),
	).Count(s.ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}

	resp = &service.StatisticsBasic{
		Count:      int64(count),
		Data:       bas,
		QuestionId: id,
	}
	return resp
}

func (s Survey) GetSurveyResponseHeatmap(id int64) (resp []*service.Heatmap, err error) {

	err = s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyID(id),
		surveyresponse2.RespondentNEQ(""),
		surveyresponse2.ResearcherNEQ(""),
		surveyresponse2.LatitudeNEQ(""),
		surveyresponse2.Delete(0),
	).
		Modify(func(s *sql.Selector) {

			s.Select(
				sql.As("SUBSTRING( latitude, 1, LOCATE( '.', latitude )+ 3 )", "lat"),
				sql.As("SUBSTRING( longitude, 1, LOCATE( '.', longitude )+ 3 )", "lng"),
				sql.As(sql.Count("*"), "count"),
			).
				GroupBy(
					"SUBSTRING( latitude, 1, LOCATE( '.', latitude )+ 3 )",
					"SUBSTRING( longitude, 1, LOCATE( '.', longitude )+ 3 )",
				)
		}).
		Scan(context.Background(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s Survey) GetSurveyStatistics(id int64) (resp *service.SurveyStatistics, err error) {
	resp = &service.SurveyStatistics{
		Count:           0,
		RespondentCount: 0,
		ResearcherCount: 0,
		VillageCount:    0,
		AnswersCount:    0,
		AnswersAverage:  0,
	}

	var predicates []predicate.SurveyResponse
	predicates = append(predicates, surveyresponse2.SurveyIDEQ(id))
	predicates = append(predicates, surveyresponse2.Delete(0))
	predicates = append(predicates, surveyresponse2.RespondentNEQ(""))
	count, err := s.db.SurveyResponse.Query().Where(predicates...).Count(s.ctx)
	if err != nil {
		hlog.Error(err)
	}
	resp.Count = int64(count)
	respondents, err := s.db.SurveyResponse.Query().Where(predicates...).GroupBy(surveyresponse2.FieldRespondent).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
	}
	resp.RespondentCount = int64(len(respondents))
	researchers, err := s.db.SurveyResponse.Query().Where(predicates...).GroupBy(surveyresponse2.FieldResearcher).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
	}
	resp.ResearcherCount = int64(len(researchers))
	villages, err := s.db.SurveyResponse.Query().Where(predicates...).GroupBy(surveyresponse2.FieldVillage).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
	}
	resp.VillageCount = int64(len(villages))

	answers, err := s.db.SurveyResponse.Query().Where(predicates...).QueryAnswers().Count(s.ctx)
	if err != nil {
		hlog.Error(err)
	}
	resp.AnswersCount = int64(answers)

	resp.AnswersAverage = int64(float64(answers) / float64(count))
	return resp, nil
}
