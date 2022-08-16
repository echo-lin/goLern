package routers

import (
	"github.com/gin-gonic/gin"
	"golern/controller"
	"golern/middleware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.RegisterUser)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleWare(), controller.UserInfo)
	r.POST("/api/document/addDoc", middleware.AuthMiddleWare(), controller.AddDoc)
	r.POST("/api/document/docList", middleware.AuthMiddleWare(), controller.DocList)
	return r
}
