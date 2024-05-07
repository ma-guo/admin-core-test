package boot

// Generated by niuhe.idl

import (
	apiViews "github.com/ma-guo/admin-core/app/v1/views"
	"github.com/ma-guo/admin-core/config"

	"github.com/ziipin-server/niuhe"
)

type AdminBoot struct {
	protocol niuhe.IApiProtocol
}

func (AdminBoot) LoadConfig(path string) error {
	return config.LoadConfig(path)
}

func (AdminBoot) InitConfig(conf config.AdminConfig) error {
	config.InitConfig(conf)
	return nil
}

func (AdminBoot) BeforeBoot(svr *niuhe.Server) {}

func (admin *AdminBoot) RegisterModules(svr *niuhe.Server) {
	if admin.protocol != nil {
		apiViews.SetProtocol(admin.protocol)
	}
	svr.RegisterModule(apiViews.GetModule())
}

func (AdminBoot) Serve(svr *niuhe.Server) {
	svr.Serve(config.Config.ServerAddr)
}

// 设置协议, 如果需要更改 请求和返回结构, 可调用本方法处理, 具体实现参考 app/v1/views/init_protocol.go
func (admin *AdminBoot) SetProtocol(protocol niuhe.IApiProtocol) {
	admin.protocol = protocol
}

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	if len(os.Args) < 2 {
// 		niuhe.LogInfo("usage: %s <config-path>", os.Args[0])
// 		return
// 	}
// 	path := os.Args[1]
// 	boot := AdminBoot{}
// 	if err := boot.LoadConfig(path); err != nil {
// 		panic(err)
// 	}
// 	svr := niuhe.NewServer()
// 	boot.BeforeBoot(svr)
// 	boot.RegisterModules(svr)
// 	boot.Serve(svr)
// }
