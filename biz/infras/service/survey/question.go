package survey

import (
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	surveyquestion2 "kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/service"
	"strconv"
)

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
