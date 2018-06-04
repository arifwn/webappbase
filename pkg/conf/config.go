package conf

import (
	"os"
)

type WebConfig struct {
	ServerAddress string
	ServerPort    string
	StaticDir     string
	UploadDir     string
	TemplateDir   string
	AppDir        string
	DBType        string
	DBConf        string
}

func Get() WebConfig {
	conf := WebConfig{
		ServerAddress: GetEnv("SERVER_ADDRESS", "127.0.0.1"),
		ServerPort:    GetEnv("SERVER_PORT", "8080"),
		StaticDir:     GetEnv("STATIC_DIR", "./web/static"),
		UploadDir:     GetEnv("UPLOAD_DIR", "./web/upload"),
		TemplateDir:   GetEnv("TEMPLATE_DIR", "./web/template"),
		AppDir:        GetEnv("APP_DIR", "./web/app"),
		DBType:        GetEnv("DB_TYPE", "sqlite3"),
		DBConf:        GetEnv("DB_CONF", "./web/sqlite.db"),
	}

	return conf
}

func GetEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	} else {
		return val
	}
}
