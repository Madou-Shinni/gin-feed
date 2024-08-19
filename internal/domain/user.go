package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

// User 用户表
type User struct {
	model.Model
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(255);comment:昵称"` // 昵称
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像"`     // 头像
}

type PageUserSearch struct {
	User
	request.PageSearch
}

func (User) TableName() string {
	return "user"
}
