package url_records

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"short-url/initialize"
	"short-url/utils"
)

type UrlRecord struct {
	Url  string
	Hash string
}

func GetUrlRecordByHash(c *gin.Context) {
	collection := initialize.Connection().Database("mongo").Collection("urls")

	hash := c.Param("hash")
	filter := bson.D{{"hash", hash}}

	var result UrlRecord
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": result.Url,
	})
}

func CreateUrlRecord(c *gin.Context) {
	collection := initialize.Connection().Database("mongo").Collection("urls")
	var requestBody UrlRecord
	c.Bind(&requestBody)
	requestBody.Hash = utils.GenerateShortLink(requestBody.Url)

	result, err := collection.InsertOne(context.TODO(), requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result.InsertedID,
	})
}
