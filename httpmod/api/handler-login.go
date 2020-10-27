package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Loginhandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"login",
	})
}


