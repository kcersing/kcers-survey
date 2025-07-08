namespace go service

include "../base/base.thrift"

service SurveyService {

	base.NilResponse CreateSurvey(1: CreateOrUpdateSurveyReq req)  (api.post = "/service/survey/create")
	base.NilResponse UpdateSurvey(1: CreateOrUpdateSurveyReq req)  (api.post = "/service/survey/update")
	base.NilResponse GetSurvey(1: base.IDReq req) (api.post = "/service/survey/info")
	base.NilResponse ListSurvey(1: SurveyListReq req) (api.post = "/service/survey/list")

	base.NilResponse DeleteSurvey(1: base.IDReq req) (api.post = "/service/survey/delete")

	base.NilResponse CreateQuestion(1: CreateOrUpdateQuestionReq req)  (api.post = "/service/survey/question/create")
	base.NilResponse UpdateQuestion(1: CreateOrUpdateQuestionReq req)  (api.post = "/service/survey/question/update")
    base.NilResponse GetQuestion(1: base.IDReq req)  (api.post = "/service/survey/question/info")
	base.NilResponse ListQuestion(1: QuestionListReq req)  (api.post = "/service/survey/question/list")
	base.NilResponse DeleteQuestion(1: base.IDReq req)  (api.post = "/service/survey/question/delete")

	base.NilResponse TreeQuestion(1: QuestionListReq req)  (api.post = "/service/survey/question/tree")


	base.NilResponse CreateResponse(1: CreateOrUpdateResponseReq req) (api.post = "/service/survey/response/create")
	base.NilResponse UpdateResponse(1: CreateOrUpdateResponseReq req) (api.post = "/service/survey/response/update")
	base.NilResponse GetResponse(1: ResponseAnswersReq req)  (api.post = "/service/survey/response/info")
    base.NilResponse GetResponseAnswers(1: ResponseAnswersReq req)  (api.post = "/service/survey/response/answers")


	base.NilResponse ListResponse(1: ResponseListReq req) (api.post = "/service/survey/response/list")
	base.NilResponse DeleteResponse(1: base.IDReq req)  (api.post = "/service/survey/response/delete")


	base.NilResponse GetNext(1: GetNextReq req)  (api.post = "/service/survey/response/getNext")





}
struct GetNextReq{
  1:optional string sn="" (api.raw = "sn")
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


      255:  optional string createdAt="" (api.raw = "createdAt")
      256:  optional string updatedAt="" (api.raw = "updatedAt")

}
struct SurveyListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string title="" (api.raw = "title")
}
struct QuestionListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional i64 surveyId=0 (api.raw = "surveyId")
    4: optional string content="" (api.raw = "content")

}

struct CreateOrUpdateQuestionReq {
  1:optional i64 id=0 (api.raw = "id")
  2:optional string content="" (api.raw = "content")
  3:optional string type="" (api.raw = "type")
  4:optional i64 surveyId=0 (api.raw = "surveyId")
  5:optional i64 parentId=0 (api.raw = "parentId")
  6:optional i64 sort=0 (api.raw = "sort")
  7:optional i64 required=1 (api.raw = "required")
  8:optional list<JumpRules> jumpRules={} (api.raw = "jumpRules")
  9:optional list<Options> options={} (api.raw = "options")
}


// QuestionRequest 问题请求
struct Question  {
	  1:optional string content="" (api.raw = "content")
      2:optional string type="" (api.raw = "type")
      3:optional list<Options> options={} (api.raw = "options")
      4:optional i64 required=1 (api.raw = "required")
      5:optional i64 sort=0 (api.raw = "sort")
      6:optional i64 id=0 (api.raw = "id")
      7:optional list<Question> children={} (api.raw = "childrens")
      8:optional list<JumpRules> jumpRules={} (api.raw = "jumpRules")
      9:optional i64 surveyId=0 (api.raw = "surveyId")
      10:optional i64 parentId=0 (api.raw = "parentId")
      11:optional string serial="" (api.raw = "serial")
      12:optional i64 valueNumber=0 (api.raw = "valueNumber")
      13:optional i64 show=0 (api.raw = "show")
      14:optional string remark="" (api.raw = "remark")

         15:optional string answerCreatedAt="" (api.raw = "answerCreatedAt")
         16:optional list<string> answer="" (api.raw = "answer")
         17:optional string answerText="" (api.raw = "answerText")
         18:optional i64 responseId=0 (api.raw = "responseId")


    255:  optional string createdAt="" (api.raw = "createdAt")
    256:  optional string updatedAt="" (api.raw = "updatedAt")
}
struct JumpRules {
      2:optional string answer="" (api.raw = "answer")// 触发条件的回答
      3:optional i64 nextQuestionId=0 (api.raw = "nextQuestionId")// 跳转的目标问题ID
      4:optional string operators="" (api.raw = "operators")
}

struct Options {
	  1:optional i64 serial=1 (api.raw = "serial")
      2:optional string content="" (api.raw = "content")
      3:optional i64 inputs=1 (api.raw = "inputs")
}

struct CreateOrUpdateResponseReq {
  1:optional i64 surveyId=0 (api.raw = "surveyId")
  3:optional string type="" (api.raw = "type")
  4:optional i64 questionId=0 (api.raw = "questionId")
  5:optional list<string> value="" (api.raw = "value")
  6:optional string sn="" (api.raw = "sn")
  7:optional string latitude="" (api.raw = "latitude")
  8:optional string longitude="" (api.raw = "longitude")
}

struct Response  {
    3:optional i64 id=0 (api.raw = "id")
    4:optional i64 surveyId=0 (api.raw = "surveyId")
    5:optional string sn="" (api.raw = "sn")
    6:optional string latitude="" (api.raw = "latitude")
    7:optional string longitude="" (api.raw = "longitude")
    8:optional list<Question> questions={} (api.raw = "questions")
    9:optional string respondent="" (api.raw = "respondent")
    10:optional string respondentPhone="" (api.raw = "respondentPhone")
    11:optional string researcher="" (api.raw = "researcher")
    12:optional string researcherPhone="" (api.raw = "researcherPhone")
    13:optional list<string> pic="" (api.raw = "researcherPhone")
    14:optional list<string> audio="" (api.raw = "audio")

    15:optional string ip="" (api.raw = "ip")
    16:optional string device="" (api.raw = "device")
    17:optional string area="" (api.raw = "area")
    18:optional string city="" (api.raw = "city")
    19:optional string district="" (api.raw = "district")
    20:optional string address="" (api.raw = "address")
    21:optional string village="" (api.raw = "village")
    22:optional i64 answerCount=0 (api.raw = "answerCount")


     255:  optional string createdAt="" (api.raw = "createdAt")
     256:  optional string updatedAt="" (api.raw = "updatedAt")
}

struct ResponseListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3:optional string sn="" (api.raw = "sn")
    4: optional i64 surveyId=0 (api.raw = "surveyId")
    9:optional string respondent="" (api.raw = "respondent")
    10:optional string respondentPhone="" (api.raw = "respondentPhone")
    11:optional string researcher="" (api.raw = "researcher")
    12:optional string researcherPhone="" (api.raw = "researcherPhone")
    13:optional string sorter="" (api.raw = "sorter")

}
  struct ResponseAnswersReq   {
      1: optional i64 id=0 (api.raw = "id")
      3:optional string sn="" (api.raw = "sn")
  }
  struct ResponseAnswers   {

      1: optional i64 id=0 (api.raw = "id")
      2: optional i64 surveyId=0 (api.raw = "surveyId")
      3: optional i64 surveyResponseId=0 (api.raw = "surveyResponseId")
      4: optional i64 surveyQuestionId=0 (api.raw = "surveyQuestionId")
       5: optional  list<string> answer={} (api.raw = "answer")
      6:optional string answerText="" (api.raw = "answerText")
      7:optional string createdAt="" (api.raw = "createdAt")
     8:optional string content="" (api.raw = "content")

  }





