module github.com/MSunFlower1014/golang-API

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/pprof v1.3.0
	//indirect 表示传递依赖，非直接依赖
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/hjr265/redsync.go v0.0.0-20160719150818-688f6d364b79
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/prometheus/common v0.10.0
	github.com/robfig/cron v1.2.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stvp/tempredis v0.0.0-20181119212430-b82af8480203 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b // indirect
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4 // indirect
	golang.org/x/sys v0.0.0-20210317091845-390168757d9c // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)

//用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。
replace (
	github.com/MSunFlower1014/golang-API/config => ./config
	github.com/MSunFlower1014/golang-API/middleware => ./middleware
	github.com/MSunFlower1014/golang-API/models => ./models
	github.com/MSunFlower1014/golang-API/pkg => ./pkg
	github.com/MSunFlower1014/golang-API/routers => ./routers
)
