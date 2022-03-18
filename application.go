package main

import (
	"context"
	"github.com/VerTox/logger"
	"github.com/VerTox/zipcodes/api"
	v1 "github.com/VerTox/zipcodes/api/v1"
	"github.com/VerTox/zipcodes/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
	Connection domain.Connection
}

func (a *Application) Run(port string) {
	r := mux.NewRouter()

	ctx := &domain.Context{
		Connection: a.Connection,
	}

	v1.NewV1(r, ctx)

	r.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("418, I'm a teapot"))
	})

	r.Use(LoggingMiddleware, ContextMiddleware())

	panic(http.ListenAndServe(port, r))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")

		if requestId == "" {
			requestId = uuid.New().String()
			r.Header.Set("X-Request-Id", requestId)
		}

		l := logger.New(requestId)

		l.Info(r.Method+" "+r.RequestURI, nil)

		next.ServeHTTP(w, r)
	})
}

func ContextMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceId := r.Header.Get("X-Request-Id")

			ctx := &api.Context{
				TraceId: traceId,
			}

			c := context.WithValue(r.Context(), "context", ctx)

			next.ServeHTTP(w, r.WithContext(c))
		})
	}
}
