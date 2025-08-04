# 本 demo 说明

本文以 [admin-core-test](https://github.com/ma-guo/admin-core-test) 为例，引用的项目有 [admin-core](https://github.com/ma-guo/admin-core), [vue3-element-admin](https://github.com/ma-guo/vue3-element-admin)。大部分内容与 [admin-core-test/README.md](https://github.com/ma-guo/admin-core-test/blob/main/README.MD) 相同



`admin-core-test` 演示使用 vscode 插件 `niuhe` 接入项目 [admin-core](https://github.com/ma-guo/admin-core/) 例子。前端搭配项目为 [vue3-element-admin](https://github.com/ma-guo/vue3-element-admin)
## 1. 定义 niuhe 文件
 在 `niuhe/all.niuhe` 文件下定义一个 api, 然后点击 ![entry](http://niuhe.zuxing.net/assets/niuhedoc05.png) 按钮 生成 `go` 项目代码
```python
#app=admincoretest

class NoneReq(Message):
    pass

class NoneRsp(Message):
    pass

with services():
    GET('测试 api', '/api/system/test/', NoneReq, wraps(NoneRsp))
```
> 更多 niuhe 定义用例可参考: https://github.com/ma-guo/admin-core-niuhe/tree/main/niuhe
## 2. 修改配置文件
在配置 [conf/admincoretest.yaml](https://github.com/ma-guo/admin-core-test/blob/main/conf/admincoretest.yaml) 文件中添加第8和第9行并修改第7行的数据库连接信息, 项目运行后会自动在对应的库中创建对应的表信息，但不会创建数据库，数据库需要提前进行创建
> 管理后台的文件服务可能使用到 `host` 字段, 如用到文件存储服务，请在配置文件中配置 `host` 信息，格式为 `http(s)://...`

## 3. 接入 AdminBoot
在 [src/admincoretest/main.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/main.go) 中接入 `AdminBoot`
```go
 
 func main() {
	boot := BaseBoot{}
	if err := boot.LoadConfig(); err != nil {
		panic(err)
	}
	svr := niuhe.NewServer()
	boot.BeforeBoot(svr)
	boot.RegisterModules(svr)
	boot.Serve(svr)
}
```
上述 `main` 方式是自定生成的, 将 `AdminBoot` 引入, 初始化 config(`LoadConfig`) 并 `RegisterModules` 即可
```go
// adminBoot "github.com/ma-guo/admin-core/boot" // import
 func main() {
	boot := BaseBoot{}
	if err := boot.LoadConfig(); err != nil {
		panic(err)
	}
	admin := adminBoot.AdminBoot{}
	if err := admin.LoadConfig(os.Args[1]); err != nil {
		niuhe.LogInfo("admin config error: %v", err)
		return
	}
	svr := niuhe.NewServer()
	boot.BeforeBoot(svr)
	admin.RegisterModules(svr)
	boot.RegisterModules(svr)
	boot.Serve(svr)
}
```
进过上面三步即可接入 `admin-core` 项目, 在 vscode 下可愉快地使用 `niuhe` 插件加速您的项目开发了。

## 4. 登录验证 - 可选
1. 自定义方法需要加入 `Bearea` 认证的请参考 [src/admincoretest/views/init.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/app/api/views/init.go) 中的使用
2. 需要修改 `Bearea` 认证盐请在[配置文件](https://github.com/ma-guo/admin-core-test/blob/main/conf/admincoretest.yaml)中配置 `secretkey` 值
3. 登录失败时会将请求的密码生成的加密字符串 log 出来，首次登录时将字符串替换到表中对应用户的密码字段即可
> 本步骤为可选, 如不需要将新加的方法加入到登录认证中，则请略过本步骤

## 5. 使用 API 级权限校验 - 可选
API 级校验接入后需要在后台进行配置对应权限, 同时服务端也需要进行一些简单接入。路由信息在运行 niuhe 插件生成的 API 后会自动添加到数据表中。使用权限校验需要两步处理
### 5.1 在 langs 中添加 route
需在 `niuhe/.config.json5` 文件的 `langs` 中添加 `route`, 配置后会在 `src/{app}/app/api/protos/gen_routes.go` 中生成项目定义的路由信息
```json5
{
	langs: ["go", "ts", "route"],
}
```
### 5.2 将路由信息接入加入到 protocol 中
接入代码参考 [views/init.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/app/api/views/init.go), 以下代码为 init.go 中的 18-26行
```go
routes := []*niuhe.RouteItem{}
for _, route := range protos.RouteItems {
        routes = append(routes, &niuhe.RouteItem{
                Method: route.Method,
                Path:   route.Path,
                Name:   route.Name,
        })
}
coreViews.GetProtocol().AddRoute("", routes)
```

## 6 生成 swagger.json API 文档 - 可选
需要生成 api 文档需要在 langs 中添加 `docs`, 添加后将在会生成文档配置文件 `niuhe/docs.json5`(存在则不生成) 和 `docs/swagger.json`(每次都会复写)。

生成的 swagger.json 文档遵循 swagger 2.0 协议，可直接导入 `postman` 和  `ApiFox(推荐)` 中使用。也可以做成 api 返回 json 作为 URL 导入到 [`ApiFox`](https://apifox.com/) 中保持实时更新, 参考 [views/system_views.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/app/api/views/system_views.go) `Docs_GET` 方法




# 参考 niuhe/.config,json5 完整配置
> .config 为本地定义文件, 一般不需要跟随 git 版本提交(多成员开发时每个人的环境可能不同)
> 
> 配置项不能添加注释, 下列说明中配置项后面的 // 注释为实例

```json5
{
	app: "", // 为生成的代码中的 app 名称, 默认为空字符串, 空字符串时同 #app=app_name
	gomod: "", // 为生成的代码中的 gomod 名称, 默认为为空字符串, 空字符串时同 app
	langs: ["ts", "docs", "route"], // 为生成的语言类型, 默认为 "go"。 同时支持 "ts","docs","route","protocol", "vite" 分别为 go, typescript, swagger.json 文档, route 为生成的 go route 信息
	tstypes: [], //  langs 中支持 "ts" 时有效, 为生成的 ts 接口文件路径, 默认为 typings/api.ts 文件, 可定义多个, 如: tsapi=["full_api_path1", "full_api_path2"]
	tsapi: [], // langs 中支持 "ts" 时有效, 为生成的 ts 类型文件路径, 默认为 typings/types.d.ts 文件, 可定义多个, 如: tstypes=["full_types_path1","full_types_path2"]
	tsoptional: false, // langs 支持 "ts" 时 optional 是否添加?, 默认为 false
	showlog: false, // 为生成代码时是否生成日志, 默认为不打印日志, 打开时，日志在项目目录下 niuhe.log 中, 生成错误时可进行排查
	endcmd: [], // 为生成代码后执行的命令, 默认为空, 一般第一个为命令名, 后续为参数, 如: go mod tidy 则定义为 ["go","mod","tidy"]
	fire: true, // 生成代码后会显示烟花效果, 默认为 true
}

```