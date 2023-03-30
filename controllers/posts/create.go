package posts

import (
	"api-test/errorHandler"
	"api-test/initializers"
	"api-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body: body.Title,
	}
	result := initializers.DB.Create(&post)
	errorHandler.BadRequest(result.Error,c)

	c.JSON(http.StatusOK, gin.H{
		"post":post,
	})
}