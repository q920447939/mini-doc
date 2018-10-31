package httpcode

const (
	Member_Email_Is_Not_Null = 10001
	MEMBER_READ_NAME_IS_EXISTS=10002
	REGISTER_IS_ERROR=10003
	ACCOUNT_OR_PASSWORD_IS_ERR=10004
	CODE_IS_EMPTY=10005
	CODE_IS_NOT_EQUAL=10006
)

var MemberHttpCodes = map[int]string{
	Member_Email_Is_Not_Null: "邮箱不能为空",
	MEMBER_READ_NAME_IS_EXISTS: "该用户已注册",
	REGISTER_IS_ERROR: "注册失败",
	ACCOUNT_OR_PASSWORD_IS_ERR:"账号或密码错误",
	CODE_IS_EMPTY:"验证码不能为空",
	CODE_IS_NOT_EQUAL:"验证码不正确",
}
