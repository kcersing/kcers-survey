namespace go auth

include "../base/base.thrift"
include "menu.thrift"
// authorization service
service AuthService {
  // Create role information | 创建角色
  base.NilResponse CreateRole(1: RoleInfo req) (api.post = "/service/role/create")

  // Update role information | 更新角色
  base.NilResponse UpdateRole(1: RoleInfo req) (api.post = "/service/role/update")

  // Delete role information | 删除角色信息
  base.NilResponse DeleteRole(1: base.IDReq req) (api.post = "/service/role/del")

  // Get role information | 获取角色信息
  base.NilResponse RoleByID(1: base.IDReq req) (api.get = "/service/role")

  // 创建菜单权限
  base.NilResponse CreateMenuAuth(1: MenuAuthInfoReq req) (api.post = "/service/auth/menu/create")

  // 更新菜单权限
  base.NilResponse UpdateMenuAuth(1: MenuAuthInfoReq req) (api.post = "/service/auth/menu/update")


  // Get role list | 获取角色列表
  base.NilResponse RoleList(1: base.PageInfoReq req) (api.get = "/service/role/list")

  base.NilResponse RoleTree(1: base.PageInfoReq req) (api.get = "/service/role/tree")

  // Set role status | 设置角色状态, 启用1/禁用0
  base.NilResponse UpdateRoleStatus(1: base.StatusCodeReq req) (api.post = "/service/role/status")

  // 创建API权限
  base.NilResponse CreateAuth(1: CreateOrUpdateApiAuthReq req) (api.post = "/service/auth/api/create")

  // 更新API权限
  base.NilResponse UpdateAuth(1: CreateOrUpdateApiAuthReq req) (api.post = "/service/auth/api/update")

  // 获取角色api权限列表
  base.NilResponse ApiAuth(1: base.IDReq req) (api.post = "/service/auth/api/role")


}

// 菜单授权请求数据
struct MenuAuthInfoReq {
    1: i64 roleId (api.raw = "roleId")
    2: list<i64> menuIds (api.raw = "menuIds")
}
struct RoleMenu {
    1: list<i64> menuIds (api.raw = "menuIds")
    2: list<menu.MenuInfo> menuInfo (api.raw = "menuInfo")
}

struct RoleListReq{
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    12:optional i64 venueId=0 (api.raw = "venueId")
}
//创建或更新角色信息参数
struct RoleInfo {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional string value="" (api.raw = "value")
    4:  optional string defaultRouter="" (api.raw = "defaultRouter")
    5:  optional i64 status=1 (api.raw = "status")
    6:  optional string remark=""  (api.raw = "remark")
    7:  optional i64 orderNo=0 (api.raw = "orderNo")
    8:  optional string createdAt="" (api.raw = "createdAt")
    9:  optional string updatedAt="" (api.raw = "updatedAt")
    10: optional list<i64> menus={} (api.raw = "menus")
    11: optional list<i64> apis={} (api.raw = "apis")

    12:optional i64 venueId=0 (api.raw = "venueId")
}


// authorization message
// API授权数据
struct ApiAuthInfo {
    1: string path (api.raw = "path")
    2: string method (api.raw = "method")
}

// 创建或更新API授权信息
struct CreateOrUpdateApiAuthReq {
    1: i64 roleId (api.raw = "roleId")
//    2:  ApiAuthInfo data (api.raw = "api_auth_info")
    2: list<i64> apis (api.raw = "apis")
}

