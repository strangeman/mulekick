package mulekick

// fake response writer wrapper to tell us when middleware writes
import (
	"net/http"
	"time"
)

type ResponseWriter struct {
	ResponseWriter  http.ResponseWriter
	responseWritten bool
	statusCode      int
	start           time.Time
}

func (wr *ResponseWriter) Header() http.Header {
	return wr.ResponseWriter.Header()
}

func (wr *ResponseWriter) Write(b []byte) (int, error) {
	wr.responseWritten = true
	return wr.ResponseWriter.Write(b)
}

func (wr *ResponseWriter) WriteHeader(header int) {
	wr.responseWritten = true
	wr.statusCode = header
	wr.ResponseWriter.WriteHeader(header)
}
