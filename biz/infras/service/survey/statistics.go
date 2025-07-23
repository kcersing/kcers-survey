package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"kcers-survey/idl_gen/model/service"
)

func (s Survey) GetQuestionStatisticsBasic(id int64) (resp *service.StatisticsBasic, err error) {

	first, err := s.db.SurveyQuestion.Query().Where(
		surveyquestion2.IDEQ(id),
	).First(s.ctx)
	if err != nil {
		return nil, err
	}

	if first.Type == "multiple_choice" {
		resp = s.answerCount1(id)
	} else {
		resp = s.answerCount(id)
	}

	return resp, err
}

type ScAll struct {
	Count  int64    `json:"count"`
	Answer []string `json:"answer"`
}

func (s Survey) answerCount1(id int64) (resp *service.StatisticsBasic) {
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
	count, err := s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyIDEQ(id),
		surveyresponse2.Delete(0),
		surveyresponse2.RespondentNEQ(""),
	).Count(s.ctx)
	if err != nil {
		hlog.Error(err)
		resp.Count = 0
	}
	resp.Count = int64(count)

	respondents, err := s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyIDEQ(id),
		surveyresponse2.Delete(0),
		surveyresponse2.RespondentNEQ(""),
	).GroupBy(surveyresponse2.FieldRespondent).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
		resp.RespondentCount = 0
	}
	resp.RespondentCount = int64(len(respondents))
	researchers, err := s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyIDEQ(id),
		surveyresponse2.Delete(0),
		surveyresponse2.RespondentNEQ(""),
	).GroupBy(surveyresponse2.FieldResearcher).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
		resp.ResearcherCount = 0
	}
	resp.ResearcherCount = int64(len(researchers))
	villages, err := s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyIDEQ(id),
		surveyresponse2.Delete(0),
		surveyresponse2.RespondentNEQ(""),
	).GroupBy(surveyresponse2.FieldVillage).Strings(s.ctx)
	if err != nil {
		hlog.Error(err)
		resp.VillageCount = 0
	}
	resp.VillageCount = int64(len(villages))

	answers, err := s.db.SurveyResponse.Query().Where(
		surveyresponse2.SurveyIDEQ(id),
		surveyresponse2.Delete(0),
		surveyresponse2.RespondentNEQ(""),
	).QueryAnswers().Count(s.ctx)
	if err != nil {
		hlog.Error(err)
		resp.AnswersCount = 0
	}
	resp.AnswersCount = int64(answers)

	return resp, nil
}
