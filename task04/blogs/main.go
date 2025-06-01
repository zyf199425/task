package main

import (
	"blogs/controller"
)

func main() {
	router := controller.StartWebServer()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
