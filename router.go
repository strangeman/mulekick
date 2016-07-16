package mulekick

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	route      string
	middleware []http.HandlerFunc
}

func (r Router) Get(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("GET")
}

func (r Router) Post(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("POST")
}

func (r Router) Put(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("PUT")
}

func (r Router) Delete(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("DELETE")
}

func (r Router) Patch(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("PATCH")
}

func (r Router) Options(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	return r.Handle(endpoint, middleware...).Methods("OPTIONS")
}

func (r Router) Handle(endpoint string, mw ...http.HandlerFunc) *mux.Route {
	middleware := make([]http.HandlerFunc, len(r.middleware))
	for i, m := range r.middleware {
		middleware[i] = m
	}

	middleware = append(middleware, mw...)

	if os.Getenv("ENV") == "debug" {
		fmt.Printf("%v%v (%d handlers)\n", r.route, endpoint, len(middleware))
	}

	route := r.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		wr := &ResponseWriter{w, false, http.StatusOK, time.Now()}

		for _, m := range middleware {
			m(wr, r)

			if wr.responseWritten {
				break
			}
		}

		wr.LogMiddleware(r)
	})

	return route
}

// Group creates a new sub-router, enabling you to group handlers
func (r Router) Group(str string, middleware ...http.HandlerFunc) Router {
	middleware = append(r.middleware, middleware...)

	newRouter := New(r.PathPrefix(str).Subrouter(), middleware...)
	newRouter.route = r.route + str
	return newRouter
}

// Use function adds middleware to the router for calls
func (r *Router) Use(middleware ...http.HandlerFunc) {
	r.middleware = append(r.middleware, middleware...)
}
