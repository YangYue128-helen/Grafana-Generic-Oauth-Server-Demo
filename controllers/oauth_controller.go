package controllers

import (
	"net/http"
	"strings"
	"log"

	"github.com/gin-gonic/gin"
	"mockAuth/utils"
)

type OAuthController struct{}

func (o *OAuthController) Auth(c *gin.Context) {
	
	// user := c.MustGet("user").(utils.User) // 请替换成正确的 User 结构
	redirectURI := c.DefaultQuery("redirect_uri", "http://localhost:3000/login/generic_oauth")
	state := c.DefaultQuery("state", "")
	// if user == nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"msg": "not login"})
	// 	return
	// }
	
	// url := redirectURI + "?state=" + state + "&code=" + user.ObjectID
	url := redirectURI + "?state=" + state + "&code=123456"
	log.Println(url)
	c.Redirect(http.StatusFound, url)
}

func (o *OAuthController) Token(c *gin.Context) {
	code := c.DefaultQuery("code", "123456")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}
	log.Println("Token return")
	token := &utils.OAuth2Token{
		AccessToken:  code,
		TokenType:    "Bearer",
		ExpiresIn:     1000000,
		RefreshToken: "",
	}

	c.JSON(http.StatusOK, token)
}

func (o *OAuthController) UserInfo(c *gin.Context) {
	headerAuth := c.GetHeader("Authorization")
	uid := strings.TrimPrefix(headerAuth, "Bearer ")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": ""})
		return
	}

	// 根据 uid 获取用户
	user, err := utils.GetUserByID(uid) // 请替换成正确的用户获取方法
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	nickName := user.NickName // 请替换成正确的用户属性
	email := user.Username + "@test" // 需要是电子邮箱格式

	c.JSON(http.StatusOK, gin.H{"email": email, "name": nickName})
}
