## steps
```
cd your project/config
cp config.yaml.local config.yaml
```

=======
## govendor
```
govendor init
govendor add +external
```
## 添加包
```
govendor add +outside
```
## 同步包
```
govendor sync
```

##controllers使用规范

1.使用目录结构作为根路径 如：controllers/api/v2/onemore/PintCtl 它的根是/api/v2/onemore

2.注册Ctl
#func init() {
#	server.Srv.Register("you", new(YourCtl))
#}
使用controller的注册的名字做为自己的路径名,这里PingCtl的名字是ping，那么它下面方法所有根路径：/api/v2/onemore/you

3.加载ctl
在/controllers/init.go中import你ctl所在的包
#_ "onemore-service-iris/controllers/api/v2/onemore"

4.Ctl的方法遵循iris/mvc的规范

5./controllers/base/ctl.go中基Ctl,用来获取iris.Context
示例:
#type YourCtl struct {
#	base.Ctl
#}