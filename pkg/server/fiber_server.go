package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/cavalheirodev/finance-app-bff/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type fiberWebServer struct {
	server *fiber.App
}

func newFiberWebServer() *fiberWebServer {
	return &fiberWebServer{}
}

func (f *fiberWebServer) initialize() {
	f.server = fiber.New(fiber.Config{
		ServerHeader:          config.APP_SERVER_HEADER,
		AppName:               config.APP_NAME,
		DisableStartupMessage: true,
	})
}

func (f *fiberWebServer) shutdown() error {
	return f.server.ShutdownWithTimeout(10 * time.Second)
}

func (f *fiberWebServer) convertUriToFiberUri(uri string, replacer *strings.Replacer) string {
	paths := strings.Split(uri, "/")

	for idx, path := range paths {
		if f.pathIsPathParam(path) {
			paths[idx] = fmt.Sprintf(":%s", replacer.Replace(path))
		}
	}

	return strings.Join(paths, "/")
}

func (f *fiberWebServer) pathIsPathParam(path string) bool {
	return strings.Contains(path, "{")
}

func (f *fiberWebServer) injectRoutes() {
	replacer := strings.NewReplacer(
		"{", "",
		"}", "",
	)

	for _, route := range serverRoutes {
		routeUri := string(route.Prefix) + f.convertUriToFiberUri(route.URI, replacer)
		fn := route.Function
		beforeEnter := route.BeforeEnter

		f.server.Add(route.Method, routeUri, func(ctx *fiber.Ctx) error {
			webContext := newFiberWebContext(ctx)
			if beforeEnter != nil {
				if err := beforeEnter(webContext); err != nil {
					ctx.Status(err.StatusCode)
					return ctx.JSON(err)
				}
			}

			fn(webContext)
			return nil
		})

		log.Info("Registered route [%7s] %s", route.Method, string(route.Prefix)+route.URI)
	}
}

func (f *fiberWebServer) listenAndServe() error {
	defer func() {
		if p := recover(); p != nil {
			log.Fatalf("panic recovering: %v", p)
		}
	}()

	addr := fmt.Sprintf(":%d", config.APP_PORT)
	return f.server.Listen(addr)
}
