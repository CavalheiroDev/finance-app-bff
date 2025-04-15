package server

import "github.com/cavalheirodev/finance-app-bff/pkg/error"

type Route struct {
	Method      string
	Prefix      string
	URI         string
	Function    func(webContext *fiberWebContext)
	BeforeEnter func(webContext *fiberWebContext) *error.ResponseError
}
