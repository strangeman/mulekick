package mulekick

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// enables middleware
type Router struct {
	*mux.Router
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

func (r Router) Handle(endpoint string, mw ...http.HandlerFunc) *mux.Route {
	middleware := make([]http.HandlerFunc, len(r.middleware))
	for i, m := range r.middleware {
		middleware[i] = m
	}

	middleware = append(middleware, mw...)

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

func (r Router) Group(str string, middleware ...http.HandlerFunc) Router {
	middleware = append(r.middleware, middleware...)
	return New(r.PathPrefix(str).Subrouter(), middleware...)
}

func (r *Router) Use(middleware ...http.HandlerFunc) {
	r.middleware = append(r.middleware, middleware...)
}
