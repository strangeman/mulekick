package mulekick

import "net/http"

func ExampleRouter_Use() {
	r := &Router{}

	r.Get("/hello")
	r.Use(func(w http.ResponseWriter, r *http.Request) {
		// sample middleware
	})
	r.Get("/world")

	// /hello call will not be affected by middleware
	// /world will have the middleware in its stack
}
