package main

import (
	"dollar-ussd/configs"
	"dollar-ussd/domain/migrations"
	"dollar-ussd/serve"
)

func main() {
	migrations.Migrations(configs.GetDB())
	serve.Serve()
}
