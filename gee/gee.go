package gee

import (
	"net/http"
)

type HandlerFunc func (*Context)

type App struct {
	router *router
}

func New() *App {
	return &App{
		router: newRouter(),
	}
}



func (app *App) GET(pattern string, handlerFunc HandlerFunc) {
	app.router.addRoute(http.MethodGet, pattern, handlerFunc)
}

func (app *App) POST(pattern string, handlerFunc HandlerFunc) {
	app.router.addRoute(http.MethodPost, pattern, handlerFunc)
}

func (app *App) Run(addr string) error {
	return http.ListenAndServe(addr, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	app.router.handle(c)
}
