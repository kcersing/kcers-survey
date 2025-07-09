package do

import (
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/service"
)

type Survey interface {
	CreateSurvey(req *service.CreateOrUpdateSurveyReq) (err error)
	UpdateSurvey(req *service.CreateOrUpdateSurveyReq) (err error)
	GetSurvey(id int64) (resp *service.Survey, err error)
	ListSurvey(req *service.SurveyListReq) (resp []*service.Survey, total int, err error)
	DeleteSurvey(id int64) (err error)

	TreeQuestion(req *service.QuestionListReq) (resp []*base.Tree, err error)

	CreateQuestion(req *service.CreateOrUpdateQuestionReq) (err error)
	UpdateQuestion(req *service.CreateOrUpdateQuestionReq) (err error)
	GetQuestion(id int64) (resp *service.Question, err error)
	ListQuestion(req *service.QuestionListReq) (resp []*service.Question, total int, err error)

	DeleteQuestion(id int64) (err error)

	CreateResponse(req *service.CreateOrUpdateResponseReq) (err error)
	UpdateResponse(req *service.CreateOrUpdateResponseReq) (err error)
	GetResponse(req *service.ResponseAnswersReq) (resp *service.Response, err error)
	GetResponseAnswers(req *service.ResponseAnswersReq) (resp []*service.ResponseAnswers, err error)
	ListResponse(req *service.ResponseListReq) (resp []*service.Response, total int, err error)
	DeleteResponse(id int64) (err error)

	GetNext(req *service.GetNextReq) (number int64, err error)

	GetQuestionStatisticsBasic(id int64) (resp *service.StatisticsBasic, err error)
}
