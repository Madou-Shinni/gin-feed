package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"gorm.io/datatypes"
)

// Moving 动态表
type Moving struct {
	model.Model
	Uid      uint                        `json:"uid" form:"uid"`                // 用户id
	Comment  string                      `json:"comment"`                       // 内容
	Pictures datatypes.JSONSlice[string] `json:"pictures" swaggerignore:"true"` // 图片墙

	// 冗余数据
	Nickname string `json:"nickname" gorm:"-"`
	Avatar   string `json:"avatar" gorm:"-"`
}

type PageMovingSearch struct {
	Moving
	request.PageSearch
}

func (Moving) TableName() string {
	return "moving"
}

// MovingPull 动态发件箱(大up发布动态)
type MovingPull struct {
}

// MovingPush 动态收件箱(小up发布动态到粉丝收件箱)
type MovingPush struct {
}
