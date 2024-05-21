package views

// Generated by niuhe.idl

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/utils"
	"github.com/ma-guo/admin-core/utils/password"
	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/admin-core/xorm/services"

	"github.com/ma-guo/niuhe"
	"github.com/xuri/excelize/v2"
)

type Users struct {
	_Gen_Users
	salt string // 密码加盐
}

// 获取当前登录用户信息
func (v *Users) Me_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.V1UsersMeRsp) error {
	auth, err := getAuthInfo(c)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	svc := services.NewSvc()
	defer svc.Close()
	user, err := svc.User().GetById(auth.Uid)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err

	}
	if user.Deleted {
		return niuhe.NewCommError(consts.AuthError, "账号已被禁用")
	}
	rsp.UserId = user.Id
	rsp.Avatar = user.Avatar
	rsp.Username = user.Username
	rsp.Nickname = user.Nickname
	roles, err := svc.Role().GetByUserId(user.Id)
	roleIds := make([]int64, 0)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Roles = make([]string, 0)
	for _, role := range roles {
		rsp.Roles = append(rsp.Roles, role.Code)
		roleIds = append(roleIds, role.Id)
	}
	if len(roleIds) > 0 {
		meus, err := svc.RoleMenu().GetRoleMenus(roleIds...)

		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		ids := make([]int64, 0)
		for _, meu := range meus {
			ids = append(ids, meu.MenuId)
		}
		items, err := svc.Menu().GetByIds(ids...)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		for _, item := range items {
			if item.Perm != "" {
				rsp.Perms = append(rsp.Perms, item.Perm)
			}
		}
	}
	if len(rsp.Perms) == 0 {
		rsp.Perms = make([]string, 0)
	}
	return nil
}

// 用户分页列表
func (v *Users) Page_GET(c *niuhe.Context, req *protos.V1UsersPageReq, rsp *protos.V1UsersPageRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	users, total, err := svc.User().GetPage(req, c.Request.URL.RawQuery)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Total = total
	rsp.List = make([]*protos.V1UserPageItem, 0)
	ids := make([]int64, 0)
	for _, user := range users {
		ids = append(ids, user.DeptId)
	}
	depts, err := svc.Dept().GetByIds(ids...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	for _, user := range users {
		item := &protos.V1UserPageItem{
			Id:          user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Mobile:      user.Mobile,
			GenderLabel: user.GenderLabel(),
			Avatar:      user.Avatar,
			Email:       user.Email,
			Status:      user.IntStatus(),
			DeptName:    "",
			RoleNames:   "",
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if dept, has := depts[user.DeptId]; has {
			item.DeptName = dept.Name
		}
		rsp.List = append(rsp.List, item)
	}
	return nil
}

// 新增用户
func (v *Users) Add_POST(c *niuhe.Context, req *protos.V1UsersAddReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	_, has, err := svc.User().GetByUsername(req.Username)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if has {
		return niuhe.NewCommError(-1, "用户名已存在")
	}
	pwd := password.NewPassword(req.Password)
	pwd.SetSalt(v.salt)
	user := &models.SysUser{
		Username: req.Username,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Password: pwd.Hash(),
		DeptId:   req.DeptId,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Deleted:  false,
	}
	user.SetStatus(req.Status)
	// 写入用户数据
	_, err = svc.User().Insert(user)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	// 写入角色数据
	for _, rid := range req.RoleIds {
		role := &models.SysUserRole{
			UserId: user.Id,
			RoleId: rid,
		}
		_, err := svc.UserRole().Add(role)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
	}
	return nil
}

// 用户表单
func (v *Users) Form_GET(c *niuhe.Context, req *protos.V1UserFormReq, rsp *protos.V1UserFormRsp) error {

	svc := services.NewSvc()
	defer svc.Close()
	user, err := svc.User().GetById(req.UserId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err

	}
	rsp.Id = user.Id
	rsp.Username = user.Username
	rsp.Nickname = user.Nickname
	rsp.Mobile = user.Mobile
	rsp.Gender = user.Gender
	rsp.Avatar = user.Avatar
	rsp.Email = user.Email
	rsp.Status = user.IntStatus()
	rsp.DeptId = user.DeptId

	roles, err := svc.Role().GetByUserId(user.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.RoleIds = make([]int64, 0)
	for _, role := range roles {
		rsp.RoleIds = append(rsp.RoleIds, role.Id)
	}
	return nil
}

// 修改用户
func (v *Users) Update_POST(c *niuhe.Context, req *protos.V1UserUpdateReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	user, err := svc.User().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if user.Deleted {
		return niuhe.NewCommError(-1, "账号已被禁用")
	}
	user.Nickname = req.Nickname
	user.Mobile = req.Mobile
	user.Gender = req.Gender
	user.Avatar = req.Avatar
	user.Email = req.Email
	user.SetStatus(req.Status)
	user.DeptId = req.DeptId
	err = svc.User().Update(user.Id, user)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	err = svc.UserRole().Update(user.Id, req.RoleIds...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}

// 删除用户
func (v *Users) Delete_POST(c *niuhe.Context, req *protos.V1UsersDeleteReq, rsp *protos.NoneRsp) error {
	ids := utils.SplitInt64(req.Ids)

	svc := services.NewSvc()
	defer svc.Close()
	for _, uid := range ids {
		row, err := svc.User().GetById(uid)
		if err != nil {
			niuhe.LogInfo("%v", err)
			continue
		}
		// 将用户标记为删除
		if !row.Deleted {
			row.Deleted = true
			err = svc.User().Update(row.Id, row)
			if err != nil {
				niuhe.LogInfo("%v", err)
				return err
			}
		}
	}
	return nil
}

// 到出用户
func (v *Users) Export_GET(c *niuhe.Context, req *protos.V1UsersPageReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	users, _, err := svc.User().GetPage(req, c.Request.URL.RawQuery)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	deps, err := svc.Dept().GetAll()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	deptMap := make(map[int64]string, 0)
	for _, dep := range deps {
		deptMap[dep.Id] = dep.Name
	}
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	sheetName := consts.Sheet1
	index, err := xlsx.NewSheet(sheetName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	xlsx.SetCellValue(sheetName, "A1", "用户名")
	xlsx.SetCellValue(sheetName, "B1", "用户昵称")
	xlsx.SetCellValue(sheetName, "C1", "部门")
	xlsx.SetCellValue(sheetName, "D1", "性别")
	xlsx.SetCellValue(sheetName, "E1", "手机号码")
	xlsx.SetCellValue(sheetName, "F1", "邮箱")
	xlsx.SetCellValue(sheetName, "G1", "创建时间")

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	rowIdx := 1
	for _, user := range users {
		rowIdx++
		rowNum := strconv.Itoa(rowIdx)
		xlsx.SetCellValue(sheetName, "A"+rowNum, user.Username)
		xlsx.SetCellValue(sheetName, "B"+rowNum, user.Nickname)
		if depName, has := deptMap[user.DeptId]; has {
			xlsx.SetCellValue(sheetName, "C"+rowNum, depName)
		}
		// xlsx.SetCellValue(sheetName, "C"+rowNum, user.DeptId)
		xlsx.SetCellValue(sheetName, "D"+rowNum, user.GenderLabel())
		xlsx.SetCellValue(sheetName, "E"+rowNum, user.Mobile)
		xlsx.SetCellValue(sheetName, "F"+rowNum, user.Email)
		xlsx.SetCellValue(sheetName, "G"+rowNum, user.CreateTime.Format(consts.FullTimeLayout))
	}
	fileName := time.Now().Format("20060102150405") + ".xlsx"
	fullName, err := utils.GetTmpFileName(fileName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	niuhe.LogInfo("export %v", fullName)
	if err = xlsx.SaveAs(fullName); err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=用户列表.xlsx")
	// 指定文件格式
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(fullName)
	// 删除文件
	os.Remove(fullName)
	return niuhe.NewCommError(consts.CodeNoCommRsp, "已经处理了返回")
}

// 用户导入模板下载
func (v *Users) Template_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.NoneRsp) error {
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	sheetName := consts.Sheet1
	index, err := xlsx.NewSheet(sheetName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	xlsx.SetCellValue(sheetName, "A1", "用户名")
	xlsx.SetCellValue(sheetName, "B1", "用户昵称")
	xlsx.SetCellValue(sheetName, "C1", "性别")
	xlsx.SetCellValue(sheetName, "D1", "手机号码")
	xlsx.SetCellValue(sheetName, "E1", "邮箱")
	xlsx.SetCellValue(sheetName, "F1", "角色")

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)

	fileName := time.Now().Format("20060102150405") + ".xlsx"
	fullName, err := utils.GetTmpFileName(fileName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	niuhe.LogInfo("template %v", fullName)
	if err = xlsx.SaveAs(fullName); err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=用户导入模板.xlsx")
	// 指定文件格式
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(fullName)
	// 删除文件
	os.Remove(fullName)
	return niuhe.NewCommError(consts.CodeNoCommRsp, "已经处理了返回")
}

// 用户导入
func (v *Users) Import_POST(c *niuhe.Context) {
	deptIdStr, has := c.GetQuery("deptId")
	if !has {
		rspWithError(c, -1, "部门ID不能为空")
		return
	}
	deptId, err := strconv.ParseInt(deptIdStr, 10, 64)
	if err != nil {
		rspWithError(c, -1, err.Error())
		return
	}
	svc := services.NewSvc()
	defer svc.Close()

	_, err = svc.Dept().GetById(deptId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		rspWithError(c, -1, err.Error())
		return
	}
	file, er := c.FormFile("file")
	if er != nil {
		niuhe.LogInfo("%v", er)
		rspWithError(c, -1, er.Error())
		return
	}
	fileName := time.Now().Format("20060102150405") + ".xlsx"
	fullName, err := utils.GetTmpFileName(fileName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		rspWithError(c, -1, err.Error())
		return
	}

	if err = c.SaveUploadedFile(file, fullName); err != nil {
		niuhe.LogInfo("%v", err)
		rspWithError(c, -1, err.Error())
		return
	}

	defer os.Remove(fullName)

	xlsx, err := excelize.OpenFile(fullName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		rspWithError(c, -1, err.Error())
		return
	}
	defer xlsx.Close()

	rows, err := xlsx.GetRows(consts.Sheet1)
	if err != nil {
		niuhe.LogInfo("%v", err)
		rspWithError(c, -1, err.Error())
		return
	}
	users := make([]*models.SysUser, 0)
	safeGet := func(row []string, idx int) string {
		if len(row) > idx {
			return strings.Trim(row[idx], " ")
		}
		return ""
	}
	// xlsx.SetCellValue(sheetName, "A1", "用户名")
	// xlsx.SetCellValue(sheetName, "B1", "用户昵称")
	// xlsx.SetCellValue(sheetName, "C1", "性别")
	// xlsx.SetCellValue(sheetName, "D1", "手机号码")
	// xlsx.SetCellValue(sheetName, "E1", "邮箱")
	// xlsx.SetCellValue(sheetName, "F1", "角色")

	roleUserMap := make(map[string]string, 0)
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		user := &models.SysUser{
			Username: safeGet(row, 0),
			Nickname: safeGet(row, 1),
			Gender:   0,
			Mobile:   safeGet(row, 3),
			Email:    safeGet(row, 4),
			DeptId:   deptId,
			Avatar:   "",
			Deleted:  false,
		}

		_, has, err := svc.User().GetByName(user.Username, user.Nickname)
		if !has && err == nil {
			// 不存在且有效时才添加
			users = append(users, user)
			role := safeGet(row, 5)
			if role != "" {
				roleUserMap[user.Username] = role
			}
		}
	}
	for _, user := range users {
		_, err = svc.User().Insert(user)
		if err == nil && user.Id > 0 {
			if roleName, has := roleUserMap[user.Username]; has {
				role, err := svc.Role().GetByName(roleName)
				if err == nil {
					_, err = svc.UserRole().Add(&models.SysUserRole{
						UserId: user.Id,
						RoleId: role.Id,
					})
					if err != nil {
						niuhe.LogInfo("%v", err)
					}
				}

			}
		}
	}

	rspWithSuccess(c, "导入成功")
}

// 修改密码
func (v *Users) Password_POST(c *niuhe.Context, req *protos.V1UsersPasswordReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	user, err := svc.User().GetById(req.UserId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if user.Deleted {
		return niuhe.NewCommError(-1, "账号已被禁用")
	}
	pwd := password.NewPassword(req.Password)
	pwd.SetSalt(v.salt)

	if pwd.Compare(user.Password) {
		return nil
	}
	// 更新用户信息
	user.Password = pwd.Hash()
	err = svc.User().Update(user.Id, user)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}

// 修改用户状态
func (v *Users) Status_POST(c *niuhe.Context, req *protos.V1UsersStatusReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	user, err := svc.User().GetById(req.UserId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	if user.IntStatus() == req.Status {
		return nil
	}
	// 更新用户信息
	user.SetStatus(req.Status)
	err = svc.User().Update(user.Id, user)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}
func init() {
	GetModule().Register(&Users{
		salt: "admincore",
	})
}
