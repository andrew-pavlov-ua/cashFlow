package main

import "github.com/andrew-pavlov-ua/pkg/logger"

func main() {
	err := logger.Init("notification-service", "dev")
	if err != nil {
		panic(err)
	}

	logger.Info("Starting Notification Service...")
}
