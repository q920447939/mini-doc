package base

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseReturnJson struct {
	ExecuteStatus bool `json:"-"` //执行状态
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetJsonStr(t *BaseReturnJson) string {
	bytes, e := json.MarshalIndent(t, " ", " ")
	if e != nil {
		panic(e)
	}
	return string(bytes)
}

func ReturnBaseCode_Success(code int, message string, v interface{} , context *gin.Context) {
	b := BaseReturnJson{
		Code:    code,
		Message: message,
		Data:    v,
	}
	context.JSON(http.StatusOK, b)
	return
}


func ReturnBaseCode_Fail(code int, message string, v interface{} , context *gin.Context) {
	b := BaseReturnJson{
		Code:    code,
		Message: message,
		Data:    v,
	}
	context.JSON(http.StatusOK, b)
	return
}

