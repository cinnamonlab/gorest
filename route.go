package gorest

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

}
