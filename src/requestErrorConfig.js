import { message, notification } from 'antd';
// 错误处理方案： 错误类型
var ErrorShowType;
(function (ErrorShowType) {
    ErrorShowType[ErrorShowType["SILENT"] = 0] = "SILENT";
    ErrorShowType[ErrorShowType["WARN_MESSAGE"] = 1] = "WARN_MESSAGE";
    ErrorShowType[ErrorShowType["ERROR_MESSAGE"] = 2] = "ERROR_MESSAGE";
    ErrorShowType[ErrorShowType["NOTIFICATION"] = 3] = "NOTIFICATION";
    ErrorShowType[ErrorShowType["REDIRECT"] = 9] = "REDIRECT";
})(ErrorShowType || (ErrorShowType = {}));
export const errorConfig = {
    // 错误处理： umi@3 的错误处理方案。
    errorConfig: {
        // 错误抛出
        errorThrower: (res) => {
            const { code, data, message } = res;
            if (code === 0) {
                const error = new Error(message);
                error.name = 'BizError';
                error.info = { code, message, data };
                throw error; // 抛出自制的错误
            }
        },
        // 错误接收及处理
        errorHandler: (error, opts) => {
            if (opts?.skipErrorHandler)
                throw error;
            // 我们的 errorThrower 抛出的错误。
            if (error.name === 'BizError') {
                const errorInfo = error.info;
                if (errorInfo) {
                    switch (errorInfo.code) {
                        case ErrorShowType.SILENT:
                            // do nothing
                            break;
                        case ErrorShowType.WARN_MESSAGE:
                            message.warning(errorInfo.message);
                            break;
                        case ErrorShowType.ERROR_MESSAGE:
                            message.error(errorInfo.message);
                            break;
                        case ErrorShowType.NOTIFICATION:
                            notification.open({
                                description: errorInfo.message,
                                message: errorInfo.code,
                            });
                            break;
                        case ErrorShowType.REDIRECT:
                            // TODO: redirect
                            break;
                        default:
                            message.error(errorInfo.message);
                    }
                }
            }
            else if (error.response) {
                // Axios 的错误
                // 请求成功发出且服务器也响应了状态码，但状态代码超出了 2xx 的范围
                message.error(`Response status:${error.response.status}`);
            }
            else if (error.request) {
                // 请求已经成功发起，但没有收到响应
                // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
                // 而在node.js中是 http.ClientRequest 的实例
                message.error('None response! Please retry.');
            }
            else {
                // 发送请求时出了点问题
                message.error('Request error, please retry.');
            }
        },
    },
    // 请求拦截器
    requestInterceptors: [
        (config) => {
            // 拦截请求配置，进行个性化处理。
            // const url = config?.url?.concat('?token = 123');
            const url = config?.url?.concat();
            return { ...config, url };
        },
    ],
    // 响应拦截器
    responseInterceptors: [
        (response) => {
            // 拦截响应数据，进行个性化处理
            const { data } = response;
            if (data?.success === false) {
                message.error('请求失败！');
            }
            return response;
        },
    ],
};
//# sourceMappingURL=requestErrorConfig.js.map