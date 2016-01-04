package gorest

type RouteContainer interface {
	initRoutes() []RoutePath;
}
