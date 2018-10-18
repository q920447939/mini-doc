package httpcode


const (
	//邮箱不能为空
	BASE_SYS_ERROR_CODE = 41000
)

var BaseHttpCodesMap = map[int]string{
	BASE_SYS_ERROR_CODE: "系统错误",
}
