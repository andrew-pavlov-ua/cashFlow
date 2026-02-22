package main

import (
	"os"

	"github.com/andrew-pavlov-ua/pkg/logger"
	app "github.com/andrew-pavlov-ua/services/api-gateway/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// 3. Initialize logger BEFORE using it
	if err := logger.Init("api-gateway", env); err != nil {
		panic(err) // Can't log yet, so panic
	}
	defer logger.Sync()

	// 4. NOW you can use logger
	logger.Info("Starting API Gateway...")
	logger.Infof("Environment: %s", env)

	a := app.NewApp()
	a.Start()
}
