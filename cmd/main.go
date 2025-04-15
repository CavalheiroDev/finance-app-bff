package main

import (
	"os/user"

	"github.com/cavalheirodev/finance-app-bff/cmd/db"
	"github.com/cavalheirodev/finance-app-bff/pkg/web/config"
	"github.com/cavalheirodev/finance-app-bff/pkg/web/server"
	"github.com/cavalheirodev/finance-app-bff/pkg/web/validator"
)

func main() {
	config.Load()
	validator.Initialize()
	db.Initialize()
	db.AutoMigrate(&user.User{})

	server.ListenAndServe()

}
