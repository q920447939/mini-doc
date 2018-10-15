package utils

import (
	"github.com/astaxie/beego/config"
	"errors"
)

/**
FormatType 文件格式
FilePath 文件类型
ConfigName 读取文件名字
 */
type Config struct {
	FormatType string
	FilePath   string
	ConfigName string
}

var AppConfig Config
var AppConfigMap map[string]string

const (
	HTTPADDR   = "httpaddr"
	BASE_NAME  = "BASE::NAME"
	BASE_URL   = "BASE::URL"
	HTML_TITLE = "BASE::HTML_TITLE"
)

func init() {
	AppConfigMap = make(map[string]string)
	loadConf()
}

func loadConf() {
	AppConfig = Config{
		FormatType: "ini",
		FilePath:   "conf/app.conf",
	}
	loadConfString(HTTPADDR, BASE_NAME, BASE_URL, HTML_TITLE)

}

func loadConfString(confName ...string) {
	for _, v := range confName {
		AppConfig.ConfigName = v
		s, err := StringConf(&AppConfig)
		if err != nil {
			panic(err)
		}

		AppConfigMap[v] = s
	}
}

func StringConf(c *Config) (s string, err error) {
	if c.FormatType == "" {
		err = errors.New(c.FormatType + " is null")
	}
	if c.FilePath == "" {
		err = errors.New(c.FilePath + " is null")
	}
	switch c.FormatType {
	case "ini":
		config, e := config.NewConfig("ini", c.FilePath)
		if e != nil {
			err = e
			return
		}
		s = config.String(c.ConfigName)
		break
	}
	return
}
