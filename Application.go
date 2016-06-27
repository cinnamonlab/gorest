package gorest

import (
	"fmt"
	"net/http"

	"github.com/cinnamonlab/gorest/queue"
)

// Applicaton packae export
type Application struct {
	Config interface{} // the application config
	Route  *Route
}

var application *Application

func SharedApplication() *Application {
	if application == nil {
		application = &Application{}

		application.Init()
	}
	return application
}

func (app *Application) Init() {
	app.Route = GetRouteInstance()
}

func (app *Application) StartHttpServer(port string) {
	err := http.ListenAndServe(port, application.Route)

	if err != nil {
		fmt.Println("Error:" + err.Error())
		panic(err)
	} else {
		fmt.Println("Http Server started at http://localhost:" + port)
	}
}

func (app *Application) StartQueue(numberProcess int) {
	queue.Start(numberProcess)
}
