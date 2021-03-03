module github.com/MSunFlower1014/golang-API

go 1.14

require (
	//indirect 表示传递依赖，非直接依赖
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/sys v0.0.0-20210303074136-134d130e1a04 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。
replace (
	github.com/MSunFlower1014/golang-API/pkg => ./pkg
	github.com/MSunFlower1014/golang-API/config => ./config
	github.com/MSunFlower1014/golang-API/middleware   => ./middleware
	github.com/MSunFlower1014/golang-API/models => ./models
	github.com/MSunFlower1014/golang-API/routers => ./routers
)
