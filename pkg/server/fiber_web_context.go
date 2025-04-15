package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cavalheirodev/finance-app-bff/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Error string `json:"error"`
}

type WebContext interface {
	RequestHeader(key string) []string
	RequestHeaders() map[string][]string
	PathParam(key string) string
	QueryParam(key string) string
	DecodeQueryParams(value any) error
	DecodeBody(dto any) error
	AddHeader(key string, value string)
	JsonResponse(statusCode int, body any)
	ErrorResponse(statusCode int, err error)
}

type fiberWebContext struct {
	ctx *fiber.Ctx
}

func (f *fiberWebContext) RequestHeader(key string) []string {
	return []string{f.ctx.Get(key, "")}
}

func (f *fiberWebContext) RequestHeaders() map[string][]string {
	headers := make(map[string][]string)

	f.ctx.Context().Request.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = strings.Split(string(value), ";")
	})

	return headers
}

func (f *fiberWebContext) PathParam(key string) string {
	return f.ctx.Params(key)
}

func (f *fiberWebContext) QueryParam(key string) string {
	return f.ctx.Query(key)
}

func (f *fiberWebContext) DecodeQueryParams(value any) error {
	queryParams := make(map[string][]string)

	f.ctx.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		queryParams[string(key)] = strings.Split(string(value), ",")
	})

	if err := validator.FormDecode(value, queryParams); err != nil {
		return err
	}

	return validator.Struct(value)
}

func (f *fiberWebContext) DecodeBody(dto any) error {
	if err := json.Unmarshal(f.ctx.Body(), dto); err != nil {
		return err
	}

	return validator.Struct(dto)
}

func (f *fiberWebContext) AddHeader(key string, value string) {
	f.ctx.Response().Header.Add(key, value)
}

func (f *fiberWebContext) JsonResponse(statusCode int, body any) {
	f.ctx.Status(statusCode)
	if err := f.ctx.JSON(body); err != nil {
		f.ErrorResponse(http.StatusInternalServerError, err)
	}
}
func (f *fiberWebContext) ErrorResponse(statusCode int, err error) {
	f.JsonResponse(statusCode, Error{err.Error()})
}

func newFiberWebContext(ctx *fiber.Ctx) *fiberWebContext {
	return &fiberWebContext{ctx: ctx}
}
