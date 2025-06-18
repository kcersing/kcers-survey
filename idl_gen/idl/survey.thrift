namespace go service

include "../base/base.thrift"

service SurveyService {


	base.NilResponse CreateSurvey(1: CreateOrUpdateSurveyReq req)  (api.get = "/service/survey/create")
	base.NilResponse UpdateSurvey(1: CreateOrUpdateSurveyReq req)  (api.get = "/service/survey/update")
	base.NilResponse GetSurvey(1: base.IDReq req) (api.get = "/service/survey/info")
	base.NilResponse ListSurvey(1: SurveyListReq req) (api.get = "/service/survey/list")

	base.NilResponse DeleteSurvey(1: base.IDReq req) (api.get = "/service/survey/delete")

	base.NilResponse CreateQuestion(1: CreateOrUpdateQuestionReq req)  (api.get = "/service/survey/question-create")
	base.NilResponse UpdateQuestion(1: CreateOrUpdateQuestionReq req)  (api.get = "/service/survey/question-update")

	base.NilResponse DeleteQuestion(1: base.IDReq req)  (api.get = "/service/survey/question-delete")

	base.NilResponse CreateResponse(1: CreateOrUpdateResponseReq req) (api.get = "/service/survey/response-create")
	base.NilResponse UpdateResponse(1: CreateOrUpdateResponseReq req) (api.get = "/service/survey/response-update")
	base.NilResponse GetResponse(1: base.IDReq req)  (api.get = "/service/survey/response-info")
	base.NilResponse ListResponse(1: ResponseListReq req) (api.get = "/service/survey/response-list")
	base.NilResponse DeleteResponse(1: base.IDReq req)  (api.get = "/service/survey/response-delete")


}

struct CreateOrUpdateSurveyReq {
  1:optional string title="" (api.raw = "title")
  2:optional string pic="" (api.raw = "pic")
  3:optional string desc="" (api.raw = "desc")
  4:optional string startAt="" (api.raw = "startAt")
  5:optional string endAt="" (api.raw = "endAt")
  6:optional i64 id=0 (api.raw = "id")
}

struct Survey {
     1:optional i64 id=0 (api.raw = "id")
     2:optional string title="" (api.raw = "title")
     3:optional string desc="" (api.raw = "desc")
     4:optional i64 status=1 (api.raw = "status")
     5:optional string startAt="" (api.raw = "startAt")
     6:optional string endAt="" (api.raw = "endAt")
     7:optional list<Question> questions={} (api.raw = "questions")

}
struct SurveyListReq {

}

struct CreateOrUpdateQuestionReq {
  1:optional string name="" (api.raw = "name")
  2:optional string pic="" (api.raw = "pic")
  3:optional string link="" (api.raw = "link")
  4:optional i64 isShow=1 (api.raw = "isShow")
  5:optional i64 id=0 (api.raw = "id")
}

// QuestionRequest 问题请求
struct Question  {
	  1:optional string content="" (api.raw = "content")
      2:optional string type="" (api.raw = "type")
      3:optional list<Options> options (api.raw = "options")
      4:optional i64 required=1 (api.raw = "required")
      5:optional i64 sort=0 (api.raw = "sort")
      6:optional i64 id=0 (api.raw = "id")
      7:optional list<Question> subQuestions={} (api.raw = "subQuestions")
}

struct Options {
	  1:optional string key="" (api.raw = "key")
      2:optional string value="" (api.raw = "value")
}

struct CreateOrUpdateResponseReq {
  1:optional i64 surveyId=0 (api.raw = "surveyId")
  2:optional list<Answer> answers={} (api.raw = "answers")
}
struct Answer  {
    1:optional i64 questionId=0 (api.raw = "questionId")
    2:optional string answers="" (api.raw = "answers")
    3:optional list<Answer> subAnswers={} (api.raw = "subAnswers")
}

struct Response  {
    1:optional i64 Id=0 (api.raw = "Id")

}

struct ResponseListReq {

}