# 本 demo 说明

本文以 [admin-core-test](https://github.com/ma-guo/admin-core-test) 为例，引用的项目有 [admin-core](https://github.com/ma-guo/admin-core), [vue3-element-admin](https://github.com/ma-guo/vue3-element-admin)。大部分内容与 [admin-core-test/README.md](https://github.com/ma-guo/admin-core-test/blob/main/README.MD) 相同



`admin-core-test` 演示使用 vscode 插件 `niuhe` 接入项目 [admin-core](https://github.com/ma-guo/admin-core/) 例子。前端搭配项目为 [vue3-element-admin](https://github.com/ma-guo/vue3-element-admin)
## 1. 定义 niuhe 文件
 在 `niuhe/all.niuhe` 文件下定义一个 api, 然后点击 `</>` 按钮 生成 `go` 项目代码
```python
#app=admindemo

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
需在 `niuhe/.config` 文件的 `#langs` 中添加 `route`, 配置后会在 `src/{app}/app/api/protos/gen_routes.go` 中生成项目定义的路由信息
```python
#langs=go,ts,route
```
### 5.2 将路由信息接入加入到 protocol 中
接入代码参考 [views/init.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/app/api/views/init.go), 以下代码为 init.go 中的 18-26行
```go
routes := []*coreProtos.RouteItem{}
for _, route := range protos.RouteItems {
        routes = append(routes, &coreProtos.RouteItem{
                Method: route.Method,
                Path:   route.Path,
                Name:   route.Name,
        })
}
coreViews.GetProtocol().AddRoute("", routes)
```

## 6 生成 swagger.json API 文档 - 可选
需要生成 api 文档需要在 langs 中添加 `docs`, 添加后将在会生成文档配置文件 `niuhe/docs.json5`(存在则不生成) 和 `docs/swagger.json`(每次都会复写)。

生成的 swagger.json 文档遵循 swagger 2.0 协议，可直接导入 `postman` 和  `ApiFox(推荐)` 中使用。也可以做成 api 返回 json 作为 URL 导入到 `ApiFox` 中保持实时更新, 参考 [views/system_views.go](https://github.com/ma-guo/admin-core-test/blob/main/src/admincoretest/app/api/views/system_views.go) `Docs_GET` 方法




# 参考 .config 完整配置
> .config 为本地定义文件, 一般不需要跟随 git 版本提交(多成员开发时每个人的环境可能不同)
> 
> 配置项不能添加注释, 下列说明中配置项后面的 // 注释为实例

|  配置项 | 配置说明  | 示例 |
|  ----  | ----  | --- |
| `#langs`  | 支持的语言, 目前支持 `go`,`ts`, `docs`, `route`, 默认支持 `go` | `#langs=ts`
| `#tstypes`  | 自定义 `types.d.ts` 存放路径列表, 以半角逗`,`号分隔 | `#tstypes=~/twerp/typings/lib.props.d.ts` |
| `#tsapi` | 自定义 `api.ts` 存放路径路径列表, 以半角逗号`,`分隔 | `#tsapi=~/twerp/src/utils/api.ts` |
| `#tsoptional` | `optional` 修饰的字段添加 `?`, 默认不添加 | `#tsoptional` |
|`#showlog`|生成代码时是否在 `niuhe.log` 文件中生成过程日志, 在发生错误时辅助调试使用|`#showlog`|
