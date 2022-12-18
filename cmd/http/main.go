package main

import (
	"context"
	"log"
	"main/config"
	_ "main/docs"
	authServices "main/internal/auth/application/services"
	authHandlers "main/internal/auth/handler/http"
	authRepos "main/internal/auth/infrastructure/repository"
	"main/pkg/databases/postgresql"
	"main/pkg/logger"
	"main/pkg/server"
	"main/pkg/utils/graceful_exit"
)

// @title Auth Service
// @version 1.0
// @description Common Auth service broker with REST endpoints
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
	pgRepo := authRepos.NewPostgresqlRepository(postgresqlDB)

	// Init services
	authService := authServices.NewAuthService(cfg, pgRepo, appLogger)

	// Interceptors
	//

	servers := server.NewServer(cfg, &ctx, appLogger)

	httpServer, errHttpServer := servers.NewHttpServer()
	if errHttpServer != nil {
		println(errHttpServer.Error())
	}
	versioning := httpServer.Group("/v1")

	// Init handlers for HTTP Server
	authHandler := authHandlers.NewHttpHandler(ctx, cfg, authService, appLogger)

	// Init routes for HTTP Server
	authHandlers.MapRoutes(authHandler, versioning)

	//telegram.SendMessage("Send Message to telegram channel")

	// Exit from application gracefully
	graceful_exit.TerminateApp(ctx)

	appLogger.Info("Server Exited Properly")
}
