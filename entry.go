package main

import (
	"apiProject/app"
	"fmt"
)

func main() {
	fmt.Println("entry => main called")

	app.SetupConfiguration()
	app.SetupDatabase()
	app.SetupRoutes()

}
