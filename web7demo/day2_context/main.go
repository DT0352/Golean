package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s ,you are at %s\n", c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})
	r.Run(":9999")
}
