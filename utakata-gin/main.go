package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(checkLoggedIn())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

func checkLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("ログイン確認処理 前")
		ctx.Next()
		fmt.Println("ログイン確認処理 後")
	}
}