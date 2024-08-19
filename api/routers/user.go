package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

var userHandle = handle.NewUserHandle()

// 注册路由
func UserRouterRegister(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	{
		userGroup.POST("", userHandle.Add)
		userGroup.DELETE("", userHandle.Delete)
		userGroup.DELETE("/delete-batch", userHandle.DeleteByIds)
		userGroup.GET("/:id", userHandle.Find)
		userGroup.GET("/list", userHandle.List)
		userGroup.PUT("", userHandle.Update)
	}
}
