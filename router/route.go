package router

import (
	"go_gin_second/common/middleware"
	"go_gin_second/controller"

	"github.com/gin-gonic/gin"
)

// CollectRoute 控制中心
func CollectRoute(g *gin.Engine) *gin.Engine {
	g.POST("/api/auth/register", controller.Register)
	g.POST("/api/auth/login", controller.Login)
	g.GET("/api/auth/info", middleware.AuthorMiddelware(), controller.Info) //引入中间件保护用户信息接口，若通过验证，则会获取到用户信息
	return g
}
