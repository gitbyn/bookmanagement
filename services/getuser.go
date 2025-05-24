package services

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUserBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Im Inggrit as User!",
	})
}