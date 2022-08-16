package middleware

import (
	"github.com/gin-gonic/gin"
	"golern/common"
	"golern/model"
	"net/http"
	"strings"
)

// AuthMiddleWare
// @Description 验证token
// @Date 2022-07-27 06:33:57
// @return gin.HandlerFunc
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// token是否为空活着Bearer开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"code": 511, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"code": 511, "msg": "权限不足"})
			c.Abort()
			return
		}
		userId := claims.UserId
		var user model.User
		common.DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"code": 511, "msg": "权限不足"})
			c.Abort()
			return
		}
		// 用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
