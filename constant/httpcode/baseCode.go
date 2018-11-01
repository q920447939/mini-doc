package httpcode


const (
	//系统错误
	BASE_SYS_ERROR_CODE = 41000
	PARAMS_IS_ERROR = 41001
)

var BaseHttpCodesMap = map[int]string{
	BASE_SYS_ERROR_CODE: "系统错误",
	PARAMS_IS_ERROR: "参数错误",
}
