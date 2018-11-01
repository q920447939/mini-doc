package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/utils"
	"wahaha/module/rbac"
	"regexp"
	"wahaha/conf"
	"strings"
	"wahaha/base"
	"wahaha/service/impl"
	"wahaha/utils/jwt"
	"wahaha/utils/gocaptcha"
	"wahaha/constant/httpcode"
	"github.com/gin-gonic/contrib/sessions"
		"github.com/goinggo/mapstructure"
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
		"captcha":             "http://" + conf.GetEnv().ServerConfig.SERVER_IP + ":" + conf.GetEnv().ServerConfig.SERVER_PORT + "/login/captcha",
	})

}

//注册
func Register(context *gin.Context) {
	var m rbac.Member
	mapInstance := make(map[string]interface{})
	if err := context.BindJSON(&mapInstance); err != nil {
		base.ReturnBaseCode_Fail(httpcode.PARAMS_IS_ERROR, httpcode.MemberHttpCodes[httpcode.PARAMS_IS_ERROR], nil, context)
		return
	}
	//将map转为struct
	err := mapstructure.Decode(mapInstance, &m)
	if err != nil {
		base.ReturnBaseCode_Fail(httpcode.BASE_SYS_ERROR_CODE, httpcode.BaseHttpCodesMap[httpcode.BASE_SYS_ERROR_CODE], nil, context)
		return
	}
	if queryCode, exists := mapInstance["code"]; !exists || queryCode == "" {
		base.ReturnBaseCode_Fail(httpcode.CODE_IS_EMPTY, httpcode.MemberHttpCodes[httpcode.CODE_IS_EMPTY], nil, context)
		return
	} else {
		//从session 获取验证码
		session := sessions.Default(context)
		if sessionCode, ok := session.Get(conf.CaptchaSessionName).(string); !ok {
			base.ReturnBaseCode_Fail(httpcode.CODE_IS_NOT_EQUAL, httpcode.MemberHttpCodes[httpcode.CODE_IS_NOT_EQUAL], nil, context)
			return
		} else {
			if strings.ToLower(queryCode.(string)) != strings.ToLower(sessionCode) {
				base.ReturnBaseCode_Fail(httpcode.CODE_IS_NOT_EQUAL, httpcode.MemberHttpCodes[httpcode.CODE_IS_NOT_EQUAL], nil, context)
				return
			}
			//检查参数
			errMsg, ok := checkMember(&m, mapInstance)
			if !ok {
				base.ReturnBaseCode_Fail(httpcode.CODE_IS_EMPTY, errMsg, nil, context)
				return
			} else {
				member := impl.Member{}
				if baseResult := member.AddMember(&m); baseResult.ExecuteStatus {
					//登陆成功
					base.ReturnBaseCode_Success(http.StatusOK, "注册成功", nil, context)
					return
				} else {
					base.ReturnBaseCode_Fail(httpcode.MEMBER_READ_NAME_IS_EXISTS, httpcode.MemberHttpCodes[httpcode.MEMBER_READ_NAME_IS_EXISTS], nil, context)
					return
				}
				return
			}
		}
	}
}

func checkMember(m *rbac.Member, mapInstance map[string]interface{}) (errMsg string, flg bool) {
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
	if mapInstance["confirmPassword"] != m.Password {
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
	if err := context.BindJSON(&m); err != nil {
		panic(err)
		return
	}
	if errMsg, flg := CheckLoginParams(&m); !flg {
		base.ReturnBaseCode_Fail(http.StatusBadRequest, errMsg, nil, context)
	}
	member := impl.Member{}

	if e := member.Login(&m); !e.ExecuteStatus {
		context.JSON(http.StatusOK, e)
		return
	}
	//放到session
	if token, err := jwt.GenerateToken(m.Account, m.Password); err != nil {
		panic(err)
	} else {
		slice := strings.Split(context.Request.URL.String(), "redirect_url")
		var redictUrl string
		if len(slice) > 1 {
			redictUrl = slice[1]
		} else {
			redictUrl = "minidoc"
		}
		context.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"data":       token,
			"redict_url": redictUrl,
		})
	}
}

func CheckLoginParams(m *rbac.Member) (errMsg string, flg bool) {
	if m.Account == "" {
		errMsg = "账号不能为空!"
		return
	}
	if m.Password == "" {
		errMsg = "密码不能为空!"
		return
	}
	flg = true
	return
}

// 验证码
func Captcha(context *gin.Context) {
	captchaImage, err := gocaptcha.NewCaptchaImage(140, 40, gocaptcha.RandLightColor())
	if err != nil {
		panic(err)
	}
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)
	txt := gocaptcha.RandText(4)
	session := sessions.Default(context)
	session.Set(conf.CaptchaSessionName, txt)
	session.Save()
	captchaImage.DrawText(txt)
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))

	captchaImage.SaveImage(context.Writer, gocaptcha.ImageFormatJpeg)
}
