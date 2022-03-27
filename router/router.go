package router

import (
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
	config.AddAllowHeaders("Authorization", "sentry-trace")
	router.Use(cors.New(config))

	// Define routing path
	router.GET("/v1/:url_id/", middleware.Cache(), controller.GetUrlHandler)
	router.POST("/v1/urls/", controller.CreateUrlHandler)

	// Handling route not existing problem
	router.NoRoute(controller.MissingRouteHandler)
	return router
}
