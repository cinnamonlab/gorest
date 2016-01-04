package RouteUtils

import (
"github.com/cinnamonlab/gorest"
"strings"
)

func newRoutePath(method string, path string, controller gorest.APIFunc) gorest.RoutePath {
	paths := strings.Split(path, "/")

	return gorest.RoutePath{Method: method, Path: path, Paths: paths, Controllers: controller}
}

// add new GET action
func CreateGetRoute(pattern string, controllerMethod gorest.APIFunc) gorest.RoutePath {
	return createAction("GET", pattern, controllerMethod)
}

// add new POST action
func CreatePostRoute(pattern string, controllerMethod gorest.APIFunc) gorest.RoutePath {
	return createAction("POST", pattern, controllerMethod)
}

// add new PUT action
func CreatePutRoute(pattern string, controllerMethod gorest.APIFunc) gorest.RoutePath {
	return createAction("PUT", pattern, controllerMethod)
}

// add new delete action
func CreateDeleteRoute(pattern string, controllerMethod gorest.APIFunc) gorest.RoutePath {
	return createAction("DELETE", pattern, controllerMethod)
}

// perform add new action
func createAction(httpMethod string, pattern string, controllerMethod gorest.APIFunc) gorest.RoutePath {
	return newRoutePath(httpMethod, pattern, controllerMethod)
}
