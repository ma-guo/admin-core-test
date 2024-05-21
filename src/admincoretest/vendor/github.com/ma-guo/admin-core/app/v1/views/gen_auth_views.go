package views

// Generated by niuhe.idl

import (
	"github.com/ma-guo/admin-core/app/v1/protos"

	"github.com/ma-guo/niuhe"
)

type _Gen_Auth struct{}

// 登录
func (v *_Gen_Auth) Login_POST(c *niuhe.Context, req *protos.AuthLoginReq, rsp *protos.AauthLoginRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 登出
func (v *_Gen_Auth) Logout_POST(c *niuhe.Context, req *protos.NoneReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 获取验证码
func (v *_Gen_Auth) Captcha_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.AuthCaptchaRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}
