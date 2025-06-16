package do

import (
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/menu"
)

type Menu interface {
	Create(menuReq *menu.MenuInfo) error
	Update(menuReq *menu.MenuInfo) error
	Delete(id int64) error
	ListByRole(roleId int64) (list []*menu.MenuInfoTree, total int64, err error)
	List(req *menu.MenuListReq) (list []*menu.MenuInfoTree, total int, err error)
	MenuTree(req *menu.MenuListReq) (list []*base.Tree, total int, err error)

	//CreateMenuParam(req *menu.MenuParam) error
	//UpdateMenuParam(req *menu.MenuParam) error
	//DeleteMenuParam(menuParamID *int64) error
	//MenuParamListByMenuID(menuID *int64) (list []menu.MenuParam, total int64, err error)
}
