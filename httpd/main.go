package main

import (
	"newfeeder/httpd/handler"

	"github.com/gin-gonic/gin"
	// "github.com/shubhamsnehi/newfeeder/httpd/handler"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet)
	r.Run()
}
