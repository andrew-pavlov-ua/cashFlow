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

	ampqDSN := Getenv("AMQP_DSN", "amqp://guest:guest@localhost:5672/")

	a := app.NewApp(ampqDSN)
	a.Start()
}

func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Warnf("Environment variable %s not set, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
