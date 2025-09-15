package do

import (
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/menu"
)

type Menu interface {
	Create(menuReq *menu.CreateOrUpdateMenuReq) error
	Update(menuReq *menu.CreateOrUpdateMenuReq) error
	Delete(id int64) error
	MenuRole(roleId []int64) (list []*menu.MenuInfo, err error)
	List(req *menu.MenuListReq) (list []*menu.MenuInfo, total int, err error)
	MenuTree(req *menu.MenuListReq) (list []*base.Tree, err error)
	MenuInfo(id int64) (info *menu.MenuInfo, err error)
	//CreateMenuParam(req *menu.MenuParam) error
	//UpdateMenuParam(req *menu.MenuParam) error
	//DeleteMenuParam(menuParamID *int64) error
	//MenuParamListByMenuID(menuID *int64) (list []menu.MenuParam, total int64, err error)
}
