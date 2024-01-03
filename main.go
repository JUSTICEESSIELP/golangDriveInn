package main

import (
	"driveinn_server/routes"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {

	godotenv.Load()
	irisApp := iris.Default()

	locationApi := irisApp.Party("api/v1/location")
	{

		locationApi.Get("/autocomplete", routes.Autocomplete)

		locationApi.Get("/search", routes.Search)

	}

	irisApp.Listen(":8080")

}
