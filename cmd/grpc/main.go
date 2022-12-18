package main

import (
	"context"
	"google.golang.org/grpc/reflection"
	"log"
	"main/config"
	authRepos "main/internal/auth/infrastructure/repository"
	"main/pkg/databases/postgresql"
	"main/pkg/logger"
	"main/pkg/server"
	"main/pkg/utils/graceful_exit"
)

// @title Auth Service
// @version 1.0
// @description Common Auth service broker with GRPC endpoints
// @contact.email ivanbarayev@hotmail.com
// @BasePath /v1
func main() {
	log.Println("Starting api server")

	cfg, errConfig := config.ParseConfig()
	if errConfig != nil {
		log.Fatal(errConfig)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.APP_VERSION, cfg.Logger.LEVEL, cfg.Server.APP_ENV)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init Clients
	postgresqlDB, err := postgresql.NewPostgresqlDB(cfg)
	if err != nil {
		appLogger.Fatal("Error when tyring to connect to Postgresql")
	} else {
		appLogger.Info("Postgresql connected")
	}

	// Init repositories
	_ = authRepos.NewPostgresqlRepository(postgresqlDB)

	// Init services
	//

	// Interceptors
	//

	servers := server.NewServer(cfg, &ctx, appLogger)

	// Init handlers for HTTP Server
	//

	// Init routes for HTTP Server
	//

	//telegram.SendMessage("Send Message to telegram channel")

	// GRPC Services
	grpcServer, errGrpcServer := servers.NewGrpcServer()
	if errGrpcServer != nil {
		cancel()
		return
	}

	if cfg.Server.APP_ENV == "dev" {
		reflection.Register(grpcServer)
	}

	// Exit from application gracefully
	graceful_exit.TerminateApp(ctx)

	grpcServer.GracefulStop()
	appLogger.Info("Server Exited Properly")
}
