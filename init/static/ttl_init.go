package static

import "github.com/astaxie/beego"
import (
	"wahaha/constant/static"
	)

var TtlConstantMap  = map[string]string{}



func init()  {
	base := beego.AppConfig.String("FONT::BASE")
	name := beego.AppConfig.String("FONT::NAME")
	style := beego.AppConfig.String("FONT::STYLE")
	if base == "" || name == "" || style == "" {
		panic("ttl filepath is empty")
	}
	TtlConstantMap[static.TTL_PATH] = base + name + style
}
