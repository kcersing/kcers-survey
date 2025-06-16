package do

import "kcers-survey/idl_gen/model/service"

type Survey interface {
	CreateSurvey(req *service.CreateOrUpdateSurveyReq) (err error)
	UpdateSurvey(req *service.CreateOrUpdateSurveyReq) (err error)
	GetSurvey(id int64) (resp *service.Survey, err error)
	ListSurvey(req *service.SurveyListReq) (resp []*service.Survey, total int, err error)

	DeleteSurvey(id int64) (err error)

	CreateQuestion(req *service.CreateOrUpdateQuestionReq) (err error)
	UpdateQuestion(req *service.CreateOrUpdateQuestionReq) (err error)

	DeleteQuestion(id int64) (err error)

	CreateResponse(req *service.CreateOrUpdateResponseReq) (err error)
	UpdateResponse(req *service.CreateOrUpdateResponseReq) (err error)
	GetResponse(id int64) (resp *service.Response, err error)
	ListResponse(req *service.ResponseListReq) (resp []*service.Response, total int, err error)
	DeleteResponse(id int64) (err error)
}

//{
//    "title": "用户满意度调查",
//    "description": "感谢您参与我们的满意度调查",
//    "status": "draft",
//    "questions": [
//        {
//            "content": "您对我们产品的整体满意度如何？",
//            "question_type": "single_choice",
//            "options": ["非常满意", "满意", "一般", "不满意", "非常不满意"],
//            "required": true,
//            "sort_order": 1
//        },
//        {
//            "content": "您使用我们产品的频率是？",
//            "question_type": "single_choice",
//            "options": ["每天", "每周几次", "每月几次", "很少使用"],
//            "required": true,
//            "sort_order": 2,
//            "subquestions": [
//                {
//                    "content": "您最喜欢我们产品的哪个功能？",
//                    "question_type": "multiple_choice",
//                    "options": ["功能A", "功能B", "功能C", "功能D"],
//                    "required": false,
//                    "sort_order": 1
//                }
//            ]
//        },
//        {
//            "content": "您对我们服务有什么建议？",
//            "question_type": "text",
//            "required": false,
//            "sort_order": 3
//        }
//    ]
//}
