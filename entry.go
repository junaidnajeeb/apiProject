package main

import (
	"apiProject/app"
	"apiProject/db"
	"apiProject/modal"
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
	db.SetupDatabase()
	//Create DB/Migrate
	db.GetDB().Debug().AutoMigrate(&modal.Account{})
	app.SetupRoutes()

}
