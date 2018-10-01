package settings

import "os"

type Database struct {
	Host string
	Port int
	User string
	Password string
	Name string
}


func GetDatabaseConfig() Database {
	Env := os.Getenv("AhaEnv")
	switch Env {  // TODO: 增加生产环境和测试环境的配置返回
	case "prod":
		return Database{}
	case "test":
		return Database{}
	default:
		return DevDBConfig
	}
}
