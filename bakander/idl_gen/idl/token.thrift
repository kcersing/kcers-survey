namespace go token

include "../base/base.thrift"

// Token信息
struct TokenInfo {
    1: optional i64 id  =0(api.raw = "id")
    2: optional string createdAt ="" (api.raw = "createdAt")
    3: optional string updatedAt  =""(api.raw = "updatedAt")
    4: optional i64 userId =0 (api.raw = "userId")
    5: optional string username  =""(api.raw = "username")
    6: optional string token ="" (api.raw = "token")
    7: optional string source  =""(api.raw = "source")
    8: optional string expiredAt ="" (api.raw = "expiredAt")
}
struct MemberTokenInfo {
    1:  i64 id (api.raw = "id")
    2:  string createdAt (api.raw = "createdAt")
    3:  string updatedAt (api.raw = "updatedAt")
    4:  i64 memberId (api.raw = "memberId")
    5:  string name (api.raw = "username")
    6:  string token (api.raw = "token")
    7:  string source (api.raw = "source")
    8:  string expiredAt (api.raw = "expiredAt")
    10: string mobile (api.raw = "mobile")
}
// token列表请求参数
struct TokenListReq {
    1: optional i64 page =0 (api.raw = "page")
    2: optional i64 pageSize =0 (api.raw = "pageSize")
    3: optional string username   =""(api.raw = "username")
    4: optional i64 userId =0 (api.raw = "userId")
}

service TokenService{
  // 更新Token
  base.NilResponse UpdateToken(1: TokenInfo req) (api.post = "/service/token/update")

  // 删除token信息
  base.NilResponse DeleteToken(1: base.IDReq req) (api.post = "/service/token")

  // 获取token列表
  base.NilResponse TokenList(1: TokenListReq req) (api.post = "/service/token/list")

}