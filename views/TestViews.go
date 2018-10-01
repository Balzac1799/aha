package views

import "github.com/gin-gonic/gin"

func TestIndexView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home Page Loaded",
	})
}