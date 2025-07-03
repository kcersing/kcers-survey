package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	survey2 "kcers-survey/biz/dal/db/mysql/ent/survey"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"kcers-survey/biz/infras/do"
	"kcers-survey/biz/infras/service/common"
	"kcers-survey/biz/pkg/utils"
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/service"
	"strconv"
	"strings"
	"time"
)

type Survey struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Survey) CreateSurvey(req *service.CreateOrUpdateSurveyReq) (err error) {
	startAt, err := utils.GetStringDateTime(req.StartAt)
	if err != nil {
		return err
	}
	endAt, err := utils.GetStringDateTime(req.EndAt)
	if err != nil {
		return err
	}
	_, err = s.db.Survey.Create().
		SetTitle(req.Title).
		SetPic(req.Pic).
		SetDesc(req.Desc).
		SetCreatedID(common.GetTokenUserID(s.c)).
		SetStartAt(startAt).
		SetEndAt(endAt).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s Survey) UpdateSurvey(req *service.CreateOrUpdateSurveyReq) (err error) {
	startAt, err := utils.GetStringDateTime(req.StartAt)
	if err != nil {
		return err
	}
	endAt, err := utils.GetStringDateTime(req.EndAt)
	if err != nil {
		return err
	}
	_, err = s.db.Survey.Update().
		Where(survey2.IDEQ(req.ID)).
		SetTitle(req.Title).
		SetPic(req.Pic).
		SetDesc(req.Desc).
		SetStartAt(startAt).
		SetEndAt(endAt).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s Survey) GetSurvey(id int64) (resp *service.Survey, err error) {
	first, err := s.db.Survey.Query().Where(survey2.IDEQ(id)).First(s.ctx)
	if err != nil {
		return nil, err
	}
	return s.entToSurvey(first), nil
}
func (s Survey) entToSurvey(v *ent.Survey) *service.Survey {

	return &service.Survey{
		ID:        v.ID,
		Title:     v.Title,
		Pic:       v.Pic,
		Desc:      v.Desc,
		StartAt:   v.StartAt.Format(time.DateTime),
		EndAt:     v.EndAt.Format(time.DateTime),
		CreatedAt: v.CreatedAt.Format(time.DateTime),
	}
}

func (s Survey) entToQuestionAll(all []*ent.SurveyQuestion, parentID int64) []*service.Question {
	if all == nil {
		return nil
	}
	var result []*service.Question
	for _, v := range all {
		if v.ParentID == parentID && v.ID != parentID {
			sq := &service.Question{
				Content:   v.Content,
				Type:      v.Type,
				Options:   v.Options,
				Required:  v.Required,
				Sort:      v.Sort,
				ID:        v.ID,
				JumpRules: v.JumpRules,
				SurveyId:  v.SurveyID,
				ParentId:  v.ParentID,
				Serial:    v.Serial,
				Show:      v.Show,
				Remark:    v.Remark,
			}

			sq.Children = s.entToQuestionAll(all, v.ID)
			result = append(result, sq)
		}
	}
	return result
}

func (s Survey) entToQuestion(v *ent.SurveyQuestion) *service.Question {

	sq := &service.Question{
		Content:   v.Content,
		Type:      v.Type,
		Options:   v.Options,
		Required:  v.Required,
		Sort:      v.Sort,
		ID:        v.ID,
		Children:  nil,
		JumpRules: v.JumpRules,
		SurveyId:  v.SurveyID,
		ParentId:  v.ParentID,
		Show:      v.Show,
		Serial:    v.Serial,
	}

	return sq
}
func (s Survey) ListSurvey(req *service.SurveyListReq) (resp []*service.Survey, total int, err error) {

	var predicates []predicate.Survey

	if req.Title != "" {
		predicates = append(predicates, survey2.Title(req.Title))
	}
	predicates = append(predicates, survey2.Delete(0))
	all, err := s.db.Survey.
		Query().
		Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	for _, v := range all {
		resp = append(resp, s.entToSurvey(v))
	}
	total, err = s.db.Survey.Query().Count(s.ctx)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}

func (s Survey) DeleteSurvey(id int64) (err error) {
	_, err = s.db.Survey.Update().
		Where(survey2.IDEQ(id)).
		SetDelete(1).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s Survey) CreateQuestion(req *service.CreateOrUpdateQuestionReq) (err error) {

	sq := s.db.SurveyQuestion.Create().
		SetContent(req.Content).
		SetType(req.Type).
		SetSurveyID(req.SurveyId).
		SetParentID(req.ParentId).
		SetSort(req.Sort).
		SetRequired(req.Required).
		SetOptions(req.Options)

	if len(req.Options) > 0 {
		sq.SetOptions(req.Options)
	}
	if len(req.JumpRules) > 0 {
		sq.SetJumpRules(req.JumpRules)
	}

	_, err = sq.Save(s.ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s Survey) UpdateQuestion(req *service.CreateOrUpdateQuestionReq) (err error) {
	sq := s.db.SurveyQuestion.Update().
		Where(surveyquestion2.IDEQ(req.ID)).
		SetContent(req.Content).
		SetType(req.Type).
		SetSurveyID(req.SurveyId).
		SetParentID(req.ParentId).
		SetSort(req.Sort).
		SetRequired(req.Required)

	if len(req.Options) > 0 {
		sq.SetOptions(req.Options)
	}
	if len(req.JumpRules) > 0 {
		sq.SetJumpRules(req.JumpRules)
	}
	_, err = sq.Save(s.ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s Survey) GetQuestion(id int64) (resp *service.Question, err error) {
	first, err := s.db.SurveyQuestion.Query().Where(surveyquestion2.IDEQ(id)).First(s.ctx)
	if err != nil {
		return nil, err
	}
	return s.entToQuestion(first), nil
}
func (s Survey) TreeQuestion(req *service.QuestionListReq) (resp []*base.Tree, err error) {
	var predicates []predicate.SurveyQuestion

	if req.SurveyId != 0 {
		predicates = append(predicates, surveyquestion2.SurveyID(req.SurveyId))
	}
	predicates = append(predicates, surveyquestion2.Delete(0))
	all, err := s.db.SurveyQuestion.
		Query().
		Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	resp = findTreeQuestionChildren(all, 0)
	return resp, nil
}

func findTreeQuestionChildren(data []*ent.SurveyQuestion, parentID int64) []*base.Tree {
	if data == nil {
		return nil
	}
	var result []*base.Tree
	for _, v := range data {
		if v.ParentID == parentID && v.ID != parentID {
			var m = new(base.Tree)
			m.Title = v.Content
			m.Value = strconv.FormatInt(v.ID, 10)
			m.Key = strconv.FormatInt(v.ID, 10)
			m.Children = findTreeQuestionChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}

func (s Survey) ListQuestion(req *service.QuestionListReq) (resp []*service.Question, total int, err error) {
	var predicates []predicate.SurveyQuestion

	if req.Content != "" {
		predicates = append(predicates, surveyquestion2.Content(req.Content))
	}

	if req.SurveyId != 0 {
		predicates = append(predicates, surveyquestion2.SurveyID(req.SurveyId))
	}
	predicates = append(predicates, surveyquestion2.Delete(0))
	all, err := s.db.SurveyQuestion.
		Query().
		Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	resp = s.entToQuestionAll(all, 0)

	total, err = s.db.Survey.Query().Count(s.ctx)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
func (s Survey) DeleteQuestion(id int64) (err error) {
	_, err = s.db.SurveyQuestion.Update().
		Where(surveyquestion2.IDEQ(id)).
		SetDelete(1).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
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
		Where(surveyquestion2.SurveyID(first.SurveyID)).
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
	if question.Level == 2 {

		for i, v := range all {
			if v == question.ID {
				return int64(i), err
			}

		}

	} else {
		split := strings.Split(question.Tree, " ")
		trimmed := strings.TrimPrefix(split[1], "tr_")
		id, err := strconv.ParseInt(trimmed, 10, 64)
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

	predicates = append(predicates)

	predicates = append(predicates, surveyresponse2.Delete(0))
	all, err := s.db.SurveyResponse.
		Query().
		Where(predicates...).Order(ent.Desc(surveyresponse2.FieldID)).
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

	aCount, _ := s.db.SurveyResponseAnswers.Query().
		Where(surveyresponseanswers2.SurveyResponseID(v.ID)).
		Count(s.ctx)

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
		Pic:             v.Pic,
		//Audio:           nil,
		IP:          v.IP,
		Device:      v.Device,
		Area:        "",
		City:        "",
		District:    "",
		Address:     "",
		Village:     "",
		AnswerCount: int64(aCount / qCount),
	}

}

func (s Survey) DeleteResponse(id int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewSurvey(ctx context.Context, c *app.RequestContext) do.Survey {
	return &Survey{
		ctx: ctx,
		c:   c,
		db:  db.DB,
	}
}
