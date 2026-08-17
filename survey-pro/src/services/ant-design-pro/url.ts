export enum Urls {
  Login = '/service/login',
  OutLogin = '/service/logout',
  UserInfo = '/service/user/info',


  CreateSurvey='/service/survey/create',
  UpdateSurvey='/service/survey/update',
  GetSurvey='/service/survey/info',
  ListSurvey='/service/survey/list',
  DeleteSurvey='/service/survey/delete',

  CreateQuestion='/service/survey/question/create',
  UpdateQuestion='/service/survey/question/update',
  GetQuestion='/service/survey/question/info',
  ListQuestion='/service/survey/question/list',
  DeleteQuestion='/service/survey/question/delete',

  TreeQuestion='/service/survey/question/tree',


  CreateRespondent='/service/survey/response/create',
  GetNext='/service/survey/response/getNext',
  GetResponse='/service/survey/response/info',
  GetResponseAnswers='/service/survey/response/answers',
  ListResponse='/service/survey/response/list',
  ListResponseExport='/service/survey/response/list-export',
  QuestionAnswersList='/service/survey/question/answers',
  Heatmap='/service/survey/response/heatmap',
  QuestionBasic='/service/survey/question/basic',
  SurveyStatistics='/service/survey/statistics',

  QueryByPhone='/service/pub/query-by-phone',

  InterviewerLogin='/service/pub/interviewer/login',
  InterviewerSurveyList='/service/survey/interviewer/list',
  InterviewerResponseUpdate='/service/survey/interviewer/response/update',
  InterviewerChangePassword='/service/survey/interviewer/change-password',
  InterviewerSetupPassword='/service/pub/interviewer/setup-password',
  InterviewerSendSms='/service/pub/interviewer/send-sms',
  InterviewerLogSearch='/service/survey/interviewer/logs',
}
