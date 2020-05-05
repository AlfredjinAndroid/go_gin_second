package middleware

import (
	"go_gin_second/common"
	"go_gin_second/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
AuthorMiddelware 验证中心
*/
func AuthorMiddelware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"code": "401",
					"msg":  "权限不足",
				},
			)
			//权限不足，将此次请求抛弃，并返回
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"code": "401",
					"msg":  "权限不足",
				},
			)
			//权限不足，将此次请求抛弃，并返回
			ctx.Abort()
		}

		//验证通过，获取claim中的userId
		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		if user.ID == 0 { //用户不存在
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"code": "401",
					"msg":  "权限不足",
				},
			)
			//权限不足，将此次请求抛弃，并返回
			ctx.Abort()
		}

		//用户存在，将user信息写入上下文
		ctx.Set("User", user)
		ctx.Next()
	}
}
