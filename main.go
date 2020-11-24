package main

import (
	"fmt"
	"gee/gee"
	"log"
)

func main() {
	app := gee.New()
	app.GET("/", func(c *gee.Context) {
		fmt.Println("你好牛逼哦")
		c.JSON(200, gee.H{
			"code": 200,
			"data": "你好",
			"msg":  "success",
		})
	})
	log.Fatal(app.Run(":9999"))
}
