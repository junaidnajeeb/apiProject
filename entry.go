package main

import (
	"apiProject/app"
	"apiProject/utils"
)

func main() {
	utils.LoggerInfo("entry => main called")
	/*
		utils.LoggerInfo("entry => main called")
		utils.LoggerDebug("entry => main called")
		utils.LoggerWarn("entry => main called")
		utils.LoggerTrace("entry => main called")
		utils.LoggerError("entry => main called")
		utils.LoggerFatal("entry => main called")
	*/
	app.SetupConfiguration()
	app.SetupDatabase()
	app.SetupRoutes()

}
