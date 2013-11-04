package main

import (
	"./controllers"
	"./controllers/middleware"
	"./controllers/authentication"
	"./helpers/database"
	"flag"
	"github.com/ninnemana/web"
	"log"
)

var (
	listenAddr = flag.String("port", "8087", "http listen address")
)

func main() {
	flag.Parse()

	err := database.PrepareAll()
	if err != nil {
		log.Fatal(err)
	}

	web.Middleware(middleware.Base)


	//authentication
	web.Get("/login", authentication.Index)
	web.Post("/login", authentication.Login)
	web.Get("/logout", authentication.Logout)


	web.Get("/", controllers.Index)
	web.Run("127.0.0.1:" + *listenAddr)
}
