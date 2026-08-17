/*
 * @Author: kcersing wt4@live.cn
 * @Date: 2025-06-20 15:01:39
 * @LastEditors: kcersing wt4@live.cn
 * @LastEditTime: 2025-06-20 15:50:39
 * @FilePath: \ant-web\src\services\ant-design-pro\survey.ts
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import { Urls } from '@/services/ant-design-pro/url';



export async function createSurvey(options?: { [key: string]: any }) {

    return request<Record<string, any>>(Urls.CreateSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }
  export async function updateSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.UpdateSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }
  export async function getSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.GetSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }

export async function listSurvey(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
    keywords?: string;
  },
  options?: { [key: string]: any },
) {

  return request<Record<string, any>>(Urls.ListSurvey, {
    method: 'POST',
    params: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}

  export async function deleteSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.DeleteSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }


export async function createRespondent(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.CreateRespondent, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function createQuestion(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.CreateQuestion, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
export async function updateQuestion(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.UpdateQuestion, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
export async function getQuestion(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.GetQuestion, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function listQuestion(
    params: {
      surveyId?: number;
      keywords?:string;
      // query
      /** 当前的页码 */
      current?: number;
      /** 页面的容量 */
      pageSize?: number;
    },
    options?: { [key: string]: any },
) {
  return request<Record<string, any>>(Urls.ListQuestion, {
    method: 'POST',
    params: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}




export async function treeQuestion(
  params: {
    surveyId?: number;
    keywords?:string;
  },
  options?: { [key: string]: any },
) {
  return request<Record<string, any>>(Urls.TreeQuestion, {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}


export async function deleteQuestion(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.DeleteQuestion, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function createResponse(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.CreateRespondent, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
export async function getNext(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.GetNext, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function getResponse(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.GetResponse, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
export async function getResponseAnswers(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.GetResponseAnswers, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}




export async function listResponse(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    surveyId?: number;
    /** 页面的容量 */
    pageSize?: number;
    keywords?: string;
  },
  options?: { [key: string]: any },
) {

  return request<Record<string, any>>(Urls.ListResponse, {
    method: 'POST',
    params: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}
export async function listResponseExport(
  params: {
    surveyId?: number;
    keywords?: string;
  },
  options?: { [key: string]: any },
) {

  return request<Record<string, any>>(Urls.ListResponseExport, {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

export async function getQuestionAnswersList(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    surveyId?: number;
    /** 页面的容量 */
    pageSize?: number;
    keywords?: string;
  },
  options?: { [key: string]: any },
) {

  return request<Record<string, any>>(Urls.QuestionAnswersList, {
    method: 'POST',
    params: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}

export async function getHeatmap(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.Heatmap, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function questionBasicData(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.QuestionBasic, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
export async function getSurveyStatistics(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.SurveyStatistics, {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

export async function queryByPhone(
  params: {
    phone: string;
    current?: number;
    pageSize?: number;
  },
) {
  return request<Record<string, any>>(Urls.QueryByPhone, {
    method: 'POST',
    data: {
      phone: params.phone,
      page: params.current || 1,
      page_size: params.pageSize || 10,
    },
  });
}

export async function interviewerLogin(data: { mobile: string; password: string }) {
  return request<Record<string, any>>(Urls.InterviewerLogin, {
    method: 'POST',
    data,
  });
}

export async function interviewerSurveyList(
  params: { current?: number; pageSize?: number },
  token: string,
) {
  return request<Record<string, any>>(Urls.InterviewerSurveyList, {
    method: 'POST',
    headers: {
      Authorization: 'Bearer ' + token,
    },
    data: {
      page: params.current || 1,
      pageSize: params.pageSize || 10,
    },
  });
}

export async function interviewerResponseUpdate(
  data: {
    surveyId: number;
    sn: string;
    questionId: number;
    answer: string[];
    answerText: string;
    type: string;
  },
  token: string,
) {
  return request<Record<string, any>>(Urls.InterviewerResponseUpdate, {
    method: 'POST',
    headers: {
      Authorization: 'Bearer ' + token,
    },
    data,
  });
}

export async function interviewerSendSms(data: { mobile: string }) {
  return request<Record<string, any>>(Urls.InterviewerSendSms, {
    method: 'POST',
    data,
  });
}

export async function interviewerSetupPassword(data: {
  mobile: string;
  name: string;
  password: string;
}) {
  return request<Record<string, any>>(Urls.InterviewerSetupPassword, {
    method: 'POST',
    data,
  });
}

export async function interviewerChangePassword(
  data: { oldPassword: string; newPassword: string },
  token: string,
) {
  return request<Record<string, any>>(Urls.InterviewerChangePassword, {
    method: 'POST',
    headers: {
      Authorization: 'Bearer ' + token,
    },
    data,
  });
}

export async function interviewerLogSearch(
  params: { keyword: string; current?: number; pageSize?: number },
  token: string,
) {
  return request<Record<string, any>>(Urls.InterviewerLogSearch, {
    method: 'POST',
    headers: {
      Authorization: 'Bearer ' + token,
    },
    data: {
      keyword: params.keyword,
      page: params.current || 1,
      pageSize: params.pageSize || 20,
    },
  });
}
