package httpmod

import "github.com/gin-gonic/gin"

func LaunchHttp() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	eng := gin.New()
	v1 := eng.Group("/api")
	v1.Use(AddLogger())
	v1.Use(AddSequence())
	{
		v1.GET("/register", registerhandler)
		v1.GET("/login", loginhandler)
		v1.GET("/logout", logouthandler)
	}
	_ = eng.Run(":8080")
}
