package gorest

type Application struct {
	Config interface{} // the application config

}

var application *Application

func SharedApplication()  {
	if application == nil {
		application = &Application{}
	}
	return application
}

func (app *Application) Init()  {

}
