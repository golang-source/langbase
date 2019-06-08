package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
	"net/http"
	"fmt"
)

func userCreateHandlerN(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"create user success"})

}

func userDeleteHandlerN(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"delete user success"})
}

func userModifyHandlerN(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"modify user success"})
}

func userQueryHandlerN(c *gin.Context) {
	c.JSON(200,gin.H{"result":0,"des":"query user success"})
}

func userUploadHandler(c * gin.Context)  {
	imgfile,err:=c.FormFile("imgfile")
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}

	//目录一定要存在，否则存储失败
	c.SaveUploadedFile(imgfile,fmt.Sprintf("g:/img/%s",imgfile.Filename))

	c.JSON(http.StatusOK,gin.H{"message":fmt.Sprintf("%s is uploaded",imgfile.Filename)})

}

func main()  {
	gin.SetMode(gin.DebugMode)
	r:=gin.Default()
	ginpprof.Wrap(r)

	r.POST("/user/info/create",userCreateHandlerN)
	r.POST("/user/info/delete",userDeleteHandlerN)
	r.POST("/user/info/modify",userModifyHandlerN)
	r.POST("/user/info/query",userQueryHandlerN)


	r.POST("/user/tools/upload",userUploadHandler)

	r.Run(":9100")
}