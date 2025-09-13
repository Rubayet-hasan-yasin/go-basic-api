package middleware

import "ecommerce/config"

type Middlewares struct {
	Cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		Cnf: cnf,
	}
}
