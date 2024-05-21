package views

// Generated by niuhe.idl

import (
	"github.com/ma-guo/admin-core/app/v1/protos"

	"github.com/ma-guo/niuhe"
)

type _Gen_Menus struct{}

// 菜单列表
func (v *_Gen_Menus) List_GET(c *niuhe.Context, req *protos.V1MenusListReq, rsp *protos.V1MenusListRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 删除菜单
func (v *_Gen_Menus) Delete_POST(c *niuhe.Context, req *protos.V1MenusDeleteReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 新增菜单
func (v *_Gen_Menus) Add_POST(c *niuhe.Context, req *protos.V1MenusAddReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 菜单下拉列表
func (v *_Gen_Menus) Options_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.V1MenusOptionsRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 路由列表
func (v *_Gen_Menus) Routes_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.MenuRouteRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 菜单表单
func (v *_Gen_Menus) Form_GET(c *niuhe.Context, req *protos.V1MenusFormReq, rsp *protos.V1MenusFormRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 修改菜单
func (v *_Gen_Menus) Update_POST(c *niuhe.Context, req *protos.V1MenusUpdateReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 修改菜单状态
func (v *_Gen_Menus) Status_POST(c *niuhe.Context, req *protos.V1MenusStatusReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}
