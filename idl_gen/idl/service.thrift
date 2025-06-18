namespace go service

include "../base/base.thrift"

service ServiceService {
    //检查系统状态
    base.NilResponse HealthCheck(1: base.Empty req) (api.get = "/service/health")

}