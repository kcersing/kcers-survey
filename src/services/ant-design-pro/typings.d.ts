// @ts-ignore
/* eslint-disable */

declare namespace API {
  type CurrentUser = {
    // name?: string;
    // avatar?: string;
    // userid?: string;
    // email?: string;
    // signature?: string;
    // title?: string;
    // group?: string;
    // tags?: { key?: string; label?: string }[];
    // notifyCount?: number;
    // unreadCount?: number;
    // country?: string;
    // access?: string;
    // geographic?: {
    //   province?: { label?: string; key?: string };
    //   city?: { label?: string; key?: string };
    // };
    // address?: string;
    // phone?: string;

    avatar?: string;
    birthday?: string;
    createdAt?: string;
    detail?: string;
    email?: string;
    gender?: string;
    id?: number;
    mobile?: string;
    name?: string;
    status?: number;
    updatedAt?: string;
    username?: string;
    wecom?: string;

    userRole?: { id?: number; name?: string;value?: string; }[];
    userRoleIds?: number[];

  };


  type LoginResult = {
    code?: number;
    data?: { token?: string; expire?: string };
    currentAuthority?: string;
  };

  type PageParams = {
    page?: number;
    pageSize?: number;
  };

  type RuleListItem = {
    key?: number;
    disabled?: boolean;
    href?: string;
    avatar?: string;
    name?: string;
    owner?: string;
    desc?: string;
    callNo?: number;
    status?: number;
    updatedAt?: string;
    createdAt?: string;
    progress?: number;
  };

  type RuleList = {
    data?: RuleListItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type LoginParams = {
    username?: string;
    password?: string;
    autoLogin?: boolean;
    type?: string;
    captcha?: string;
    captchaId?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    code: number;
    /** 业务上的错误信息 */
    message?: string;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };





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
    children?: Questions[];
    surveyId?: number;
    parentId?: number;
    jumpRules?: JumpRules;
    matrixRows?: string;
    matrixColumns?: string;

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
