package main

import (
	"github.com/cavalheirodev/finance-app-bff/internal/entity/user"
	"github.com/cavalheirodev/finance-app-bff/pkg/config"
	"github.com/cavalheirodev/finance-app-bff/pkg/db"
	"github.com/cavalheirodev/finance-app-bff/pkg/server"
	"github.com/cavalheirodev/finance-app-bff/pkg/validator"
)

func main() {
	config.Load()
	validator.Initialize()
	db.Initialize()
	db.AutoMigrate(&user.User{})

	server.ListenAndServe()

}
