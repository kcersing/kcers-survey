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
import { Urls } from '@/services/url';
export async function createSurvey(options) {
    return request(Urls.CreateSurvey, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function updateSurvey(options) {
    return request(Urls.UpdateSurvey, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getSurvey(options) {
    return request(Urls.GetSurvey, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function listSurvey(params, options) {
    return request(Urls.ListSurvey, {
        method: 'POST',
        params: {
            page: params.current,
            ...params,
        },
        ...(options || {}),
    });
}
export async function deleteSurvey(options) {
    return request(Urls.DeleteSurvey, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function createRespondent(options) {
    return request(Urls.CreateRespondent, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function createQuestion(options) {
    return request(Urls.CreateQuestion, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function updateQuestion(options) {
    console.log(options);
    return request(Urls.UpdateQuestion, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getQuestion(options) {
    return request(Urls.GetQuestion, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function listQuestion(params, options) {
    return request(Urls.ListQuestion, {
        method: 'POST',
        params: {
            page: params.current,
            ...params,
        },
        ...(options || {}),
    });
}
export async function treeQuestion(params, options) {
    return request(Urls.TreeQuestion, {
        method: 'POST',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}
export async function deleteQuestion(options) {
    return request(Urls.DeleteQuestion, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function createResponse(options) {
    return request(Urls.DeleteQuestion, {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getNext(options) {
    return request('/service/survey/response/getNext', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getResponse(options) {
    return request('/service/survey/response/info', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getResponseAnswers(options) {
    return request('/service/survey/response/answers', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function listResponse(params, options) {
    return request("/service/survey/response/list", {
        method: 'POST',
        params: {
            page: params.current,
            ...params,
        },
        ...(options || {}),
    });
}
export async function listResponseExport(params, options) {
    return request("/service/survey/response/list-export", {
        method: 'POST',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}
export async function getQuestionAnswersList(params, options) {
    return request("/service/survey/question/answers", {
        method: 'POST',
        params: {
            page: params.current,
            ...params,
        },
        ...(options || {}),
    });
}
export async function getHeatmap(options) {
    return request('/service/survey/response/heatmap', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function questionBasicData(options) {
    return request('/service/survey/question/basic', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
export async function getSurveyStatistics(options) {
    return request('/service/survey/statistics', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
//# sourceMappingURL=survey.js.map
