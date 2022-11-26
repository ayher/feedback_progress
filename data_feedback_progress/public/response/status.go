package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sucess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func False(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
