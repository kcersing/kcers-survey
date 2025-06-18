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
}

service SysService {


    base.NilResponse RoleList(1: SysListReq req) (api.post = "/api/sys/role/list")

}



