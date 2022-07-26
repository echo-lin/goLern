package routers

import (
	"github.com/gin-gonic/gin"
	"golern/controller"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.RegisterUser)
	return r
}
