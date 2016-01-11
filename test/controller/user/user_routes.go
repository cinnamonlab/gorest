package user

import (
	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/RouteUtils"
)

type UserController struct { // implement BaseController interface
	routes []gorest.RoutePath
}

func NewController() gorest.BaseController {
	u := &UserController{}
	u.initRoutes()

	return u
}

// implement function from baseController interface
func (u *UserController) Routes() []gorest.RoutePath {
	return u.routes
}

// init all routes for user
func (u *UserController) initRoutes() {


	u.routes = []gorest.RoutePath{
		RouteUtils.CreateGetRoute("/users/me",u.GetMySelf),
		RouteUtils.CreateGetRoute("/users/{id}",u.GetMyByID),
	}
}
