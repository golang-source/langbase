package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
)

func userCreateHandler(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"create user success"})

}

func userDeleteHandler(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"delete user success"})
}

func userModifyHandler(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"modify user success"})
}

func userQueryHandler(c *gin.Context) {
	username:=c.DefaultQuery("username","wenweiping")
	address:= c.Query("address")

	c.JSON(200,gin.H{"result":0,"des":"query user success","username":username,"address":address})
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": 1})
}

func main() {

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	ginpprof.Wrap(r)
	r.GET("/", rootHandler)

	r.POST("/user/info", userCreateHandler)

	r.DELETE("/user/info", userDeleteHandler)

	r.PUT("/user/info", userModifyHandler)

	r.GET("/user/info", userQueryHandler)
	r.Run(":9090")

}
