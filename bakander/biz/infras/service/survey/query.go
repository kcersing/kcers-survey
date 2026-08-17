package survey

import (
	"time"

	"kcers-survey/biz/dal/db/ent"
	survey2 "kcers-survey/biz/dal/db/ent/survey"
	surveyresponse2 "kcers-survey/biz/dal/db/ent/surveyresponse"
	"kcers-survey/idl_gen/model/pub"
)

func (s Survey) QueryByPhone(req *pub.QueryByPhoneReq) (resp []*pub.QueryByPhoneResp, total int, err error) {
	predicates := []surveyresponse2.OrderOption{
		ent.Desc(surveyresponse2.FieldCreatedAt),
	}

	all, err := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.ResearcherPhone(req.Phone),
			surveyresponse2.SnNEQ(""),
			surveyresponse2.AnswersCountNEQ(0),
			surveyresponse2.Delete(0),
		).
		Order(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		All(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	surveyIDs := make([]int64, 0, len(all))
	for _, v := range all {
		surveyIDs = append(surveyIDs, v.SurveyID)
	}

	surveyMap := make(map[int64]string)
	if len(surveyIDs) > 0 {
		surveys, err := s.db.Survey.Query().
			Where(survey2.IDIn(surveyIDs...)).
			All(s.ctx)
		if err != nil {
			return nil, 0, err
		}
		for _, sv := range surveys {
			surveyMap[sv.ID] = sv.Title
		}
	}

	for _, v := range all {
		resp = append(resp, &pub.QueryByPhoneResp{
			SurveyID:        v.SurveyID,
			SurveyTitle:     surveyMap[v.SurveyID],
			Sn:              v.Sn,
			Respondent:      v.Respondent,
			RespondentPhone: v.RespondentPhone,
			Address:         v.Address,
			CreatedAt:       v.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
			AnswersCount:    v.AnswersCount,
		})
	}

	total, err = s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.ResearcherPhone(req.Phone),
			surveyresponse2.SnNEQ(""),
			surveyresponse2.AnswersCountNEQ(0),
			surveyresponse2.Delete(0),
		).
		Count(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}