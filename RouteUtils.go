package gorest
import "strings"

type RouteUtils struct {
}

func (utils *RouteUtils) newRoutePath(method string, path string, controller APIFunc) RoutePath {
	paths := strings.Split(path,"/")
	return RoutePath{method: method, path : path, paths : paths, controller : controller}
}

// add new GET action
func (route *RouteUtils) createGetRoute(pattern string, controllerMethod APIFunc) RoutePath  {
	return route.createAction("GET",pattern,controllerMethod)
}
// add new POST action
func (route *RouteUtils) createPostRoute(pattern string, controllerMethod APIFunc) RoutePath  {
	return  route.createAction("POST",pattern,controllerMethod)
}
// add new PUT action
func (route *RouteUtils) createPutRoute(pattern string, controllerMethod APIFunc) RoutePath {
	return route.createAction("PUT",pattern,controllerMethod)
}
// add new delete action
func (route *RouteUtils) createDeleteRoute(pattern string, controllerMethod APIFunc) RoutePath {
	return route.createAction("DELETE",pattern,controllerMethod)
}

// perform add new action
func (route *RouteUtils) createAction(httpMethod string,pattern string, controllerMethod APIFunc) RoutePath  {
	return route.newRoutePath(httpMethod,pattern,controllerMethod)
}
