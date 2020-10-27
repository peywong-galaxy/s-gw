package httpmod

import (
	"github.com/gin-gonic/gin"
	"github.com/peywong-galaxy/s-gw/httpmod/api"
	"github.com/peywong-galaxy/s-gw/httpmod/gin-middleware"
)

func LaunchHttp() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	eng := gin.New()
	v1 := eng.Group("/api")
	v1.Use(gin_middleware.AddLogger())
	v1.Use(gin_middleware.AddSequence())
	{
		v1.GET("/register", api.Registerhandler)
		v1.GET("/login", api.Loginhandler)
		v1.GET("/logout", api.Logouthandler)
	}
	_ = eng.Run(":8080")
}
