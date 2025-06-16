package do

import "kcers-survey/idl_gen/model/sys"

type Sys interface {
	RoleList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
}
