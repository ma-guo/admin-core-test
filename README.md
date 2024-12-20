# 教程请查阅 [DOCS.md](./DOCS.md)
交流群 QQ [971024252] [点击链接加入群聊【niuhe插件支持】：https://qm.qq.com/q/vTaMfRNwL]


以下文档为插件首次运行时自动生成

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
| `#tstypes`  | 自定义 `types.d.ts` 路径 | `#tstypes=~/admincoretest/typings/lib.props.d.ts` |
| `#tsapi` | 自定义 `api.ts` 路径 | `#tsapi=~/admincoretest/src/utils/api.ts` |
| `#tsoptional` | `optional` 修饰的字段添加 `?`, 默认不添加 | `#tsoptional` |
