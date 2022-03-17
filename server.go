package main

import (
	"github.com/SwarzChen/url-shortener/router"
	"github.com/SwarzChen/url-shortener/utils"
	"log"
)

func main() {
	// create required table
	err := utils.CreateTable()
	if err != nil {
		log.Fatalln(err)
	}

	engine := router.InitRoutes()
	panic(engine.Run())
}
