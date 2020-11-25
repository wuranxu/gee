package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func main() {
	app := gee.New()
	app.GET("/oo/:xixixi", func(context *gee.Context) {
		param := context.Param("xixixi")
		fmt.Println(param)
		context.JSON(200, gee.H{
			"data": "别吵了 特么饿的",
		})
	})
	app.GET("/nidaye", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"code": 200,
			"data": "你好",
			"msg":  "success",
		})
	})

	v1 := app.Group("/admin")
	v1.GET("/panic", func(c *gee.Context) {
		name := []string{"2"}
		c.String(http.StatusOK, name[223])
	})
	v1.GET("/hello", func(context *gee.Context) {
		context.String(http.StatusOK, "hello, 小弟弟")
	})
	app.Use(gee.Logger(), gee.Recovery())
	app.Run(":9999")
}
