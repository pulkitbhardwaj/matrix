package runtime

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Method string

const (
	GET  Method = http.MethodGet
	POST Method = http.MethodPost
)

type Middleware func(http.Handler) http.Handler

type Router interface {
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Use(middlewares ...func(http.Handler) http.Handler)
}

func NewRouter() Router {
	return chi.NewRouter()
}
