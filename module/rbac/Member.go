package rbac

import "time"

type Member struct {
	Id uint    `gorm:"column:id;auto_increment" ;json:"id" `
	MemberId string  `gorm:"column:member_id;	type:varchar(50);	unique;	not null;" `
	Account  string `gorm:"column:account;form:account;	type:varchar(50);	unique;	not null; binding:required"  json:"account" `
	RealName string  `gorm:"column:real_name;form:realname;	type:varchar(100);	not null " ;json:"real_name"`
	Password string `gorm:"column:password;form:password;	type:varchar(20);	not null ; binding:required" ;json:"password"`
	//认证方式: local 本地数据库 /ldap LDAP
	AuthMethod  int `gorm:"column:auth_method;form:authmethod;	type:int(10);	default:1" ;json:"auth_method"`
	Description string `gorm:"column:description;form:description;	type:varchar(200); " ;json:"description"`
	Email       string `gorm:"column:email;form:email;	type:varchar(50); " ;json:"email"`
	Phone       string `gorm:"column:phone;form:phone;	type:varchar(15); " ;json:"phone"`
	Avatar      string `gorm:"column:avatar;form:avatar;	type:varchar(10); " ;json:"avatar"`
	Status        int       `gorm:"column:status;	type:int(2);	not null	default:1" ;json:"status"` //用户状态：0 正常/1 禁用
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	LastLoginTime time.Time `gorm:"column:last_login_time;	type:datetime;" ;json:"last_login_time"`
}

