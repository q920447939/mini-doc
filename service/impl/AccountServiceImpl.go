package impl

import (
	"wahaha/module/rbac"
		"wahaha/models"
		"wahaha/base"
	"github.com/satori/go.uuid"
		"net/http"
	"wahaha/constant/httpcode"
	"encoding/json"
	"wahaha/constant"
	"wahaha/connections/redis"
	)

type Member struct {
}

func (member *Member) AddMember(m rbac.Member) (e *base.BaseReturnJson) {
	var r  rbac.Member
	e = new(base.BaseReturnJson)
	if error := models.Model.Where("account = ?", m.Account).First(&r).Error ; error != nil{
		e.Code = httpcode.MEMBER_READ_NAME_IS_EXISTS
		e.Message = httpcode.MemberHttpCodes[httpcode.MEMBER_READ_NAME_IS_EXISTS]
		return
	}

	if !AddUser(m) {
		e.Code = httpcode.REGISTER_IS_ERROR
		e.Message = httpcode.BaseHttpCodesMap[httpcode.REGISTER_IS_ERROR]
		return e
	} else {
		e.Code = http.StatusOK
		e.ExecuteStatus = true
		return e
	}
}

func AddUser(m rbac.Member) bool {
	id, e := uuid.NewV4()
	if e != nil {
		panic(e)
	}
	m.MemberId = id.String()

	err := models.Model.Create(&m).Error
	if err != nil {
		panic(err)
	}else{
		//把用户数据保存到redis
		go func() {
			m.Password = ""
			json, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			redis.Client.Set(constant.MEMBERS_JSON+m.MemberId, string(json), constant.EXPIRE_TIME)
		}()
		//todo 发送邮箱验证用户

		return true
	}

	return false
}
