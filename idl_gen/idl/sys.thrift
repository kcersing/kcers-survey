namespace go sys

include "../base/base.thrift"

struct SysList  {
	 1: optional i64 id=0 (api.raw = "id")
	 2: optional string name="" (api.raw = "name")
	 3: optional string key ="" (api.raw = "key")
}


struct SysListReq {
    1: optional string name="" (api.raw = "name")
    2: optional i64 dictionaryId=0 (api.raw = "dictionaryId" )
    3: optional string type="" (api.raw = "type" )
    4: optional i64 id=0 (api.raw = "id")
}

service SysService {


    base.NilResponse RoleList(1: SysListReq req) (api.post = "/service/sys/role/list")

    base.NilResponse Area(1: SysListReq req) (api.get = "/service/sys/area")
    base.NilResponse City(1: SysListReq req) (api.get = "/service/sys/city")
}



