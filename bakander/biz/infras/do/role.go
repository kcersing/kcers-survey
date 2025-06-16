package do

import "kcers-survey/idl_gen/model/auth"

type Role interface {
	Create(req *auth.RoleInfo) error
	Update(req *auth.RoleInfo) error
	Delete(id int64) error
	RoleInfoById(id int64) (roleInfo *auth.RoleInfo, err error)
	List(req *auth.RoleListReq) (roleInfoList []*auth.RoleInfo, total int, err error)
	UpdateStatus(id, status int64) error
}
