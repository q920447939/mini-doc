package conf

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	ServerConfig   ServerConf
	RedisConfig    RedisConf
	DatebaseConfig DateBaseConf
}

func init() {
	serverConf := ServerConf{
		DEBUG:           true,
		SERVER_IP:       "127.0.0.1",
		SERVER_PORT:     "4000",
		ACCESS_LOG:      false,
		ACCESS_LOG_PATH: "storage/logs/access.log",

		ERROR_LOG:      false,
		ERROR_LOG_PATH: "storage/logs/error.log",

		INFO_LOG:      false,
		INFO_LOG_PATH: "storage/logs/info.log",

		TEMPLATE_PATH: "views/**",

		//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
		APP_SECRET: "something-very-secret",
	}
	redisConf := RedisConf{
		REDIS_IP:       "127.0.0.1",
		REDIS_PORT:     "6379",
		REDIS_PASSWORD: "",
		REDIS_DB:       0,

		REDIS_SESSION_DB: 1,
		REDIS_CACHE_DB:   2,
	}
	datebaseConf := DateBaseConf{
		DATABASE_IP:       "127.0.0.1",
		DATABASE_PORT:     "3306",
		DATABASE_USERNAME: "root",
		DATABASE_PASSWORD: "123456",
		DATABASE_NAME:     "gin-template",
	}
	EnvConf = Env{
		ServerConfig:   serverConf,
		RedisConfig:    redisConf,
		DatebaseConfig: datebaseConf,
	}
}

var EnvConf Env

func GetEnv() *Env {
	return &EnvConf
}

type ServerConf struct {
	SERVER_IP   string
	SERVER_PORT string
	DEBUG       bool
	APP_SECRET  string

	ACCESS_LOG      bool
	ACCESS_LOG_PATH string
	ERROR_LOG       bool
	ERROR_LOG_PATH  string
	INFO_LOG        bool
	INFO_LOG_PATH   string
	TEMPLATE_PATH   string // 静态文件相对路径
}

type RedisConf struct {
	REDIS_IP         string
	REDIS_PORT       string
	REDIS_PASSWORD   string
	REDIS_DB         int
	REDIS_SESSION_DB int
	REDIS_CACHE_DB   int
}

type DateBaseConf struct {
	DATABASE_IP       string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	SQL_LOG           bool
}
