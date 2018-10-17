package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/utils"
	"wahaha/models"
	"wahaha/module/rbac"
	"wahaha/constant"
	"wahaha/base"
	"regexp"
	"wahaha/conf"
	"strings"
		"github.com/satori/go.uuid"
)

var db = models.Model

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
	if context.BindJSON(m, ) != nil {
		errMsg, ok := checkMember(m, context)
		if !ok {
			b := base.ReturnCode(constant.BaseCode, errMsg, m)
			context.String(b.Code, base.GetJsonStr(b))
		} else {
			AddUser(m)

		}
	}
}

func checkMember(m rbac.Member, context *gin.Context) (errMsg string, flg bool) {
	//code := context.PostForm("code")
	email := context.PostForm("email")
	confirmPassword := context.PostForm("confirmPassword")

	if m.Email == "" {
		errMsg = "邮箱不能为空"
		return
	}
	if ok, err := regexp.MatchString(conf.RegexpAccount, m.Account); !ok || err != nil {
		errMsg = "账号只能由英文字母数字组成，且在3-50个字符"
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
	if ok, err := regexp.MatchString(conf.RegexpEmail, email); !ok || err != nil || email == "" {
		errMsg = "邮箱格式不正确"
		return
	}
	m.Email = email
	flg = true
	return
}

func AddUser(m  rbac.Member) {
	uuid, e := uuid.NewV4()
	if e != nil {
		panic(e)
	}
	m.MemberId =  uuid.String()
	oK := db.NewRecord(&m)
	if oK {
		go func() {}()
	}
}


