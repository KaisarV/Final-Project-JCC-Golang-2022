package main

import (
	"Final-Project-JCC-Golang-2022/docs"
	"Final-Project-JCC-Golang-2022/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	docs.SwaggerInfo.Title = "E-Commerce API"
	docs.SwaggerInfo.Description = "This API is used for online website."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := routes.SetupRouter()
	r.Run()
}
