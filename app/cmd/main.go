package main

import "github.com/gulizay91/go-rest-api/cmd/services"

// @title Go Rest API Starter Doc
// @version 1.0
// @description Go - RESTful
// @termsOfService https://swagger.io/terms/

// @contact.name GulizAY
// @contact.url https://github.com/gulizay91
// @contact.email gulizay91@gmail.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8091
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Bearer-Token
func main() {
	services.Run()
}
