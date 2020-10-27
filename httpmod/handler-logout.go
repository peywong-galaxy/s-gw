package httpmod

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func logouthandler(c*gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"logout",
	})
}
