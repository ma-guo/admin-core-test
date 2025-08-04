# 教程请查阅 [DOCS.md](./DOCS.md)
交流群 QQ [971024252] [点击链接加入群聊【niuhe插件支持】：https://qm.qq.com/q/vTaMfRNwL]
![971024252](./assets/qrcode.jpg)


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
在 `niuhe` 文件夹下新建文件 `.config.json5`, 注:
- .config.json5 为本地定义文件, 不需要跟随 git 版本提交
- 配置项不能添加注释, 下列说明中配置项后面的 // 注释为实例

## 支持生成 typescript
项目配置项位于 `niuhe/.config.json5` 文件中, 关于生成 `ts` 代码, 在 `langs` 中增加 `"ts"` 即可

```sh
{
    "langs": ["ts"], // 在 langs 中添加 "ts" 支持
}
```
完整示例
```json5
{
    "langs": ["ts"], // 在 langs 中添加 "ts" 支持
    "tstypes": ["~/admincoretest/typings/lib.props.d.ts"], // 自定义 types.d.ts 路径
    "tsapi": ["~/admincoretest/src/utils/api.ts"], // 自定义 api.ts 路径
}
```
