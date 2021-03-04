package setting

import (
	"github.com/astaxie/beego/validation"
	"github.com/go-ini/ini"
)
import (
	"log"
	"time"
)

const ConfigPath = "D:\\Study\\Golang\\golang-API\\config\\app.ini"

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load(ConfigPath)

	if err != nil {
		log.Fatalf("Fail to parse %v , error : %v", ConfigPath, err)
	}

	LoadBase()
	loadServer()
	loadJwt()
}

func LoadBase() {
	RunMode = Cfg.Section("gin").Key("RUN_MODE").MustString("debug")
}

const ServerSection = "server"

func loadServer() {
	serverConfig, err := Cfg.GetSection(ServerSection)
	if err != nil {
		log.Fatalf("GetSection %v  error : %v", ServerSection, err)
	}
	HTTPPort = serverConfig.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(serverConfig.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(serverConfig.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	valid := validation.Validation{}
	valid.Required(HTTPPort, "Port").Message("port cant empty")
	if valid.HasErrors() {
		log.Fatalf("validation error : %v", valid.Errors)
	}
}

const JwtConfig = "jwt"

func loadJwt() {
	jwtConfig, err := Cfg.GetSection(JwtConfig)
	if err != nil {
		log.Fatalf("GetSection %v  error : %v", JwtConfig, err)
	}

	JwtSecret = jwtConfig.Key("JWT_SECRET").MustString("test")
}
