package gorest

type BaseController interface {
	Routes() []RoutePath
}
