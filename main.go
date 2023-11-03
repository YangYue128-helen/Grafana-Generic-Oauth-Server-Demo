package main

import (
	"github.com/gin-gonic/gin"
	"mockAuth/routes"
)

func main() {
	router := gin.Default()

	// 配置路由
	routes.SetupRoutes(router)

	// 启动HTTP服务器，监听端口8080
	router.Run(":8080")
}