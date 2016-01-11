package main

import (
	"net/http"

	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/test/controller/user"
	"log"
)

func main() {
	route := gorest.GetRouteInstance()

	userController := user.NewController()

	route.AddRoute(userController)

	http.ListenAndServe(":3000", route)

	log.Print("server is running at http://localhost:3000/")
}
