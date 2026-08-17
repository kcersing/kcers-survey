namespace go pub
include "../base/base.thrift"

struct UploadReq{
    1:  binary file
}

struct QueryByPhoneReq{
    1:  string phone
    2:  i64 page=1
    3:  i64 page_size=10
}

struct QueryByPhoneResp{
    1:  i64 survey_id=0
    2:  string survey_title=""
    3:  string sn=""
    4:  string respondent=""
    5:  string respondent_phone=""
    6:  string address=""
    7:  string created_at=""
    8:  i64 answers_count=0
}

struct InterviewerLoginReq{
    1:  string mobile
    2:  string password
}

struct InterviewerSendSmsReq{
    1:  string mobile
}

struct InterviewerSetupReq{
    1:  string mobile
    2:  string name
    3:  string password
    4:  string smsCode
}

// pub service
service pubService {
//    base.NilResponse Upload(1: UploadReq req) (api.post = "/service/pub/upload/")
    base.NilResponse QueryByPhone(1: QueryByPhoneReq req) (api.post = "/service/pub/query-by-phone")
    base.NilResponse InterviewerLogin(1: InterviewerLoginReq req) (api.post = "/service/pub/interviewer/login")
    base.NilResponse InterviewerSendSms(1: InterviewerSendSmsReq req) (api.post = "/service/pub/interviewer/send-sms")
    base.NilResponse InterviewerSetupPassword(1: InterviewerSetupReq req) (api.post = "/service/pub/interviewer/setup-password")
}