// @ts-ignore
/* eslint-disable */
import { Urls } from '@/services/url';
import { request } from '@umijs/max';
/** 获取当前的用户 GET /api/currentUser */
export async function currentUser(options) {
    return request(Urls.UserInfo, {
        method: 'GET',
        headers: {
            Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
        },
        ...(options || {}),
    });
}
/** 退出登录接口 POST /api/login/outLogin */
export async function outLogin(options) {
    return request(Urls.OutLogin, {
        method: 'POST',
        ...(options || {}),
    });
}
/** 登录接口 POST /api/login/account */
export async function login(body) {
    return request(Urls.Login, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
    });
}
/** 此处后端没有提供注释 GET /api/notices */
export async function getNotices(options) {
    return request('/api/notices', {
        method: 'GET',
        ...(options || {}),
    });
}
/** 获取规则列表 GET /api/rule */
export async function rule(params, options) {
    return request('/api/rule', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}
/** 更新规则 PUT /api/rule */
export async function updateRule(options) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            method: 'update',
            ...(options || {}),
        },
    });
}
/** 新建规则 POST /api/rule */
export async function addRule(options) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            method: 'post',
            ...(options || {}),
        },
    });
}
/** 删除规则 DELETE /api/rule */
export async function removeRule(options) {
    return request('/api/rule', {
        method: 'POST',
        data: {
            method: 'delete',
            ...(options || {}),
        },
    });
}
export async function queryProvince() {
    return request('/service/sys/area');
}
export async function queryCity(area) {
    return request(`/service/sys/city?id=${area}`);
}
export async function pubUpload(options) {
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
    return request('/service/pub/upload/', {
        method: 'POST',
        data: formData,
        // headers: {
        //   // 若需要认证，添加认证信息
        //   Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
        // },
    });
}
//# sourceMappingURL=api.js.map
