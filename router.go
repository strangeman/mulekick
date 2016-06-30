package mulekick

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type H map[string]interface{}

func Bind(w http.ResponseWriter, r *http.Request, out interface{}) error {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	return err
}

func WriteJSON(w http.ResponseWriter, code int, out interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(out); err != nil {
		panic(err)
	}
}

// enables middleware
type Router struct {
	*mux.Router
	middleware []http.HandlerFunc
}

func NewRouter(r *mux.Router, middleware ...http.HandlerFunc) Router {
	return Router{
		Router:     r,
		middleware: middleware,
	}
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

func (r Router) Handle(endpoint string, middleware ...http.HandlerFunc) *mux.Route {
	middleware = append(r.middleware, middleware...)

	route := r.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		wr := NewResponseWriter(w)

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
	return NewRouter(r.PathPrefix(str).Subrouter(), middleware...)
}

func (r *Router) Use(middleware ...http.HandlerFunc) {
	r.middleware = append(r.middleware, middleware...)
}
