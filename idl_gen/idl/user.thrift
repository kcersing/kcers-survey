namespace go user

include "../base/base.thrift"
include "dictionary.thrift"
struct UserInfo {
	1:i64 id (api.raw = "id")
	2:i64 status (api.raw = "status")
	3:string username (api.raw = "username")
	4:string password (api.raw = "password")
	5:string name (api.raw = "name")
	7:string mobile (api.raw = "mobile")
	8:string avatar (api.raw = "avatar")
	9:string createdAt (api.raw = "createdAt")
	10:string updatedAt  (api.raw = "updatedAt")

	13:string gender (api.raw = "gender")
	15:string birthday (api.raw = "birthday")
    16:string detail (api.raw = "detail")

    21: list<UserRole> userRole (api.raw = "userRole")
    22: list<i64> userRoleIds (api.raw = "userRoleIds")

    23:  optional string email (api.raw = "email")
    25:  optional string wecom (api.raw = "wecom")

}


// login request | 登录参数
struct LoginReq {
    1:  string username (api.raw = "username")
    2:  string password (api.raw = "password")
    3:  string captchaId (api.raw = "captchaId")
    4:  string captcha (api.raw = "captcha")
}
struct LoginResp{
    1:  i64 userId (api.raw = "userId")
    2:  string username (api.raw = "username")
    3:  list<UserRole> userRole (api.raw = "userRole")
    4: list<i64> userRoleIds (api.raw = "userRoleIds")
    5: string roleIdStr (api.raw = "roleIdStr")
}

struct UserRole{
    1:optional  string name="" (api.raw = "name")
    2:optional  string value="" (api.raw = "value")
    3:optional  i64 id=0 (api.raw = "id")
}

// register request | 注册参数
struct RegisterReq {
    1:  optional string username (api.raw = "username")
    2:  optional string password (api.raw = "password")
    3:  optional string captchaId (api.raw = "captchaId")
    4:  optional string captcha (api.raw = "captcha")
    5:  optional string email (api.raw = "email")
}

// change user's password request | 修改密码请求参数
struct ChangePasswordReq {
    1: i64 userId (api.raw = "userId")
    2: string newPassword (api.raw = "newPassword")
}

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateUserReq {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string avatar="" (api.raw = "avatar")
    4:  optional string mobile="" (api.raw = "mobile")
    6:  optional i64 status=1 (api.raw = "status")
    7:  optional string name="" (api.raw = "name")
    8:  optional string gender="" (api.raw = "gender")
    9:  optional list<i64> roleId=0 (api.raw = "roleId")
    10: optional i64 createId=0 (api.raw = "createId")
    12: optional string password="" (api.raw = "password")
    13:  optional string username="" (api.raw = "username")
    14:  optional list<string> functions="" (api.raw = "functions")
    15:  optional string detail="" (api.raw = "detail")

	24:string birthday="" (api.raw = "birthday")
    23:  optional string email="" (api.raw = "email")
    25:  optional string wecom="" (api.raw = "wecom")

}


// Get user list request | 获取用户列表请求参数
struct UserListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    4:  optional string name="" (api.raw = "name")
    6:  optional string mobile="" (api.raw = "mobile")
    7:  optional i64 roleId=0 (api.raw = "roleId")
    8:  optional i64 status=0 (api.raw = "status")

}

struct SetUserRole{
    1:  optional i64 userId=0 (api.raw = "userId")
    2:  optional list<i64> roleId=0 (api.raw = "roleId")
}



service UserService {
  // 登录
//  base.NilResponse Login(1: LoginReq req) (api.post = "/service/login")

  // 注册
  //base.NilResponse Register(1: RegisterReq req) (api.post = "/service/register")

  /**修改密码*/
 base.NilResponse ChangePassword(1: ChangePasswordReq req) (api.post = "/service/user/change-password")

  /**新增用户*/
  base.NilResponse CreateUser(1: CreateOrUpdateUserReq req) (api.post = "/service/user/create")

  /**更新用户*/
  base.NilResponse UpdateUser(1: CreateOrUpdateUserReq req) (api.post = "/service/user/update")

  /**获取用户基本信息*/
  base.NilResponse UserInfo(1: base.IDReq req)  (api.get = "/service/user/info")

  /**获取用户列表*/
  base.NilResponse UserList(1: UserListReq req) (api.post = "/service/user/list")

  /**删除用户信息*/
  base.NilResponse DeleteUser(1: base.IDReq req) (api.post = "/service/user")

  /**更新用户状态*/
  base.NilResponse UpdateUserStatus(1: base.StatusCodeReq req) (api.post = "/service/user/status")

  /**设置用户角色*/
  base.NilResponse SetUserRole(1: SetUserRole req) (api.post = "/service/user/set-role")


}
