package middelware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// responseWriterWrapper перехватывает код ответа
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем writer для перехвата statusCode
		ww := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		statusCode := ww.statusCode

		// Определяем наличие ошибки
		var errMsg string
		if ctxErr := r.Context().Err(); ctxErr != nil {
			errMsg = fmt.Sprintf("context_cancelled: %v", ctxErr)
		} else if statusCode >= 500 {
			errMsg = fmt.Sprintf("server_error_%d", statusCode)
		}

		// Логирование
		slog.Info("request_completed",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status_code", statusCode),
			slog.Duration("duration", duration),
			slog.String("ip", r.RemoteAddr),
			slog.Any("error", errMsg),
		)
	})
}
