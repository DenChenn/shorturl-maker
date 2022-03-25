package main

import (
	"github.com/SwarzChen/url-shortener/router"
)

func main() {
	engine := router.InitRoutes()
	panic(engine.Run())
}
