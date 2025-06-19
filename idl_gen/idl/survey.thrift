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
     8:optional i64 questionCount=0 (api.raw = "questionCount")
     9:optional i64 responseCount=0 (api.raw = "responseCount")
     10:optional string pic="" (api.raw = "pic")

}
struct SurveyListReq {
    1: optional i64 page=0 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string title="" (api.raw = "title")
}

struct CreateOrUpdateQuestionReq {
  1:optional i64 id=0 (api.raw = "id")
  2:optional string content="" (api.raw = "content")
  3:optional string type="" (api.raw = "type")
  4:optional i64 surveyId=0 (api.raw = "surveyId")
  5:optional i64 parentId=0 (api.raw = "parentId")
  6:optional i64 sort=0 (api.raw = "sort")
  7:optional i64 required=1 (api.raw = "required")
  8:optional i64 to=0 (api.raw = "to")
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
      8:optional JumpRules jumpRules (api.raw = "jumpRules")
      9:optional i64 surveyId=0 (api.raw = "surveyId")
      10:optional i64 parentId=0 (api.raw = "parentId")

}
struct JumpRules {
	  1:optional i64 questionId=0 (api.raw = "questionId")// 触发跳题的问题ID
      2:optional string answer="" (api.raw = "answer")// 触发条件的回答
      3:optional i64 nextQuestionId=0 (api.raw = "nextQuestionId")// 跳转的目标问题ID
      4:optional string operators="" (api.raw = "operators")
}


struct Options {
	  1:optional string serial="" (api.raw = "serial")
      2:optional string content="" (api.raw = "content")
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
    1: optional i64 page=0 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
}