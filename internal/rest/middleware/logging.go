package middleware

import (
	"net/http"

	"github.com/strpc/resume-success/pkg/logging"
)

type LoggingMiddleware struct {
	logger *logging.Logger
}

func NewLoggingMiddleware(logger *logging.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (l *LoggingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.logger.WithFields(map[string]interface{}{
			"path":   r.RequestURI,
			"method": r.Method,
		}).Info("")
		next.ServeHTTP(w, r)
	})
}
