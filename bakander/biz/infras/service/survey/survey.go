package survey

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/db/ent"
	"kcers-survey/biz/dal/db/ent/area"
	"kcers-survey/biz/dal/db/ent/predicate"
	survey2 "kcers-survey/biz/dal/db/ent/survey"
	"kcers-survey/biz/infras/do"
	"kcers-survey/biz/infras/service/common"
	"kcers-survey/biz/pkg/utils"
	"kcers-survey/idl_gen/model/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
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
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,
		MaxCost:     1 << 25, // 32MB
		BufferItems: 64,
	})
	return &Survey{
		ctx:   ctx,
		c:     c,
		db:    db.DB,
		cache: cache,
	}
}

// getAreaName 通过缓存获取区域名称，避免 N+1 查询
func (s Survey) getAreaName(idStr string) string {
	if idStr == "" || s.cache == nil {
		return ""
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ""
	}

	cacheKey := fmt.Sprintf("area:name:%d", id)
	if name, ok := s.cache.Get(cacheKey); ok {
		return name.(string)
	}

	first, err := s.db.Area.Query().Where(area.ID(id)).First(s.ctx)
	if err != nil {
		return idStr
	}

	s.cache.SetWithTTL(cacheKey, first.Name, 1, 10*time.Minute)
	return first.Name
}
