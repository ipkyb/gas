package api

type Route struct {
	Path     string
	Services []Service
}

type Service struct {
	Method   string
	Handlers []Handler
}
