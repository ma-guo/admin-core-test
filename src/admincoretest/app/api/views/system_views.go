package views

// Generated by niuhe.idl

import (
	"admincoretest/app/api/protos"

	"github.com/ziipin-server/niuhe"
)

type System struct {
	_Gen_System
}

// 测试  api
func (v *System) Test_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}
func init() {
	GetModule().Register(&System{})
}
