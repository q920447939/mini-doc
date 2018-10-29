package module

import "time"

// Book dto .
type Book struct {
	BookId int `gorm:"pk;auto;unique;column:book_id" json:"book_id"`
	// BookName 项目名称.
	BookName string `gorm:"column:book_name;size(500)" json:"book_name"`
	// Identify 项目唯一标识.
	Identify string `gorm:"column:identify;size(100);unique" json:"identify"`
	//是否是自动发布 0 否/1 是
	AutoRelease int `gorm:"column:auto_release;type(int);default(0)" json:"auto_release"`
	//是否开启下载功能 0 是/1 否
	IsDownload int `gorm:"column:is_download;type(int);default(0)" json:"is_download"`
	OrderIndex int `gorm:"column:order_index;type(int);default(0)" json:"order_index"`
	// Description 项目描述.
	Description string `gorm:"column:description;size(2000)" json:"description"`
	//发行公司
	Publisher string `gorm:"column:publisher;size(500)" json:"publisher"`
	Label     string `gorm:"column:label;size(500)" json:"label"`
	// PrivatelyOwned 项目私有： 0 公开/ 1 私有
	PrivatelyOwned int `gorm:"column:privately_owned;type(int);default(0)" json:"privately_owned"`
	// 当项目是私有时的访问Token.
	PrivateToken string `gorm:"column:private_token;size(500);null" json:"private_token"`
	//状态：0 正常/1 已删除
	Status int `gorm:"column:status;type(int);default(0)" json:"status"`
	//默认的编辑器.
	Editor string `gorm:"column:editor;size(50)" json:"editor"`
	// DocCount 包含文档数量.
	DocCount int `gorm:"column:doc_count;type(int)" json:"doc_count"`
	// CommentStatus 评论设置的状态:open 为允许所有人评论，closed 为不允许评论, group_only 仅允许参与者评论 ,registered_only 仅允许注册者评论.
	CommentStatus string `gorm:"column:comment_status;size(20);default(open)" json:"comment_status"`
	CommentCount  int    `gorm:"column:comment_count;type(int)" json:"comment_count"`
	//封面地址
	Cover string `gorm:"column:cover;size(1000)" json:"cover"`
	//主题风格
	Theme string `gorm:"column:theme;size(255);default(default)" json:"theme"`
	// CreateTime 创建时间 .
	CreateTime time.Time `gorm:"type(datetime;column:create_time);auto_now_add" json:"create_time"`
	//每个文档保存的历史记录数量，0 为不限制
	HistoryCount int `gorm:"column:history_count;type(int);default(0)" json:"history_count"`
	//是否启用分享，0启用/1不启用
	IsEnableShare int       `gorm:"column:is_enable_share;type(int);default(0)" json:"is_enable_share"`
	MemberId      int       `gorm:"column:member_id;size(100)" json:"member_id"`
	ModifyTime    time.Time `gorm:"type(datetime);column:modify_time;null;auto_now" json:"modify_time"`
	Version       int64     `gorm:"type(bigint);column:version" json:"version"`
	//是否使用第一篇文章项目为默认首页,0 否/1 是
	IsUseFirstDocument int `gorm:"column:is_use_first_document;type(int);default(0)" json:"is_use_first_document"`
	//是否开启自动保存：0 否/1 是
	AutoSave	int 		`gorm:"column:auto_save;type(tinyint);default(0)" json:"auto_save"`
}
