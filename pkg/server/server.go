package server

import (
	"log"

	"github.com/cavalheirodev/finance-app-bff/pkg/config"
)

var (
	serverRoutes []Route
	server       *fiberWebServer
)

type Server interface {
	initialize()
	shutdown() error
	injectMiddlewares()
	injectCustomMiddlewares()
	injectRoutes()
	listenAndServe() error
}

func AddRoutes(routes []Route) {
	serverRoutes = append(serverRoutes, routes...)
}

func ListenAndServe() {
	server = newFiberWebServer()
	server.initialize()
	server.injectRoutes()

	log.Printf("Service '%s' version %s running on port %d", config.APP_NAME, config.APP_VERSION, config.APP_PORT)
	if err := server.listenAndServe(); err != nil {
		log.Fatalf("Error rest server: %v", err)
	}
}
