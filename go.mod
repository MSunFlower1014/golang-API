module github.com/MSunFlower1014/golang-API

go 1.14

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	//indirect 表示传递依赖，非直接依赖
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/prometheus/common v0.10.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/sys v0.0.0-20210303074136-134d130e1a04 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。
replace (
	github.com/MSunFlower1014/golang-API/config => ./config
	github.com/MSunFlower1014/golang-API/middleware => ./middleware
	github.com/MSunFlower1014/golang-API/models => ./models
	github.com/MSunFlower1014/golang-API/pkg => ./pkg
	github.com/MSunFlower1014/golang-API/routers => ./routers
)
