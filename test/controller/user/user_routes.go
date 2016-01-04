package user
import (
	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/test/controller"
)


type UserController struct {
	routes []gorest.RoutePath
}

func newUserController() *controller.RouteBase  {
	u := &UserController{}
	u.initRoutes()

	return  u
}

func (u *UserController) Routes() []gorest.RoutePath {
	return u.routes
}

func (u *UserController) initRoutes()  {
	u.routes = []gorest.RoutePath {
	}
}