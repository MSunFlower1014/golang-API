# golang-API
go语言学习

1. go mod
```shell script
# 打开 Go modules 开关（目前在 Go1.13 中默认值为 auto）。   
go env -w GO111MODULE=on  
# 设置 GOPROXY 代理，这里主要涉及到两个值，第一个是 https://goproxy.cn，它是由七牛云背书的一个强大稳定的 Go 模块代理，可以有效地解决你的外网问题；
# 第二个是 direct，它是一个特殊的 fallback 选项，它的作用是用于指示 Go 在拉取模块时遇到错误会回源到模块版本的源地址去抓取（比如 GitHub 等）。
go env -w GOPROXY=https://goproxy.cn,direct  
# 初始化 Go modules，它将会生成 go.mod 文件
go mod init github.com/MSunFlower1014/golang-API  
```

2. gin 依赖  
```shell script
# 开源web框架
go get -u github.com/gin-gonic/gin

#go mod tidy 可以用来整理依赖
go mod tidy
```

3. go ini 依赖  
```shell script
go get -u github.com/go-ini/ini

##在init方法中获取初始化配置信息

```

4. validation 依赖
```shell script
# 表单验证
go get -u github.com/astaxie/beego/validation
```

5. gRpc 依赖
```shell script
go get -u google.golang.org/grpc
```

6. swagger go
```shell script
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template

#初始化swagger
swag init

#路由器初始化需要导入生成doc地址
_ "github.com/MSunFlower1014/golang-API/docs"

#访问地址 
# http://127.0.0.1:8080/swagger/index.html
```