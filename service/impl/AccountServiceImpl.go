package impl

import (
	"wahaha/module/rbac"
	"wahaha/models"
	"wahaha/base"
	"github.com/satori/go.uuid"
	"encoding/json"
	"wahaha/constant"
	"wahaha/connections/redis"
	"wahaha/constant/httpcode"
	"net/http"
	"sync"
)

type Member struct {
}

func (member *Member) AddMember(m *rbac.Member) (e *base.BaseReturnJson) {
	var r rbac.Member
	e = new(base.BaseReturnJson)
	if  models.Model.Where("account = ?", m.Account).First(&r); r.MemberId != "" {
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

func (member *Member) Login(m *rbac.Member) (e *base.BaseReturnJson) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	e = new(base.BaseReturnJson)
	if err := models.Model.Where("account = ? and password = ? and status =1", m.Account,m.Password).Find(m).Error; err == nil {
		e.Code = httpcode.ACCOUNT_OR_PASSWORD_IS_ERR
		e.Message = httpcode.MemberHttpCodes[httpcode.ACCOUNT_OR_PASSWORD_IS_ERR]
		return
	}
	memberChannel := make(chan string, 1)
	go func() {
		memberJson, _ := redis.Client.Get(constant.MEMBERS_JSON + m.MemberId)
		if memberJson == "" {
			addMemberToRedis(m)
		}
		memberChannel <- memberJson
	}()
	if <-memberChannel != "" {
		e.Code = http.StatusOK
		e.ExecuteStatus = true
	}
	return

}

func AddUser(m *rbac.Member) bool {
	id, e := uuid.NewV4()
	if e != nil {
		panic(e)
	}
	m.MemberId = id.String()

	err := models.Model.Create(&m).Error
	if err != nil {
		panic(err)
	} else {
		//把用户数据保存到redis
		go func() {
			addMemberToRedis(m)
		}()
		//todo 发送邮箱验证用户

		return true
	}

	return false
}

func addMemberToRedis(m *rbac.Member) {
	m.Password = ""
	json, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	redis.Client.Set(constant.MEMBERS_JSON+m.MemberId, string(json), constant.EXPIRE_TIME)
}
