package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	ginpprof.Wrap(r)

	r.GET("/hello.html", func(context *gin.Context) {

	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping/go", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"haha",
		})
	})


	r.Run(":8000")
}