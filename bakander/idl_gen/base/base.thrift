namespace go base
//bool: 布尔值，对应Java中的boolean，
//byte: 有符号字节，对应Java中的byte，对应MySQL的tinyint
//i16: 16位有符号整型，对应Java中的short，对应MySQL的smallint
//i32: 32位有符号整型，对应Java中的int，对应MySQL的int
//i64: 64位有符号整型，对应Java中的long，对应MySQL的bigint
//double: 64位浮点型，对应Java中的double
//string: 字符串，对应Java中的String
//binary: Blob 类型，对应Java中的byte[]

struct BaseResp {
    1: string message  ="" (api.raw = "message")
    2: i32 code = 0  (api.raw = "code")
    3: optional map<string, string> Extra={} (api.raw = "extra")
}

struct BaseResponse {
    1: i64 code  =0 (api.raw = "code")
    2: string message="" (api.raw = "message")
}

struct NilResponse {

}

struct Empty {

}

struct List {
    1: optional i64 id= 0  (api.raw ="id")
    2: optional string name ="" (api.raw ="name")
}
struct IDReq{
    1: i64 id,
}
struct Ids{
    1: list<i64> ids,
}

struct PageInfoReq{
    1: i64 page=1(api.raw = "page")
    2: i64 pageSize=100 (api.raw = "pageSize")
}

struct StatusCodeReq {
    1: i64 id =0 (api.raw = "id")
    2: i64 status=0 (api.raw = "status")
}

struct Tree  {
 1:	string title ="" (api.raw = "title")
 2:	string value="" (api.raw = "value")
 3:	string key="" (api.raw = "key")
 4:	string method="" (api.raw = "method")
 5:	list<Tree> children={} (api.raw = "children")
}

struct Seat{
    /**编号*/
    1: optional i64 num = 0 (api.raw = "num" )
    2: optional i64 x =0 (api.raw = "x" )
    3: optional i64 y =0 (api.raw = "y" )
}












enum Err {
    Success            = 0,
    NoRoute            = 1,
    NoMethod           = 2,
    BadRequest         = 10000,
    ParamsErr          = 10001,
    AuthorizeFail      = 10002,
    TooManyRequest     = 10003,
    ServiceErr         = 20000,
    RPCUserSrvErr      = 30000,
    UserSrvErr         = 30001,
    RPCBlobSrvErr      = 40000,
    BlobSrvErr         = 40001,
    RPCCarSrvErr       = 50000,
    CarSrvErr          = 50001,
    RPCProfileSrvErr   = 60000,
    ProfileSrvErr      = 60001,
    RPCTripSrvErr      = 70000,
    TripSrvErr         = 70001,
    RecordNotFound     = 80000,
    RecordAlreadyExist = 80001,
    DirtyData          = 80003,
}