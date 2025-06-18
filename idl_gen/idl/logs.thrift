namespace go logs

include "../base/base.thrift"

//日志列表请求数据
struct LogsListReq {
    1: optional i64 page=0 (api.raw = "page")
    2: optional i64 pageSize=0 (api.raw = "pageSize")
    3: optional string type="" (api.raw = "type")
    4: optional string method="" (api.raw = "method")
    5: optional string api="" (api.raw = "api")
    6: optional bool success= true (api.raw = "success")
    7: optional string operatorsr ="" (api.raw = "operatorsr")
    8: optional i64 identity =0(api.raw = "identity")
}

//日志信息
struct LogsInfo {
    1: optional string type="" (api.raw = "type")
    2: optional string method =""(api.raw = "method")
    3: optional string api =""(api.raw = "api")
    4: optional bool success = true (api.raw = "success")
    5: optional string reqContent="" (api.raw = "reqContent")
    6: optional string respContent="" (api.raw = "respContent")
    7: optional string ip="" (api.raw = "ip")
    8: optional string userAgent =""(api.raw = "userAgent")
    9: optional string operatorsr =""(api.raw = "operatorsr")
    10: optional i64 time=0 (api.raw = "time")
    11: optional string createdAt =""(api.raw = "createdAt")
    12: optional string updatedAt =""(api.raw = "updatedAt")
    13: optional i64 identity =0(api.raw = "identity")

    251: optional i64 id = 0 (api.raw = "id")
}

service LogsService{
  // Get logs list | 获取日志列表
  base.NilResponse GetLogsList(1: LogsListReq req) (api.post = "/service/logs/list")

  // Delete logs | 删除日志信息
  base.NilResponse DeleteLogs(1: base.Empty req) (api.post = "/service/logs/deleteAll")

}