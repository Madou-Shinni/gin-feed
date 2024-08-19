package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

var videoHandle = handle.NewVideoHandle()

// 注册路由
func VideoRouterRegister(r *gin.RouterGroup) {
	videoGroup := r.Group("video")
	{
		videoGroup.POST("", videoHandle.Add)
		videoGroup.DELETE("", videoHandle.Delete)
		videoGroup.DELETE("/delete-batch", videoHandle.DeleteByIds)
		videoGroup.GET("/:id", videoHandle.Find)
		videoGroup.GET("/list", videoHandle.List)
		videoGroup.PUT("", videoHandle.Update)
	}
}
