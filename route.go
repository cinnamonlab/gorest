package gorest

import (
	"net/http"
	//	"path"
	"path"
	"regexp"
	"strings"
)

type Route struct { // implement http.Handler interface
	controllers map[string]map[string]RoutePath
}

type APIFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string) error

var routeInstance *Route

// Singleton Route instance
func GetRouteInstance() *Route {
	if routeInstance == nil {
		routeInstance = &Route{}
	}
	return routeInstance
}

func checkPath(routePaths []string, requestPaths []string, restParams *map[string]string) bool {

	isMatch := true
	if len(routePaths) != len(requestPaths) {
		return false
	} else {
		for key, val := range routePaths {
			r1, _ := regexp.Compile("^:(.*)$")
			r2, _ := regexp.Compile("^(.*)}$")

			match1 := r1.FindStringSubmatch(val)
			match2 := r2.FindStringSubmatch(val)

			match := []string{}
			if len(match1) >= 0 {
				match = match1
			} else if len(match2) >= 0 {
				match = match2
			}

			if len(match) > 0 {
				//				restParams[match[1]]=requestPaths[key]
			} else {
				if requestPaths[key] != val {
					return false
				}
			}
		}
		return isMatch
	}
}

func (route *Route) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	reqPath := path.Clean(req.URL.Path)

	pathElems := strings.Split(reqPath, "/")

	var matchRoute RoutePath

	var restParams map[string]string

	routePaths := route.controllers[req.Method]

	for _, controller := range routePaths {

		inputPrams := map[string]string{}

		if checkPath(controller.paths, pathElems, &inputPrams) {
			matchRoute = controller
			//			restParams = &inputPrams
			break
		}
	}
	if &matchRoute == nil {
		// not found
		//rsp.Write("Error!")
	} else {
		matchRoute.controller(rsp, req, restParams)
	}
}

//define RoutePath struct
type RoutePath struct {
	method     string
	path       string
	paths      []string
	controller APIFunc
}

func (route *Route) addRoute(routeBase *BaseController) {
	routes := routeBase.Routes()

	for _, r := range routes {
		val, ok := route.controllers[r.method]
		if !ok {

		}
	}

}

/*func (routePath *RoutePath) splitPaths() {
	routePath.paths = strings.Split(path,"/")
}*/
