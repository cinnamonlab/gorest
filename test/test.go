package main

import (
	"net/http"

	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/test/controller/user"
)

func main() {
	route := gorest.GetRouteInstance()

	userController := user.NewController()

	route.AddRoute(userController)

	http.ListenAndServe(":3000", route)
}
