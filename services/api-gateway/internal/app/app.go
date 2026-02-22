package app

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	utils "github.com/andrew-pavlov-ua/pkg"
	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/handler"
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/server"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const HTTPPort = ":8080"
const GRPCPort = ":50051"

type App struct {
	Server       *http.Server
	GRPCServer   *grpc.Server
	GRPCListener net.Listener
	Handler      *handler.Handler
}

func NewApp() *App {
	amqpDSN := utils.Getenv("AMQP_DSN", "amqp://guest:guest@rabbitmq:5672/")

	a := &App{Handler: handler.NewHandler(amqpDSN)}
	a.InitHTTPServer()
	a.InitGRPCServer()

	return a
}

func (a *App) InitHTTPServer() {
	httpPort := utils.Getenv("REST_PORT", HTTPPort)

	r := gin.Default()

	r.POST("/new-transaction", a.Handler.NewTransaction)
	// more post requests can be added here

	// server/server-http init here to implement

	a.Server = &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: r,
	}
}

func (a *App) InitGRPCServer() {
	grpcPort := utils.Getenv("GRPC_PORT", GRPCPort)

	lis, err := net.Listen("tcp", "0.0.0.0:"+grpcPort)
	if err != nil {
		logger.Fatal(err)
	}

	a.GRPCListener = lis
	a.GRPCServer = grpc.NewServer()
	grpcServer := server.NewGRPCServer(a.Handler)
	grpcServer.RegisterServices(a.GRPCServer)

	logger.Infof("gRPC server initialized on %s", grpcPort)
}

func (a *App) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Infof("HTTP running on %s", a.Server.Addr)
		err := a.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Panic("error creating app instance: ", err)
		}
	}()

	go func() {
		logger.Infof("gRPC running on %s", a.GRPCListener.Addr().String())
		err := a.GRPCServer.Serve(a.GRPCListener)
		if err != nil {
			logger.Panic("error starting gRPC server: ", err)
		}
	}()
	defer a.Handler.RabbitMQPublisher.Close()

	// Wait for shutdown signal
	<-ctx.Done()
	logger.Info("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful stop
	a.GRPCServer.GracefulStop()
	_ = a.Server.Shutdown(shutdownCtx)

	logger.Info("Stopped cleanly")
}

func (a *App) Stop() {
	// This method can be used for additional cleanup if needed
}
