package server

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Degrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在gin上下文中定义变量
		c.Set("cache", "miss")

		// 请求前

		c.Next()

		// 请求后
	}
}
