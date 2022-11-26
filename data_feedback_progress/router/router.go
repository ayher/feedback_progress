package router

import (
	"data_feedback_progress/public/common"
	"data_feedback_progress/public/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var router = common.GetRouter()

	r.Use(middleware.Validate(router))
	for k, v := range router {
		r.POST(k, common.GetComponent(v.Component))
	}

	r.Run(":3008") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
