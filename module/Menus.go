package module

import "time"

type Menus struct {
	Id uint    `gorm:"column:id;auto_increment" ;json:"id" `
	Name string  `gorm:"column:name;	type:varchar(50);	unique;	not null;" ;json:"name" `
	Types        int       `gorm:"column:types;	type:int(5);	not null	default:1" ;json:"types"`
	Status        int       `gorm:"column:status;	type:int(2);	not null	default:1" ;json:"status"` //用户状态：0 正常/1 禁用
	CreatedAt time.Time `;json:"-"`
	UpdatedAt time.Time `;json:"-"`
	LastLoginTime time.Time `gorm:"column:last_login_time;	type:datetime;" ;json:"-"`
}
