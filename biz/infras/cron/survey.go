package cron

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/logs"
	"kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	surveyService "kcers-survey/biz/infras/service/survey"
	"kcers-survey/idl_gen/model/service"
)

func setResponseAnswersCount() {
	return
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

type Req struct {
	SurveyId   int64    `json:"surveyId,omitempty"`
	QuestionId int64    `json:"questionId,omitempty"`
	Type       string   `json:"type,omitempty"`
	Value      []string `json:"value,omitempty"`
	Sn         string   `json:"sn,omitempty"`
	Latitude   string   `json:"latitude,omitempty"`
	Longitude  string   `json:"longitude,omitempty"`
}

func xiufu2() {
	hlog.Info("==================================")
	all, err := db.DB.Logs.Query().
		Where(
			//func(s *sql.Selector) {
			//	s.Where(sql.Like(logs.FieldReqContent, "%"+"kdd9srslw5"+"%"))
			//},
			logs.API("/service/survey/response/create")).
		Order(ent.Asc(logs.FieldID)).
		All(context.Background())
	if err != nil {
		return
	}
	for _, v := range all {
		hlog.Info(v)
		var rwq Req
		err := json.Unmarshal([]byte(v.ReqContent), &rwq)
		if err != nil {
			continue
		}

		err = surveyService.NewSurvey(context.Background(), nil).CreateResponse(&service.CreateOrUpdateResponseReq{
			SurveyId:   rwq.SurveyId,
			Type:       rwq.Type,
			QuestionId: rwq.QuestionId,
			Value:      rwq.Value,
			Sn:         rwq.Sn,
			Latitude:   rwq.Latitude,
			Longitude:  rwq.Longitude,
		})
		if err != nil {
			return
		}
	}
	hlog.Info("完成")
}

//func xiufu() {
//
//	hlog.Info("==================================")
//	all, err := db.DB.Logs.Query().
//		Where(
//			logs.API("/service/survey/response/create")).
//		Order(ent.Asc(logs.FieldID)).
//		All(context.Background())
//	if err != nil {
//		hlog.Error(err)
//	}
//	responses, _ := db.DB.SurveyResponse.Query().
//		Where(surveyresponse2.Delete(0)).
//		Select("sn").
//		Strings(context.Background())
//
//	var snArs []string
//	for _, v := range all {
//		var rwq Req
//		err := json.Unmarshal([]byte(v.ReqContent), &rwq)
//		if err != nil {
//			continue
//		}
//		hlog.Info(rwq.Sn)
//		bl := false
//		for _, sn1 := range responses {
//			if sn1 == rwq.Sn {
//				bl = true
//			}
//		}
//		if !bl {
//			snArs = append(snArs, rwq.Sn)
//		}
//	}
//	hlog.Info("snsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsn")
//	//hlog.Info(snArs + \n")
//	fmt.Println(snArs)
//	hlog.Info("snsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsnsn")
//}
