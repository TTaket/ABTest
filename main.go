package main

import (
	"ABTest/api/routes"
	_ "ABTest/docs" // 导入swagger文档
	"fmt"
	"log"
)

func main() {
	// 初始化路由
	router := routes.SetupRouter()

	// 启动HTTP服务器
	port := 8080
	serverAddr := fmt.Sprintf(":%d", port)
	log.Printf("HTTP服务器启动在 http://localhost%s", serverAddr)
	log.Printf("Swagger文档地址: http://localhost%s/swagger/index.html", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
