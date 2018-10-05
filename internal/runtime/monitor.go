/*
	Web Service, provide other to look the available event queue, history ... etc
*/

package monitor

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

// Monitoring server
type Handler func(*Context)

type Route struct {
	Pattern *regexp.Regexp 
	Handler Handler
}

type Monitor struct {
	Routes			[]Route 
	DefaultRoute 	Handler
}

// Monitor Server method
func CreateMonitor() *Monitor {
	app := &Monitor{
		DefaultRoute: func(ctx *Context){
			ctx.Text(http.StatusNotFound, "text/plain", "Not Found")
		},
	}

	return app 
}

func (m *Monitor) Handle(pattern string, handler Handler){
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler}

	m.Routes = append(m.Routes, route)
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{Request: r, ResponseWriter: w}

	for _, rt := range m.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return 
		}
	}

	m.DefaultRoute(ctx)
}

type Context struct {
	http.ResponseWriter
	*http.Request 
	Params []string
}

func (c *Context) Text(code int, content_type, body string){
	// FIXME: content-type
	c.ResponseWriter.Header().Set("Content-Type", content_type)
	c.WriteHeader(code)

	io.WriteString(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}

// Running Monitor Server - Example 
func Run() {
	app := CreateMonitor()

	// Handler 
	app.Handle(`^/hello$`, func(ctx *Context){
		ctx.Text(http.StatusOK, "text/plain","Hello World")
	})
	app.Handle(`/hello/([\w\._-]+)$`, func(ctx *Context){
		// TODO: subpath
		ctx.Text(http.StatusOK, "text/plain", fmt.Sprintf("Hello %s", ctx.Params[0]))
	})

	// Run on specific port
	err := http.ListenAndServe(":9000", app)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}