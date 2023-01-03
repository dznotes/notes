package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		// 设置 example 变量
		c.Set("example", "12345")

		// before request
		// 请求前
		c.Next()

		// after request
		// 请求后
		latency := time.Since(t)
		log.Print("latency:\t", latency)

		// access the status we are sending
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("status:\t", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	// http://127.0.0.1:8080/test
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
