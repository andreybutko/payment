package middleware

import (
	"github.com/codegangsta/negroni"
	metr "github.com/andreybutko/payment/pkg/metric"
	"net/http"
	"strconv"
	"time"
)


type responseWriterWithMetrics struct {
	http.ResponseWriter
	StatusCode int
	isHeaderWritten bool
	start           time.Time
}

func NewResponseWriterWithMetrics(w http.ResponseWriter) *responseWriterWithMetrics {
    return &responseWriterWithMetrics{w, http.StatusOK, false, time.Now()}
}

func (w *responseWriterWithMetrics) WriteHeader(statusCode int) {
	duration := time.Now().Sub(w.start)
	us := int(duration.Truncate(1000*time.Nanosecond).Nanoseconds() / 1000)
	w.Header().Set("X-Response-Time", strconv.Itoa(us)+"us")

	w.ResponseWriter.WriteHeader(statusCode)
	w.isHeaderWritten = true
}

func (w *responseWriterWithMetrics) Write(b []byte) (int, error) {
	if !w.isHeaderWritten {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(b)
}

// Metrics service to proses data
func Metrics(metricService metr.Service) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request,  next http.HandlerFunc) {
		next(NewResponseWriterWithMetrics(w), r)
	}
}