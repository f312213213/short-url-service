package main

import (
	"github.com/gin-gonic/gin"
	"short-url/initialize"
	"short-url/url_records"
)

func init() {
	initialize.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/:hash", url_records.GetUrlRecordByHash)
	r.POST("/", url_records.CreateUrlRecord)

	r.Run()
}
