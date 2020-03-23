package main

import (
	"apiProject/app"
	"apiProject/greet"
	"fmt"
)

func main() {
	fmt.Println("entry => main called")
	fmt.Println(greet.Morning)

	app.SetupConfiguration()
	app.SetupRoutes()

}
