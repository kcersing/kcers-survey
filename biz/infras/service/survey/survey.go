package survey

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	survey2 "kcers-survey/biz/dal/db/mysql/ent/survey"
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
		ID:        v.ID,
		Title:     v.Title,
		Pic:       v.Pic,
		Desc:      v.Desc,
		StartAt:   v.StartAt.Format(time.DateTime),
		EndAt:     v.EndAt.Format(time.DateTime),
		CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
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
		Order(ent.Desc(survey2.FieldID)).
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

func NewSurvey(ctx context.Context, c *app.RequestContext) do.Survey {
	return &Survey{
		ctx: ctx,
		c:   c,
		db:  db.DB,
	}
}
