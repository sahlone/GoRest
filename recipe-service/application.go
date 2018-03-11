package main

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/http"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"net"
	"os"
	"strconv"
)

func main() {
	app := http.Initialize()
	port := os.Getenv("PORT")
	if len(port) < 4 {
		logger.Fatal("Error:Port environment variable not found")
	} else if val, err := strconv.ParseInt(port, 10, 16); err != nil {
		logger.Fatal("Error:Port environment variable not found")
	} else if val < 1024 || val > 65536 {
		logger.Fatal("Error:Port environment variable not found")
	}
	port = net.JoinHostPort("", port)
	logger.Info("Starting container with binding : %v", port)
	app.Run(port)
}
