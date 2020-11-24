package gee

import (
	"fmt"
	"net/http"
)

type App struct {
	router map[string]http.HandlerFunc
}

func New() *App {
	return &App{
		router: make(map[string]http.HandlerFunc),
	}
}

func (app *App) addRoute(method string, pattern string, handlerFunc http.HandlerFunc) {
	key := method + "-" + pattern
	app.router[key] = handlerFunc
}

func (app *App) GET(pattern string, handlerFunc http.HandlerFunc) {
	app.addRoute(http.MethodGet, pattern, handlerFunc)
}

func (app *App) POST(pattern string, handlerFunc http.HandlerFunc) {
	app.addRoute(http.MethodPost, pattern, handlerFunc)
}

func (app *App) Run(addr string) error {
	return http.ListenAndServe(addr, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := app.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND")
	}
}
