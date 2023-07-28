package chain

import "AELS/ahttp"

type MiddlewareFunction func(ahttp.Handler) ahttp.Handler

type Chain struct {
	middlewares []MiddlewareFunction
}

func New(middlewares ...MiddlewareFunction) Chain {
	return Chain{append(([]MiddlewareFunction)(nil), middlewares...)}
}

func (c Chain) Then(h ahttp.Handler) ahttp.Handler {
	for i := range c.middlewares {
		h = c.middlewares[len(c.middlewares)-1-i](h)
	}

	return h
}
