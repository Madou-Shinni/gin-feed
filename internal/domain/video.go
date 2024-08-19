package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

// Video 视频表
type Video struct {
	model.Model
	Title     string `json:"title"`     // 标题
	Url       string `json:"url"`       // 链接
	Cover     string `json:"cover"`     // 封面
	Likes     int    `json:"likes"`     // 喜欢数
	Views     int    `json:"views"`     // 观看数
	Favorites int    `json:"favorites"` // 收藏数
	Comments  int    `json:"comments"`  // 评论数
}

type PageVideoSearch struct {
	Video
	request.PageSearch
}

func (Video) TableName() string {
	return "video"
}
