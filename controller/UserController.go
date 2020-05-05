package controller

import (
	"go_gin_second/common"
	"go_gin_second/common/dto"
	"go_gin_second/common/dto/response"
	"go_gin_second/model"
	"go_gin_second/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Info 获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("User")
	// ctx.JSON(
	// 	http.StatusOK,
	// 	gin.H{
	// 		"code": 200,
	// 		"data": gin.H{
	// 			"user": dto.ToUserDto(user.(model.User)),
	// 		},
	// 	},
	// )
	response.Success(ctx, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "success")
}

//Login 登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("tel")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "手机号必须11位",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}

	if len(password) < 6 {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "密码必须大于6位",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 { //用户不存在
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "用户不存在",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{
		// 	"code": 422,
		// 	"msg":  "密码错误",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseTOken(user)
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": 500,
		// 	"msg":  "系统异常",
		// })
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"msg":  "登录成功",
	// 	"data": gin.H{
	// 		"token": token,
	// 	},
	// })
	response.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")
}

//Register 注册
func Register(context *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := context.PostForm("name")
	telephone := context.PostForm("tel")
	password := context.PostForm("password")
	//数据验证

	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		// context.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "手机号必须11位",
		// })
		return
	}

	if len(password) < 6 {
		// context.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "密码必须大于6位",
		// })
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}

	//如果名称为空，给一个十位数的随机字符
	if len(name) == 0 {
		name = util.RandomName(10)

	}

	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		// context.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "用户已存在",
		// })
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//创建用户
	//1.创建用户时加密密码
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": 500,
		// 	"msg":  "加密错误",
		// })
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasPassword),
	}

	DB.Create(&newUser)

	//返回结果
	// context.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"msg":  "创建成功",
	// })
	response.Success(context, nil, "创建成功")
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
