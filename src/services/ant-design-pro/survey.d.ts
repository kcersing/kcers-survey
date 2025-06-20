// @ts-ignore
/* eslint-disable */

declare namespace SURVEY {

  type Survey = {
    id?: number;
    title?: string;
    desc?: string;
    status?: number;
    startAt?: string;
    endAt?: string;
    questions?: Questions[];
    questionCount?: number;
    responseCount?: number;
    pic?: string;
  };

  type Questions = {
    id?: number;
    sort?: number;
    type?: string;
    content?: string;
    required?: number;
    options?: Options[];
    subQuestions?: Questions[];
    surveyId?: number;
    parentId?: number;
    jumpRules?: JumpRules;
  };
  type JumpRules = {
    questionId?: number;// 触发跳题的问题ID
    answer?: string;// 触发条件的回答
    nextQuestionId?: number;// 跳转的目标问题ID
    operators?: string;
  };
  type Options = {
    serial?: string;
    content?: string;
  };
}
