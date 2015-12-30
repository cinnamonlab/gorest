package gorest
import (
	"net/http"
//	"path"
	"strings"
	"regexp"
	"path"
)

type Route struct {
	controllers map[string]RoutePath
}

type HttpMethod string

const (
	HTTP_GET    HttpMethod ="GET"
	HTTP_POST   HttpMethod ="POST"
	HTTP_PUT    HttpMethod ="PUT"
	HTTP_DELETE HttpMethod ="DELETE"
)

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
	if(len(routePaths)!=len(requestPaths)) {
		return  false;
	} else  {
		for key, val := range routePaths  {
			r1,_ := regexp.Compile("^:(.*)$")
			r2,_ := regexp.Compile("^(.*)}$")

			match1 := r1.FindSubmatch(val);
			match2 := r2.FindSubmatch(val);

			match :=[]string{}
			if(len(match1)>=0) {
				match := match1
			} else if(len(match2)>=0) {
				match := match2
			}

			if(len(match)>0) {
				restParams[match[1]]=requestPaths[key]
			} else  {
				if(requestPaths[key]!=val) {
					return false
				}
			}
		}
		return  isMatch
	}
}

func (route *Route) ServeHTTP(rsp http.ResponseWriter,req *http.Request)  {
	reqPath := path.Clean(req.URL.Path)

	pathElems := strings.Split(reqPath,"/")

	matchRoute := RoutePath{}

	restParams :=[]string{}

	for _, controller := range route.controllers {

		inputPrams :=[]string{}

		if(checkPath(controller.paths,pathElems,inputPrams)) {
			matchRoute=controller
			restParams = &inputPrams
			break;
		}
	}


}

// add new GET action
func (route *Route) GET(pattern string, controllerMethod string)  {
	route.Action(HTTP_GET,pattern,controllerMethod);
}
// add new POST action
func (route *Route) POST(pattern string, controllerMethod string)  {
	route.Action(HTTP_POST,pattern,controllerMethod);
}
// add new PUT action
func (route *Route) PUT(pattern string, controllerMethod string)  {
	route.Action(HTTP_PUT,pattern,controllerMethod);
}
// add new delete action
func (route *Route) DELETE(pattern string, controllerMethod string)  {
	route.Action(HTTP_DELETE,pattern,controllerMethod);
}
// perform add new action
func (route *Route) Action(httpMethod HttpMethod,pattern string, controllerMethod string)  {
	routePath := newRoutePath(pattern,controllerMethod)

	if route.controllers==nil {
		route.controllers = make(map[string]RoutePath)
	}

	route.controllers[pattern]=routePath
}

//define RoutePath struct
type RoutePath struct {
	paths []string
	controller string
}

//define new RoutePath
func newRoutePath(path string, controller string) RoutePath {
	paths := strings.Split(path,"/")

	return RoutePath{paths,controller}
}
