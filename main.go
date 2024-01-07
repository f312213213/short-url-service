package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"short-url/initialize"
	"short-url/url_records"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	initialize.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/:hash", url_records.GetUrlRecordByHash)
	r.POST("/", url_records.CreateUrlRecord)

	r.Run()
}
