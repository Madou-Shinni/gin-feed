package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

var movingHandle = handle.NewMovingHandle()

// 注册路由
func MovingRouterRegister(r *gin.RouterGroup) {
	movingGroup := r.Group("moving")
	{
		movingGroup.POST("", movingHandle.Add)
		movingGroup.DELETE("", movingHandle.Delete)
		movingGroup.DELETE("/delete-batch", movingHandle.DeleteByIds)
		movingGroup.GET("/:id", movingHandle.Find)
		movingGroup.GET("/list", movingHandle.List)
		movingGroup.PUT("", movingHandle.Update)
	}
}
