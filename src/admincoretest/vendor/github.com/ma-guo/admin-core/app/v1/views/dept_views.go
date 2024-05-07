package views

// Generated by niuhe.idl

import (
	"fmt"
	"strings"

	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/utils"
	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/admin-core/xorm/services"

	"github.com/ziipin-server/niuhe"
)

type Dept struct {
	_Gen_Dept
}

// 获取部门列表
func (v *Dept) List_GET(c *niuhe.Context, req *protos.V1DeptListReq, rsp *protos.V1DeptListRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	if !strings.Contains(c.Request.URL.RawQuery, "status") {
		req.Status = -1
	}
	rows, err := svc.Dept().Search(req.Keyword, req.Status)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	idsMap := make(map[int64]int64)
	var findChildMenuItem func(parent *protos.V1DeptItem, rows []*models.SysDept)
	// 查找子菜单
	findChildMenuItem = func(parent *protos.V1DeptItem, rows []*models.SysDept) {
		for _, row := range rows {
			if row.ParentId != 0 && row.ParentId == parent.Id {
				if _, ok := idsMap[row.Id]; ok {
					continue
				}
				idsMap[row.Id] = row.Id
				item := row.ToProtos()
				findChildMenuItem(item, rows)
				parent.Children = append(parent.Children, item)
			}
		}
	}
	isInSearch := len(req.Keyword) > 0 || req.Status >= 0
	for _, row := range rows {
		if row.ParentId != 0 && !isInSearch {
			continue
		}
		if _, ok := idsMap[row.Id]; ok {
			continue
		}
		idsMap[row.Id] = row.Id
		item := row.ToProtos()
		findChildMenuItem(item, rows)

		rsp.Items = append(rsp.Items, item)
	}

	if len(rsp.Items) == 0 {
		rsp.Items = make([]*protos.V1DeptItem, 0)
	}
	return nil
}

// 添加部门
func (v *Dept) Add_POST(c *niuhe.Context, req *protos.V1DeptAddReq, rsp *protos.NoneRsp) error {
	auth, err := getAuthInfo(c)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	svc := services.NewSvc()
	defer svc.Close()
	row := &models.SysDept{
		Name:     req.Name,
		ParentId: req.ParentId,
		Sort:     req.Sort,
		Status:   req.Status,
		Deleted:  0,
		TreePath: "0",
		CreateBy: auth.Uid,
		UpdateBy: auth.Uid,
	}
	if req.ParentId > 0 {
		parent, err := svc.Dept().GetById(req.ParentId)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		row.TreePath = fmt.Sprintf("%v,%v", parent.TreePath, req.ParentId)
	}

	_, err = svc.Dept().Add(row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	return nil
}

// 获取部门下拉列表
func (v *Dept) Options_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.V1DeptOptionsRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	rows, err := svc.Dept().GetAll()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	var findChildDeptOption func(parent *protos.LongOptions, rows []*models.SysDept)
	// 查找子部门-选项
	findChildDeptOption = func(parent *protos.LongOptions, rows []*models.SysDept) {
		for _, row := range rows {
			if row.ParentId != 0 && row.ParentId == parent.Value {
				item := &protos.LongOptions{
					Children: make([]*protos.LongOptions, 0),
					Value:    row.Id,
					Label:    row.Name,
				}
				findChildDeptOption(item, rows)
				parent.Children = append(parent.Children, item)
			}
		}
	}
	for _, row := range rows {
		if row.ParentId != 0 {
			continue
		}
		item := &protos.LongOptions{
			Children: make([]*protos.LongOptions, 0),
			Value:    row.Id,
			Label:    row.Name,
		}
		findChildDeptOption(item, rows)

		rsp.Items = append(rsp.Items, item)
	}

	if len(rsp.Items) == 0 {
		rsp.Items = make([]*protos.LongOptions, 0)
	}
	return nil
}

// 修改部门
func (v *Dept) Update_POST(c *niuhe.Context, req *protos.V1DeptUpdateReq, rsp *protos.NoneRsp) error {
	auth, err := getAuthInfo(c)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	svc := services.NewSvc()
	defer svc.Close()
	row, err := svc.Dept().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	row.Name = req.Name
	row.Status = req.Status
	row.Sort = req.Sort
	row.UpdateBy = auth.Uid
	if req.ParentId > 0 {
		parent, err := svc.Dept().GetById(req.ParentId)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		row.TreePath = fmt.Sprintf("%v,%v", parent.TreePath, req.ParentId)
		row.ParentId = req.ParentId
	}

	err = svc.Dept().Update(row.Id, row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	return nil
}

// 获取部门表单数据
func (v *Dept) Form_GET(c *niuhe.Context, req *protos.V1DeptFormReq, rsp *protos.V1DeptFormRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	row, err := svc.Dept().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	rsp.Id = row.Id
	rsp.Name = row.Name
	rsp.ParentId = row.ParentId
	rsp.Sort = row.Sort
	rsp.Status = row.Status

	return nil
}

// 删除部门
func (v *Dept) Delete_POST(c *niuhe.Context, req *protos.V1DeptDeleteReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	ids := utils.SplitInt64(req.Ids)
	for _, id := range ids {
		row, err := svc.Dept().GetById(id)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		err = svc.Dept().Delete(row.Id)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
	}
	return nil
}
func init() {
	GetModule().Register(&Dept{})
}
