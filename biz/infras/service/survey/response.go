package survey

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"kcers-survey/idl_gen/model/service"
	"strconv"
	"strings"
	"time"
)

func (s Survey) GetNext(req *service.GetNextReq) (number int64, err error) {

	if req.Sn == "" {
		return 0, nil
	}
	first, err := s.db.SurveyResponse.
		Query().
		Where(surveyresponse2.Sn(req.Sn)).
		First(s.ctx)
	if err != nil {
		return 0, err
	}
	all, err := s.db.SurveyQuestion.Query().
		Where(surveyquestion2.SurveyID(first.SurveyID), surveyquestion2.Level(2)).
		IDs(s.ctx)
	if err != nil {
		return 0, err
	}
	answers, err := s.db.SurveyResponseAnswers.Query().
		Where(surveyresponseanswers2.SurveyResponseID(first.ID)).
		Order(ent.Desc(surveyresponseanswers2.FieldID)).
		First(s.ctx)
	if err != nil {
		return 0, err
	}
	question, err := s.db.SurveyQuestion.Query().
		Where(surveyquestion2.IDEQ(answers.SurveyQuestionID)).
		First(s.ctx)
	if err != nil {
		return 0, err
	}
	hlog.Info(question)

	hlog.Info(all)

	if question.Level == 2 {
		for i, v := range all {
			if v == question.ID {
				return int64(i), err
			}
		}
	} else {
		split := strings.Split(question.Tree, " ")
		hlog.Info(split)
		trimmed := strings.TrimPrefix(split[1], "tr_")
		hlog.Info(trimmed)
		id, err := strconv.ParseInt(trimmed, 10, 64)
		hlog.Info(id)
		hlog.Info(all)
		if err != nil {
			return 0, err
		}
		for i, v := range all {
			if v == id {
				return int64(i), err
			}

		}
	}

	return 0, nil
}

func (s Survey) CreateResponse(req *service.CreateOrUpdateResponseReq) (err error) {
	sa, _ := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(req.SurveyId),
			surveyresponse2.Sn(req.Sn),
		).
		First(s.ctx)
	if sa == nil {
		sa, err = s.db.SurveyResponse.Create().
			SetSurveyID(req.SurveyId).
			SetSn(req.Sn).
			SetIP(s.c.ClientIP()).
			SetDevice(string(s.c.Request.Header.UserAgent())).
			Save(s.ctx)
		if err != nil {
			return err
		}
	}
	sau := sa.Update()
	if req.Type == "location" {
		sau.SetLatitude(req.Latitude)
		sau.SetLongitude(req.Longitude)
	}
	if req.Type == "respondent" {
		sau.SetRespondent(req.Value[0])
	}
	if req.Type == "respondentPhone" {
		sau.SetRespondentPhone(req.Value[0])
	}
	if req.Type == "researcher" {
		sau.SetResearcher(req.Value[0])
	}
	if req.Type == "researcherPhone" {
		sau.SetResearcherPhone(req.Value[0])
	}

	if req.Type == "researcherPhone" {
		sau.SetResearcherPhone(req.Value[0])
	}

	if req.Type == "audio" {
		sau.AppendAudio(req.Value)
	}

	if req.Type == "image" {
		sau.AppendPic(req.Value)
	}
	if req.Type == "area" {
		sau.SetArea(req.Value[0])
	}
	if req.Type == "city" {
		sau.SetCity(req.Value[0])
	}
	if req.Type == "district" {
		sau.SetDistrict(req.Value[0])
	}
	if req.Type == "village" {
		sau.SetVillage(req.Value[0])
	}
	if req.Type == "address" {
		sau.SetAddress(req.Value[0])
	}

	if req.QuestionId > 0 {
		ra, _ := s.db.SurveyResponseAnswers.Query().
			Where(
				surveyresponseanswers2.SurveyID(req.SurveyId),
				surveyresponseanswers2.SurveyResponseID(sa.ID),
				surveyresponseanswers2.SurveyQuestionID(req.QuestionId),
			).
			First(s.ctx)
		if ra == nil {
			ra, err = s.db.SurveyResponseAnswers.Create().
				SetSurveyID(req.SurveyId).
				SetSurveyResponseID(sa.ID).
				SetSurveyQuestionID(req.QuestionId).
				Save(s.ctx)
			if err != nil {
				return err
			}
		}
		rau := ra.Update()

		if req.Type == "input" {

			rau.SetAnswerText(req.Value[0])

		} else {
			rau.SetAnswer(req.Value)
		}

		_, err = rau.Save(s.ctx)
		if err != nil {
			return err
		}
	}

	_, err = sau.Save(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s Survey) UpdateResponse(req *service.CreateOrUpdateResponseReq) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s Survey) GetResponse(id int64) (resp *service.Response, err error) {
	first, err := s.db.SurveyResponse.
		Query().
		Where(surveyresponse2.IDEQ(id)).
		First(s.ctx)
	if err != nil {
		return nil, err
	}
	hlog.Info(first)
	return
}

func (s Survey) ListResponse(req *service.ResponseListReq) (resp []*service.Response, total int, err error) {
	var predicates []predicate.SurveyResponse

	if req.Respondent != "" {
		predicates = append(predicates, surveyresponse2.Respondent(req.Respondent))
	}
	if req.RespondentPhone != "" {
		predicates = append(predicates, surveyresponse2.RespondentPhone(req.RespondentPhone))
	}
	if req.Researcher != "" {
		predicates = append(predicates, surveyresponse2.Researcher(req.Researcher))
	}
	if req.ResearcherPhone != "" {
		predicates = append(predicates, surveyresponse2.ResearcherPhone(req.ResearcherPhone))
	}
	if req.Sn != "" {
		predicates = append(predicates, surveyresponse2.Sn(req.Sn))
	}
	if req.SurveyId > 0 {
		predicates = append(predicates, surveyresponse2.SurveyID(req.SurveyId))
	}
	predicates = append(predicates, surveyresponse2.SnNEQ(""))
	predicates = append(predicates, surveyresponse2.AnswersCountNEQ(0))
	predicates = append(predicates, surveyresponse2.Delete(0))

	all, err := s.db.SurveyResponse.
		Query().
		Where(predicates...).
		Order(ent.Desc(surveyresponse2.FieldAnswersCount)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		All(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	for _, v := range all {
		qCount, _ := s.db.SurveyQuestion.Query().
			Where(surveyquestion2.SurveyID(v.SurveyID), surveyquestion2.DeleteEQ(0)).
			Count(s.ctx)
		resp = append(resp, s.entToResponse(v, qCount))
	}

	total, err = s.db.SurveyResponse.Query().Where(predicates...).Count(s.ctx)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil

}

func (s Survey) entToResponse(v *ent.SurveyResponse, qCount int) *service.Response {

	return &service.Response{
		ID:              v.ID,
		SurveyId:        v.SurveyID,
		Sn:              v.Sn,
		Latitude:        v.Latitude,
		Longitude:       v.Longitude,
		Respondent:      v.Respondent,
		RespondentPhone: v.RespondentPhone,
		Researcher:      v.Researcher,
		ResearcherPhone: v.ResearcherPhone,
		CreatedAt:       v.CreatedAt.Format(time.DateTime),
		Pic:             v.Pic,
		//Audio:           nil,
		IP:          v.IP,
		Device:      v.Device,
		Area:        "",
		City:        "",
		District:    "",
		Address:     "",
		Village:     "",
		AnswerCount: v.AnswersCount,
	}

}

func (s Survey) DeleteResponse(id int64) (err error) {
	_, err = s.db.SurveyResponse.Update().
		Where(surveyresponse2.IDEQ(id)).
		SetDelete(1).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
}
