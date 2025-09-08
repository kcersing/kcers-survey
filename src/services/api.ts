// @ts-ignore
/* eslint-disable */
import { Urls } from '@/services/url';
import { request } from '@umijs/max';

/** 获取当前的用户 GET /api/currentUser */
export async function currentUser(options?: { [key: string]: any }) {
  return request<{
    data: API.CurrentUser;
  }>(Urls.UserInfo, {
    method: 'GET',
    headers: {
      Authorization:'Bearer ' + sessionStorage.getItem('token') || '',
    },
    ...(options || {}),
  });
}

/** 退出登录接口 POST /api/login/outLogin */
export async function outLogin(options?: { [key: string]: any }) {
  return request<Record<string, any>>(Urls.OutLogin, {
    method: 'POST',
    ...(options || {}),
  });
}

/** 登录接口 POST /api/login/account */
export async function login(body: API.LoginParams) {
  return request<API.LoginResult>(Urls.Login, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
  });
}

/** 此处后端没有提供注释 GET /api/notices */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/api/notices', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取规则列表 GET /api/rule */
export async function rule(
  params: {
    // query
    /** 当前的页码 */
    page?: number;
    /** 页面的容量 */
    pageSize?: number;
  },
  options?: { [key: string]: any },
) {
  return request<API.RuleList>('/api/rule', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 更新规则 PUT /api/rule */
export async function updateRule(options?: { [key: string]: any }) {
  return request<API.RuleListItem>('/api/rule', {
    method: 'POST',
    data: {
      method: 'update',
      ...(options || {}),
    },
  });
}

/** 新建规则 POST /api/rule */
export async function addRule(options?: { [key: string]: any }) {
  return request<API.RuleListItem>('/api/rule', {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}

/** 删除规则 DELETE /api/rule */
export async function removeRule(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/rule', {
    method: 'POST',
    data: {
      method: 'delete',
      ...(options || {}),
    },
  });
}


export type AreaItemType = {
  name: string;
  id: string;
};

export async function queryProvince(): Promise<{ data: AreaItemType[] }> {
  return request('/service/sys/area');
}

export async function queryCity(area: string): Promise<{ data: AreaItemType[] }> {
  return request(`/service/sys/city?id=${area}`);
}

type PubUploadOptions = {
  file: File; // 假设需要上传文件
  // 可以添加其他可选参数
  [key: string]: any;
};

export async function pubUpload(options?: PubUploadOptions) {
  const formData = new FormData();
  if (options?.file) {
    formData.append('files', options.file);
  }
  // 添加其他参数
  if (options) {
    Object.keys(options).forEach(key => {
      if (key !== 'file') {
        formData.append(key, options[key]);
      }
    });
  }

  return request<Record<string, any>>('/service/pub/upload/', {
    method: 'POST',
    data: formData,
    // headers: {
    //   // 若需要认证，添加认证信息
    //   Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
    // },
  });
}


export async function fetchMenuData(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/service/sys/menu-list', {
    method: 'POST',
    data: {
      method: 'post',
      ...(options || {}),
    },
  });
}
