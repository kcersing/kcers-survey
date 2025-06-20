package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	survey2 "kcers-survey/biz/dal/db/mysql/ent/survey"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	"kcers-survey/biz/infras/do"
	"kcers-survey/biz/infras/service/common"
	"kcers-survey/biz/pkg/utils"
	"kcers-survey/idl_gen/model/service"
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
		ID:      v.ID,
		Title:   v.Title,
		Pic:     v.Pic,
		Desc:    v.Desc,
		StartAt: v.StartAt.Format(time.DateTime),
		EndAt:   v.EndAt.Format(time.DateTime),
	}
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

	_, err = s.db.SurveyQuestion.Create().
		SetContent(req.Content).
		SetType(req.Type).
		SetSurveyID(req.SurveyId).
		SetParentID(req.ParentId).
		SetSort(req.Sort).
		SetRequired(req.Required).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s Survey) UpdateQuestion(req *service.CreateOrUpdateQuestionReq) (err error) {
	_, err = s.db.SurveyQuestion.Update().
		Where(surveyquestion2.IDEQ(req.ID)).
		SetContent(req.Content).
		SetType(req.Type).
		SetSurveyID(req.SurveyId).
		SetParentID(req.ParentId).
		SetSort(req.Sort).
		SetRequired(req.Required).
		Save(s.ctx)

	if err != nil {
		return err
	}
	return nil
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
	//TODO implement me
	panic("implement me")
}

func (s Survey) UpdateResponse(req *service.CreateOrUpdateResponseReq) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s Survey) GetResponse(id int64) (resp *service.Response, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Survey) ListResponse(req *service.ResponseListReq) (resp []*service.Response, total int, err error) {
	//TODO implement me
	panic("implement me")
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
