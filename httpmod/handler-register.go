package httpmod

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerhandler(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"register",
	})
}
