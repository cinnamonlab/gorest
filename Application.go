package gorest
import (
	"net/http"
	"fmt"
)

type Application struct {
	Config interface{} // the application config
	Route *Route
}

var application *Application

func SharedApplication() *Application {
	if application == nil {
		application = &Application{}

		application.Init()
	}
	return application
}

func (app *Application) Init()  {
	app.Route = GetRouteInstance()
}

func (app *Application) StartHttpServer(port string)  {
	err := http.ListenAndServe(port,application.Route);

	if(err!=nil) {
		fmt.Println("Error:"+err.Error())
		panic(err)
	} else  {
		fmt.Println("Http Server started at http://localhost:" + port);
	}
}
