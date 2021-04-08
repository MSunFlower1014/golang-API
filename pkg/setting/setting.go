package setting

import (
	"github.com/astaxie/beego/validation"
	"github.com/go-ini/ini"
	"github.com/prometheus/common/log"
	"os"
)
import (
	"time"
)

const ConfigPath = "config/app.ini"

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string

	BookDataJsonPath string
	BookLifeJsonPath string
)

func init() {
	var err error
	_, err = os.Open(ConfigPath)
	if !os.IsExist(err) {
		log.Errorf("Fail to parse %v , error : %v", ConfigPath, err)
		defaultInit()
		return
	}
	Cfg, err = ini.Load(ConfigPath)

	if err != nil {
		log.Errorf("Fail to parse %v , error : %v", ConfigPath, err)
		return
	}

	LoadBase()
	loadServer()
	loadJwt()
	loadBook()
}

func LoadBase() {
	RunMode = Cfg.Section("gin").Key("RUN_MODE").MustString("debug")
}

const ServerSection = "server"

func loadServer() {
	serverConfig, err := Cfg.GetSection(ServerSection)
	if err != nil {
		log.Errorf("GetSection %v  error : %v", ServerSection, err)
	}
	HTTPPort = serverConfig.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(serverConfig.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(serverConfig.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	valid := validation.Validation{}
	valid.Required(HTTPPort, "Port").Message("port cant empty")
	if valid.HasErrors() {
		log.Errorf("validation error : %v", valid.Errors)
	}
}

const JwtConfig = "jwt"

func loadJwt() {
	jwtConfig, err := Cfg.GetSection(JwtConfig)
	if err != nil {
		log.Errorf("GetSection %v  error : %v", JwtConfig, err)
	}

	JwtSecret = jwtConfig.Key("JWT_SECRET").MustString("test")
}

const BookConfig = "book"

func loadBook() {
	jwtConfig, err := Cfg.GetSection(BookConfig)
	if err != nil {
		log.Errorf("GetSection %v  error : %v", JwtConfig, err)
		BookDataJsonPath = "."
		BookLifeJsonPath = "."
		return
	}

	BookDataJsonPath = jwtConfig.Key("DATA_JSON_PATH").MustString(".")
	BookLifeJsonPath = jwtConfig.Key("LIFT_JSON_PATH").MustString(".")
}

func defaultInit() {
	BookDataJsonPath = "./book_data.json"
	BookLifeJsonPath = "./book_life.json"
}
