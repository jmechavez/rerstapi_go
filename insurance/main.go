package main

import (
	"github.com/jmechavez/restapi_go/insurance/app"
	"github.com/jmechavez/restapi_go/insurance/logger"
)

func main() {
	logger.Info("Starting the application")

	app.Start()
}
