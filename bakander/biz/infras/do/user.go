package do

import "kcers-survey/idl_gen/model/user"

type User interface {
	Create(req *user.CreateOrUpdateUserReq) error
	Update(req *user.CreateOrUpdateUserReq) error
	ChangePassword(userId int64, newPassword string) error
	Info(id int64) (info *user.UserInfo, err error)
	List(req *user.UserListReq) (userList []*user.UserInfo, total int, err error)
	UpdateUserStatus(id, status int64) error
	DeleteUser(id int64) error
	SetRole(id int64, roleID []int64) error

	GetUserName(id int64) (name string)
	Login(req *user.LoginReq) (res *user.LoginResp, err error)
}
