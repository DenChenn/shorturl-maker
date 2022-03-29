package router

import (
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/controller"
	"github.com/SwarzChen/url-shortener/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Setting up basic cors config
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// Define routing path
	router.GET("/"+constants.CURRENT_VERSION+"/:url_id/", middleware.Cache(), controller.GetUrlHandler)
	router.POST("/"+constants.CURRENT_VERSION+"/urls/", controller.CreateUrlHandler)

	// Handling route not existing problem
	router.NoRoute(controller.MissingRouteHandler)
	return router
}
