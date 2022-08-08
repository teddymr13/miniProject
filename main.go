package main

import (
	"invertory/app"
	"invertory/logger"
)

func main() {
	//*membuat setup run application
	logger.Info("start application...")
	app.Start()
}
