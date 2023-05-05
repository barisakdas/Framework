package log

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/barisakdas/Framework/Log/Interfaces"
	. "github.com/barisakdas/Framework/Log/Models"
)

type LoggingMiddleware struct {
	logger  IFileLoggerService
	handler http.Handler
}

func (m *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log incoming request
	m.logger.LogRequest(RequestModel{
		Method: r.Method,
		Path:   r.URL.Path,
		Body:   string(getRequestBody(r)),
	})

	// Wrap the original http.Handler with a custom ResponseWriter to capture the response status code
	rw := &responseWriter{ResponseWriter: w}
	m.handler.ServeHTTP(rw, r)

	// Log outgoing response
	m.logger.LogResponse(ResponseModel{
		StatusCode: rw.statusCode,
		Body:       string(rw.body.Bytes()),
	})
}

func getRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.body.Write(body)
	return rw.ResponseWriter.Write(body)
}
