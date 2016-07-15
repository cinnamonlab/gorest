package RouteUtils

import (
"github.com/cinnamonlab/gorest"
"strings"
)

func newRoutePath(method string, path string, controller gorest.APIFunc,ignorePaths []string) gorest.RoutePath {
	paths := strings.Split(path, "/")

	newPaths := make([]string,0)

	for _, value := range paths {
		if len(value)>0 {
			newPaths = append(newPaths,value)
		}
	}

	ignores := map[string]bool {}
	for _, path := range ignorePaths {
		ignores[path] = true
	}

	return gorest.RoutePath{Method: method, Path: path, Paths: newPaths, Controllers: controller, IgnorePaths: ignores}
}

// add new GET action
func CreateGetRoute(pattern string, controllerMethod gorest.APIFunc, ignorePaths []string) gorest.RoutePath {
	return createAction("GET", pattern, controllerMethod, ignorePaths)
}

// add new POST action
func CreatePostRoute(pattern string, controllerMethod gorest.APIFunc, ignorePaths []string) gorest.RoutePath {
	return createAction("POST", pattern, controllerMethod, ignorePaths)
}

// add new PUT action
func CreatePutRoute(pattern string, controllerMethod gorest.APIFunc, ignorePaths []string) gorest.RoutePath {
	return createAction("PUT", pattern, controllerMethod, ignorePaths)
}

// add new delete action
func CreateDeleteRoute(pattern string, controllerMethod gorest.APIFunc, ignorePaths []string) gorest.RoutePath {
	return createAction("DELETE", pattern, controllerMethod, ignorePaths)
}

// perform add new action
func createAction(httpMethod string, pattern string, controllerMethod gorest.APIFunc, ignorePaths []string) gorest.RoutePath {
	return newRoutePath(httpMethod, pattern, controllerMethod, ignorePaths)
}
