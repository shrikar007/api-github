package main

import (
	"github-integration/app"
	"github-integration/drivers"
	"github-integration/helper"


	"github.com/sirupsen/logrus"
)
 
func main() {
	app.Initlogger()
	logrus.Print("Load configuration")
	app.InitConfig()
	configObj := drivers.GetConfig()
	logrus.Print("Start Application")
	appObj := new(app.App)
	appObj.DbInitialize(configObj)
	defer helper.Close(appObj)
	appObj.Run(":8000")
}
