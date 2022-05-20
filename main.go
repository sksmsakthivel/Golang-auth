package main

import (
	"goauth/config"
	"goauth/routes"
)

func main() {
	config.SetUpDB()
	r := routes.SetUpRoutes()

	r.Run()

}
