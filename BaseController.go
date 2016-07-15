package gorest

type BaseController interface {
	Routes() []RoutePath
	IgnorePaths() []string
}
