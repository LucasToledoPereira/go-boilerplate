package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router  *gin.Engine
	Version *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
	Admin   *gin.RouterGroup
}

func New() (r *Router) {
	//Configurate gin default router and cors
	router := gin.Default()
	router.Use(cors.New(getCorsConfig()))

	//Start version group and add middlewares to log info and recovery server in case of panic
	version := router.Group("v1/")

	//Create group for private routes and initizalize Auth0 middleware
	private := version.Group("private/")

	r = &Router{
		Router:  router,
		Version: version,
		Public:  version,
		Private: private,
		Admin:   version.Group("admin/"),
	}

	return r
}

/*
* Get Cors config to use with gin.Use
 */
func getCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	return config
}
