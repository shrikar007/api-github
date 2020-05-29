package main

import (
	"github-integration/app"
	"github.com/sirupsen/logrus"
)

func main() {
	apps,err := app.New()
	if err != nil {
		logrus.WithError(err).Error("Failed to start server")

	}
	apps.Run(":8000")
	defer apps.Close()
}

