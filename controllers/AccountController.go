package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/utils"
	"wahaha/module/rbac"
	"regexp"
	"wahaha/conf"
	"strings"
	"wahaha/utils/httpUtils"
	"wahaha/base"
	"wahaha/service/impl"
)

//注册页面
func RegisteredHtml(context *gin.Context) {
	baseName, _ := utils.AppConfigMap[utils.BASE_NAME]
	htmlTitle, _ := utils.AppConfigMap[utils.HTML_TITLE]
	baseUrl, _ := utils.AppConfigMap[utils.BASE_URL]
	context.HTML(http.StatusOK, "register.html", gin.H{
		"baseName":            baseName,
		"registeredHTMLTitle": htmlTitle,
		"baseUrl":             baseUrl,
	})

}

//注册
func Register(context *gin.Context) {
	var m rbac.Member
	//由于gin BindJson使用有问题,所以自己写了一个json转结构体的工具类
	//转为map
	httpMap := httpUtils.GetHtppJsonToMap(context.Request)
	//转为struct
	httpUtils.MapToStruct(httpMap, &m)
	errMsg, ok := checkMember(httpMap, &m)
	if !ok {
		b := base.ReturnCode(http.StatusOK, errMsg, nil)
		context.JSON(http.StatusOK, b)
	} else {
		member := impl.Member{}
		e := member.AddMember(&m)
		context.JSON(http.StatusOK, e)
	}
}

func checkMember(httpMap map[string]interface{}, m *rbac.Member) (errMsg string, flg bool) {
	confirmPassword := httpMap["confirmPassword"].(string)

	if ok, err := regexp.MatchString(conf.RegexpAccount, m.Account); !ok || err != nil {
		errMsg = "账号只能由英文字母数字组成，且在3-50个字符"
		return
	}
	if l := strings.Count(m.RealName, ""); l > 50 {
		errMsg = "读者姓名不能为空"
		return
	}

	if l := strings.Count(m.Password, ""); l > 50 || l < 6 {
		errMsg = "密码必须在6-50个字符之间"
		return
	}
	if confirmPassword != m.Password {
		errMsg = "两次密码不一致"
		return
	}
	if ok, err := regexp.MatchString(conf.RegexpEmail, m.Email); !ok || err != nil || m.Email == "" {
		errMsg = "邮箱格式不正确"
		return
	}
	flg = true
	return
}

func Login(context *gin.Context) {
	var m rbac.Member
	httpMap := httpUtils.GetHtppJsonToMap(context.Request)
	httpUtils.MapToStruct(httpMap, &m)
	if errMsg, flg := CheckLoginParams(httpMap); !flg {
		b := base.ReturnCode(http.StatusOK, errMsg, nil)
		context.JSON(http.StatusOK, b)
	}
	member := impl.Member{}
	e := member.Login(&m)
	context.JSON(http.StatusOK, e)
}

func CheckLoginParams(httpMap map[string]interface{}) (errMsg string, flg bool) {
	if httpMap["account"] == "" {
		errMsg = "账号不能为空!"
		return
	}
	if httpMap["password"] == "" {
		errMsg = "密码不能为空!"
		return
	}
	flg = true
	return
}
