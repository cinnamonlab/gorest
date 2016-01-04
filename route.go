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

type APIFunc func(w *Response, r *Request, vars map[string]string)

var routeInstance *Route

// Singleton Route instance
func GetRouteInstance() *Route {
	if routeInstance == nil {
		routeInstance = &Route{}
	}
	return routeInstance
}

func checkPath(routePaths []string, requestPaths []string, restParams map[string]string) bool {

	isMatch := true
	if len(routePaths) != len(requestPaths) {
		return false
	} else {
		r1, _ := regexp.Compile("^:(.*)$")
		r2, _ := regexp.Compile("^{(.*)}$")
		for key, val := range routePaths {

			match1 := r1.FindStringSubmatch(val)
			match2 := r2.FindStringSubmatch(val)

			match := []string{}
			if len(match1) > 0 {
				match = match1
			} else if len(match2) > 0 {
				match = match2
			}

			if len(match) > 0 {
				restParams[match[1]]=requestPaths[key]
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

	if err := req.ParseForm(); err != nil {
		http.Error(rsp, err.Error(), http.StatusBadRequest)
		return
	}

	request := NewRequest(req)

	response := NewResponse(rsp)

	reqPath := path.Clean(req.URL.Path)

	pathElems := strings.Split(reqPath, "/")

	var matchRoute RoutePath

	isMatched := false

	var restParams map[string]string

	routePaths := route.controllers[req.Method]

	for _, controller := range routePaths {

		inputPrams := map[string]string{}

		if checkPath(controller.Paths, pathElems, inputPrams) {
			matchRoute = controller
			restParams = inputPrams
			isMatched=true
			break
		}
	}
	if !isMatched {
		// not found
		response.Text("URL Not Found!",400)
	} else {
		matchRoute.Controllers(response, request, restParams)
	}
}

//define RoutePath struct
type RoutePath struct {
	Method     string
	Path       string
	Paths      []string
	Controllers APIFunc
}

// register routes from controller instance
func (route *Route) AddRoute(routeBase BaseController) {
	routes := routeBase.Routes()
	if route.controllers == nil {
		route.controllers = make(map[string]map[string]RoutePath)
	}

	for _, r := range routes {
		_, ok := route.controllers[r.Method]
		if !ok {
			route.controllers[r.Method] = make(map[string]RoutePath)
		}
		route.controllers[r.Method][r.Path] = r
	}
}
