package urls

import (
	"aha/views"
	"github.com/gin-gonic/gin"
)

func SetUpUrls() *gin.Engine {
	r := gin.Default()
	r.GET("/test", views.TestIndexView)
	return r
}