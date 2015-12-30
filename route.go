package gorest
import (
	"net/http"
	"path"
	"strings"
	"fmt"
)

type Route struct {
	controllers map[string]string
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

func (route *Route) ServeHTTP(rsp http.ResponseWriter,req *http.Request)  {
	reqPath := path.Clean(req.URL.Path)

	pathElems := strings.Split(reqPath,"//")

	fmt.Println(pathElems)
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
func newRoutePath(path string, controller string) *RoutePath {
	paths := strings.Split(path,"//")

	return &RoutePath{paths,controller}
}
