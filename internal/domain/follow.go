package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

// Follow 关注表
type Follow struct {
	model.Model
	Follower uint `json:"follower" gorm:"column:follower;type:int(11);comment:关注者"`  // 关注者
	Followed uint `json:"followed" gorm:"column:followed;type:int(11);comment:被关注者"` // 被关注者
}

type PageFollowSearch struct {
	Follow
	request.PageSearch
}

func (Follow) TableName() string {
	return "follow"
}
