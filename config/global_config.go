package config

import (
	"os"
)

const (
	//ScreenshotDB Schema Name for storing screenshots
	ScreenshotDB = "screenshots"
)

var (
	//DatabaseServer Address
	DatabaseServer = os.Getenv("MONGO_SERVER_ADDRESS")
)
