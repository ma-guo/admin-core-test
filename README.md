# 本 demo 说明
本项目演示使用 vscode niuhe 插件 admin-core 接入例子。 
### 1. 定义 niuhe 文件
 在 `niuhe/all.niuhe` 文件下定义一个 api, 然后点击 `</>` 按钮 生成 go 项目代码
```python
#app=admincoretest

class NoneReq(Message):
    pass

class NoneRsp(Message):
    pass


with services():
    GET('测试  api', '/api/system/test/', NoneReq, wraps(NoneRsp))
```
### 2. 修改配置文件
在配置 conf/admincoretest.yaml 文件中添加第8和第9行并修改第7行的数据库连接信息
### 3. 接入 AdminBoot
在 `src/admincoretest/main.go` 中接入 `AdminBoot`
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
上述 `main` 方式是自定生成的, 将 `AdminBoot` 引入, 初始化 config(`LoadConfig`) 并 `RegisterModules 即可`
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
下面为插件自动生成的文档说明


# mod
```sh
    go env -w GO111MODULE=auto
    cd src/admincoretest && go mod init admincoretest && go mod tidy && go mod vendor && cd ../../ && make run
```

db 配置格式
```yaml
db:
	main:user:pwd@tcp(host:port)/database_name?charset=utf8mb4
```

# 更多自定义信息
在 `niuhe` 文件夹下新建文件 `.config`, 注: 
- .config 为本地定义文件, 不需要跟随 git 版本提交
- 配置项不能添加注释, 下列说明中配置项后面的 // 注释为实例
## 支持生成 typescript
```sh
#langs=ts
#tstypes=full_types_file_path // 完整文件地址
#tsapi=full_api_file_path // 完整文件地址
#tsoptional // ts 中 optional 转换为 ?
```
完整示例
|  配置项 | 配置说明  | 示例 |
|  ----  | ----  | --- |
| `#langs`  | 支持的语言, 目前支持 `go`,`ts`, 默认支持 `go` | `#langs=ts`
| `#tstypes`  | 自定义 `types.d.ts` 路径 | `#tstypes=~/twerp/typings/lib.props.d.ts` |
| `#tsapi` | 自定义 `api.ts` 路径 | `#tsapi=~/twerp/src/utils/api.ts` |
| `#tsoptional` | `optional` 修饰的字段添加 `?`, 默认不添加 | `#tsoptional` |
