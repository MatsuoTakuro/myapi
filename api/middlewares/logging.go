package middlewares

import (
	"log"
	"net/http"

	"github.com/MatsuoTakuro/myapi-go-intermediate/api/contexts"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := contexts.NewTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)

		ctx := contexts.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
