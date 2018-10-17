package constant

const (
	//member 10000 - 10100

	//邮箱不能为空
	BaseCode = 10000
	Member_Email_Is_Not_Null = 10001
)

var BaseHttpCodes = map[int]string{
	BaseCode: "",
	Member_Email_Is_Not_Null: "邮箱不能为空",
}
