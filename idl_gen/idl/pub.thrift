namespace go pub
include "../base/base.thrift"

struct UploadReq{
    1:  binary file
}
// pub service
service pubService {
    base.NilResponse Upload(1: UploadReq req) (api.post = "/service/pub/upload/")
}
