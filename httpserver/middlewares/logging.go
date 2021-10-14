package middlewares

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out:          os.Stdout,
	Formatter:    &logrus.JSONFormatter{PrettyPrint: true},
	ReportCaller: true,
	Level:        logrus.InfoLevel,
}

func Logging(next http.Handler) http.Handler {
	// write your middleware definition here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request log
		logger.WithFields(logrus.Fields{
			"clientIp":      r.RemoteAddr,
			"path":          r.URL.Path,
			"method":        r.Method,
			"userAgent":     r.UserAgent(),
			"contentLength": r.ContentLength,
		}).Info("Incoming request")

		next.ServeHTTP(w, r)

		logger.WithFields(logrus.Fields{
			"statusCode":   w.Header().Get("statusCode"),
			"conentLength": w.Header().Get("contentLength"),
		}).Info("Outgoing response")
	})
}
