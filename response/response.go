package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义返回数据格式
func Response(ctx *gin.Context, httpState int, code int, data gin.H, msg string) {
	ctx.JSON(httpState, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//Success 定义正确的基本返回格式
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//Fail 定义失败的基本返回格式
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
