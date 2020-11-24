package main

import (
	"fmt"
	"gee/gee"
	"log"
	"net/http"
)

func main() {
	app := gee.New()
	app.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("你好牛逼哦")
		fmt.Fprintf(writer, "welcome")
	})
	log.Fatal(app.Run(":9999"))
}
