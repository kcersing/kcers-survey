package cron

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
)

func setResponseAnswersCount() {

	var scAll []struct {
		Count            int64 `json:"count"`
		SurveyResponseID int64 `json:"survey_response_id"`
	}
	err := db.DB.SurveyResponseAnswers.
		Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(sql.Count("*"), "count"),
				surveyresponseanswers.FieldSurveyResponseID,
			).
				GroupBy(surveyresponseanswers.FieldSurveyResponseID).
				OrderBy(sql.Desc(surveyresponseanswers.FieldSurveyResponseID))
		}).
		Scan(context.Background(), &scAll)
	if err != nil {
		hlog.Error(err)
		return
	}
	for _, sc := range scAll {
		err = db.DB.SurveyResponse.UpdateOneID(sc.SurveyResponseID).
			SetAnswersCount(sc.Count).
			Exec(context.Background())
		if err != nil {
			hlog.Error(err)
			return
		}
	}

}
