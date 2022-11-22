package main

import (
	"EmployeeSelection/internal/api"
	"EmployeeSelection/internal/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	api.StartWebService()
}
