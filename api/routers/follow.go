package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

var followHandle = handle.NewFollowHandle()

// 注册路由
func FollowRouterRegister(r *gin.RouterGroup) {
	followGroup := r.Group("follow")
	{
		followGroup.POST("", followHandle.Add)
		followGroup.DELETE("", followHandle.Delete)
		followGroup.DELETE("/delete-batch", followHandle.DeleteByIds)
		followGroup.GET("/:id", followHandle.Find)
		followGroup.GET("/list", followHandle.List)
		followGroup.PUT("", followHandle.Update)
	}
}
