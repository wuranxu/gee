package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type App struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}

func New() *App {
	app := &App{
		router: newRouter(),
	}
	app.RouterGroup = &RouterGroup{app: app}
	app.groups = []*RouterGroup{app.RouterGroup}
	return app
}

func (app *App) GET(pattern string, handlerFunc HandlerFunc) {
	app.router.addRoute(http.MethodGet, pattern, handlerFunc)
}

func (app *App) POST(pattern string, handlerFunc HandlerFunc) {
	app.router.addRoute(http.MethodPost, pattern, handlerFunc)
}

func (app *App) Run(addr string) error {
	log.Printf("Listen on %s", addr)
	return http.ListenAndServe(addr, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range app.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	app.router.handle(c)
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	app         *App
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	app := group.app
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		app:    app,
	}
	app.groups = append(app.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	group.app.router.addRoute(method, pattern, handlerFunc)
}

func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handlerFunc)
}

func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handlerFunc)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
