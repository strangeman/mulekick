package mulekick

import (
	"net/http"
	"time"
)

func PongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		CorsMiddleware(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	wr := &ResponseWriter{w, false, http.StatusOK, time.Now()}
	wr.WriteHeader(http.StatusNotFound)
	wr.Write([]byte("404 not found"))

	wr.LogMiddleware(r)
}

func CorsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method != "OPTIONS" {
		return
	}

	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	w.WriteHeader(http.StatusNoContent)
}
