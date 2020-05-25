package main

import (
	"github-integration/app"
	"github-integration/config"
	"github-integration/helper"
	"log"
)
 
func main() {
	log.Println("Load configuration")
	app.InitConfig()
	configObj := config.GetConfig()
	log.Println("Start Application")
	appObj := new(app.App)
	appObj.DbInitialize(configObj)
	defer helper.Close(appObj)
	appObj.Run(":8000")
}
