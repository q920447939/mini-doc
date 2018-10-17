package rbac

import "time"

type Member struct {
	MemberId string    `gorm:"column:member_id;	type:string(20);	primary_key;	auto_increment" json:"member_id;" `
	Account  string `gorm:"column:account;	type:varchar(10);	unique;	not null" json:"account"  binding:"required"`
	RealName string  `gorm:"column:real_name;	type:varchar(10);	not null " json:"real_name"`
	Password string `gorm:"column:password;	type:varchar(10);	not null " json:"password"  binding:"required"`
	//认证方式: local 本地数据库 /ldap LDAP
	AuthMethod  string `gorm:"column:auth_method;	type:int(10);	not null;	default:1" json:"auth_method"`
	Description string `gorm:"column:description;	type:varchar(200); " json:"description"`
	Email       string `gorm:"column:email;	type:varchar(50); " json:"email"`
	Phone       string `gorm:"column:phone;	type:varchar(15); " json:"phone"`
	Avatar      string `gorm:"column:avatar;	type:varchar(10); " json:"avatar"`
	Status        int       `gorm:"column:status;	type:int(2);	not null	default:1" json:"status"` //用户状态：0 正常/1 禁用
	CreateTime    time.Time `gorm:"column:create_time;	type:timestamp(100); " json:"create_time"`
	CreateAt      int       `gorm:"column:create_at;	type:varchar(11);;	not null" json:"create_at"`
	LastLoginTime time.Time `gorm:"column:last_login_time;	type:timestamp(100);" json:"last_login_time"`
}

