package base

import "encoding/json"

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

func ReturnCode(code int, message string, v interface{}) *BaseReturnJson {
	b := BaseReturnJson{
		Code:    code,
		Message: message,
		Data:    v,
	}
	return &b
}
