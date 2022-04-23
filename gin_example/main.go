package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 基础
//func main() {
//	r := gin.Default()
//	r.GET("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "base gin example of json")
//		c.String(http.StatusOK, "base gin example of string")
//	})
//	r.Run(":9999")
//}
// 路径拼接请求
//func main() {
//	r := gin.Default()
//	r.GET("/user/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		action = strings.Trim(action, "/")
//		c.String(http.StatusOK, name+" is "+action)
//	})
//	r.Run(":9999")
//}
//  get请求
//func main() {
//	r := gin.Default()
//	r.GET("/user", func(c *gin.Context) {
//		name := c.Query("name")
//		defaultQuery := c.DefaultQuery("action", "my_action")
//		c.String(http.StatusOK, fmt.Sprintf("hello %s action is %s", name, defaultQuery))
//	})
//	r.Run(":9999")
//}
// post 请求 from 表单
//func main() {
//	r := gin.Default()
//	r.POST("/form", func(c *gin.Context) {
//		types := c.DefaultPostForm("type", c.Request.URL.Path)
//		msg := c.PostForm("msg")
//		c.JSON(http.StatusOK, fmt.Sprintf("types is %v, msg is %s", types, msg))
//	})
//	r.Run(":9999")
//}
// 上传单个文件
//func main() {
//	r := gin.Default()
//	r.MaxMultipartMemory = 2 << 10
//	r.POST("/upload", func(c *gin.Context) {
//		file, err := c.FormFile("file")
//		if err != nil {
//			c.String(500, "文件上传失败")
//		}
//		c.SaveUploadedFile(file, file.Filename)
//		c.String(200, "upload file success")
//	})
//	r.Run(":9999")
//}
// 上传多个文件
//func main() {
//	r := gin.Default()
//	r.POST("/upload", func(c *gin.Context) {
//		form, err := c.MultipartForm()
//		if err != nil {
//			c.String(500, fmt.Sprintf("upload files fail,reason：%s", err.Error()))
//		}
//		files := form.File["files"]
//		for _, file := range files {
//
//			if err := c.SaveUploadedFile(file, "./upload"+file.Filename); err != nil {
//				c.String(500, fmt.Sprintf("upload %s fail, resaon:%s", file, err.Error()))
//				return
//			}
//		}
//		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
//
//	})
//	r.Run(":9999")
//}
type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.String(200, fmt.Sprintf("username=%v password=%v", json.User, json.Password))
	})
	r.Run(":9999")

}
