package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/queue"
	"github.com/cinnamonlab/gorest/test/controller/user"
)

func main() {
	queue.Start(10)

	queue.NewTask(func(p ...interface{}) {
		fmt.Println(p)
		fmt.Println("Hello " + p[0].(string))
	}, "world")

	route := gorest.GetRouteInstance()

	userController := user.NewController()

	route.AddRoute(userController)

	http.ListenAndServe(":3000", route)

	log.Print("server is running at http://localhost:3000/")
}
