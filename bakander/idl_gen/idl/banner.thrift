namespace go banner

include "../base/base.thrift"

// Banner
service BannerService {

  base.NilResponse CreateBanner(1: BannerInfo req) (api.post = "/service/banner/create")

  base.NilResponse UpdateBanner(1: BannerInfo req) (api.post = "/service/banner/update")

  base.NilResponse DelBanner(1: base.IDReq req) (api.post = "/service/banner/del")

  base.NilResponse BannerList(1: BannerListReq req) (api.post = "/service/banner/list")

  base.NilResponse UpdateBannerShow(1: base.StatusCodeReq req) (api.post = "/service/banner/show")
}

struct BannerInfo {
  1:optional string name="" (api.raw = "name")
  2:optional string pic="" (api.raw = "pic")
  3:optional string link="" (api.raw = "link")
  4:optional i64 isShow=1 (api.raw = "isShow")
  5:optional i64 id=0 (api.raw = "id")
}
struct BannerListReq {
  1:optional i64 page=1 (api.raw = "page")
  2:optional i64 pageSize=100 (api.raw = "pageSize")
 }
