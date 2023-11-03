package routes

import (
	"github.com/gin-gonic/gin"
	"mockAuth/controllers"
)

func SetupRoutes(router *gin.Engine) {
	oauthController := &controllers.OAuthController{}

	// OAuth 认证路由
	router.GET("/login/oauth/authorize", oauthController.Auth)
	router.POST("/login/oauth/token", oauthController.Token)
	router.GET("/userinfo", oauthController.UserInfo)
}
