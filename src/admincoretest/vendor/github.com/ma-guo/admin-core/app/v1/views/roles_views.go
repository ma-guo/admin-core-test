package views

// Generated by niuhe.idl

import (
	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/utils"
	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/admin-core/xorm/services"

	"github.com/ma-guo/niuhe"
)

type Roles struct {
	_Gen_Roles
}

// 角色分页列表
func (v *Roles) Page_GET(c *niuhe.Context, req *protos.V1RolesPageReq, rsp *protos.V1RolesPageRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	roles, total, err := svc.Role().GetPage(req.Keywords, req.PageNum, req.PageSize)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Total = total
	for _, role := range roles {
		rsp.List = append(rsp.List, &protos.V1RolesItem{
			Id:         role.Id,
			Name:       role.Name,
			Status:     role.Status,
			Sort:       role.Sort,
			Code:       role.Code,
			CreateTime: role.CreateTime.Format(consts.FullTimeLayout),
			UpdateTime: role.UpdateTime.Format(consts.FullTimeLayout),
		})
	}
	if len(rsp.List) == 0 {
		rsp.List = make([]*protos.V1RolesItem, 0)
	}
	return nil
}

// 角色下拉列表
func (v *Roles) Options_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.V1RolesOptiosRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	roles, err := svc.Role().GetAll()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	for _, role := range roles {
		if role.Deleted {
			continue
		}
		rsp.Items = append(rsp.Items, &protos.LongOptionItem{
			Value: role.Id,
			Label: role.Name,
		})
	}
	if len(rsp.Items) == 0 {
		rsp.Items = make([]*protos.LongOptionItem, 0)
	}
	return nil
}

// 新增角色
func (v *Roles) Add_POST(c *niuhe.Context, req *protos.V1RolesAddReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()

	row := &models.SysRole{
		Id:        req.Id,
		Name:      req.Name,
		Code:      req.Code,
		Sort:      req.Sort,
		Status:    req.Status,
		DataScope: req.DataScope,
	}
	row.Status = req.Status
	_, err := svc.Role().Add(row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	return nil
}

// 角色表单数据
func (v *Roles) Form_GET(c *niuhe.Context, req *protos.V1RolesFormReq, rsp *protos.V1RolesFormRsp) error {
	svc := services.NewSvc()
	defer svc.Close()

	row, err := svc.Role().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if row.Deleted {
		return niuhe.NewCommError(-1, "角色已被删除")
	}

	rsp.Id = row.Id
	rsp.Name = row.Name
	rsp.Code = row.Code
	rsp.Sort = row.Sort
	rsp.Status = row.Status
	rsp.DataScope = row.DataScope
	return nil
}

// 修改角色信息
func (v *Roles) Update_POST(c *niuhe.Context, req *protos.V1RolesUpdateReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()

	row, err := svc.Role().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	row.Name = req.Name
	row.Code = req.Code
	row.Status = req.Status
	row.Sort = req.Sort
	row.DataScope = req.DataScope

	err = svc.Role().Update(row.Id, row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}

// 删除角色
func (v *Roles) Delete_POST(c *niuhe.Context, req *protos.V1RolesDeleteReq, rsp *protos.NoneRsp) error {
	ids := utils.SplitInt64(req.Ids)

	svc := services.NewSvc()
	defer svc.Close()
	for _, uid := range ids {
		row, err := svc.Role().GetById(uid)
		if err != nil {
			niuhe.LogInfo("%v", err)
			continue
		}
		// 将用户标记为删除
		if !row.Deleted {
			row.Deleted = true
			err = svc.Role().Update(row.Id, row)
			if err != nil {
				niuhe.LogInfo("%v", err)
				return err
			}
		}
	}
	return nil
}

// 获取角色的菜单ID集合
func (v *Roles) Menuids_GET(c *niuhe.Context, req *protos.V1RolesMenuIdsReq, rsp *protos.V1RolesMenuIdsRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	rows, err := svc.RoleMenu().GetRoleMenus(req.RoleId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	for _, row := range rows {
		rsp.Items = append(rsp.Items, row.MenuId)
	}
	if len(rsp.Items) == 0 {
		rsp.Items = make([]int64, 0)
	}
	return nil
}

// 分配菜单权限给角色
func (v *Roles) Menus_POST(c *niuhe.Context, req *protos.V1RolesMenusReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	err := svc.RoleMenu().Update(req.RoleId, req.MenuIds...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}

// 修改角色状态
func (v *Roles) Status_POST(c *niuhe.Context, req *protos.V1RolesStatusReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	row, err := svc.Role().GetById(req.RoleId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if row.Status == req.Status {
		return nil
	}
	row.Status = req.Status
	err = svc.Role().Update(row.Id, row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	return nil
}
func init() {
	GetModule().Register(&Roles{})
}
