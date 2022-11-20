package router

import (
	"data_feedback_progress/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get_article",components.GetArticle)
	r.Run(":3008") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
