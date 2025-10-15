package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware){
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler{
	n := next

	for _, middleware := range middlewares{
		n = middleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(mux http.Handler) http.Handler{
	h := mux

	// Preflight, Cors, Logger,
	//logger(cors(Preflight(h)))
	for _, middleware := range mngr.globalMiddlewares{
		h = middleware(h)
	}

	return h
}