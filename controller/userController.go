package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golern/common"
	"golern/dto"
	"golern/model"
	"golern/response"
	"golern/util"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reflect"
)

// RegisterUser
// @Description 注册用户
// @Date 2022-07-26 21:15:35
// @param c
func RegisterUser(c *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	userData := model.User{}
	c.Bind(&userData)
	name := userData.Nickname
	telephone := userData.Telephone
	password := userData.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "手机号必须为11位", nil)
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "密码必须不小于6位", nil)
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, "手机号存在，不允许注册", nil)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, "系统错误："+err.Error(), nil)
		return
	}
	newUser := model.User{
		Nickname:  name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, "注册失败", nil)
		log.Printf("generate token failed: %v", err)
		return
	}
	response.Success(c, "注册成功", gin.H{"token": token})
}

func Login(c *gin.Context) {
	userData := model.User{}
	c.Bind(&userData)
	telephone := userData.Telephone
	password := userData.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "手机号码不小于11位", nil)
		return
	}

	// 判断手机号是否存在
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "用户不存在", nil)
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, "密码错误", nil)
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, "登录失败", nil)
		log.Printf("generate token failed: %v", err)
		return
	}
	response.Success(c, "登录成功", gin.H{"token": token})

}

// UserInfo
// @Description 获取用户信息
// @Date 2022-07-27 08:21:44
// @param c
func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	fmt.Println(reflect.TypeOf(user))
	response.Success(c, "操作成功", gin.H{"user": dto.ToUserDto(user.(model.User))})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
