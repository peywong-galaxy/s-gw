package httpmod

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginhandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"login",
	})
}
