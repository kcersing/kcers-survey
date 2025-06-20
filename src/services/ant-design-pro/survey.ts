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

    return request<Record<string, any>>(Urls.ListSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }
  export async function updateSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.ListSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }
  export async function getSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.ListSurvey, {
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
  },
  options?: { [key: string]: any },
) {

  return request<Surveys.Survey>(Urls.ListSurvey, {
    method: 'POST',
    params: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}


  export async function deleteSurvey(options?: { [key: string]: any }) {
    return request<Record<string, any>>(Urls.ListSurvey, {
      method: 'POST',
      data: {
        method: 'post',
        ...(options || {}),
      },
    });
  }

