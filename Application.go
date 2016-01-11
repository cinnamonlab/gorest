package gorest
import "net/http"

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

func (app *Application) StartHttpServer(port int) error  {
	err := http.ListenAndServe(port,application.Route);
	return err
}
