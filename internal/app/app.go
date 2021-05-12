package app

import (
	"context"
	"github.com/labstack/echo/v4"
	_ "myapp/docs"
	"myapp/internal/delivery/http/v1"
	"myapp/internal/helpers"
	"myapp/internal/repository"
	"myapp/internal/server"
	"myapp/internal/service"
	"myapp/internal/validators"
	"myapp/pkg/database"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title jph_app
// @version 1.0
// @description This is a pet project of written in echo

// @host localhost:8001
// @BasePath /api/v1/

// @securityDefinitions.apikey UserAuth
// @in cookie
// @name Authorization

func InitApp() error {
	err := helpers.LoadConfig()
	if err != nil {
		return err
	}

	if helpers.LoadEnvVariables() != nil {
		return err
	}

	return nil
}

func Run() {
	e := echo.New()
	validators.InitValidator(e)

	err := InitApp()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	db, err := database.InitDatabase()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := v1.NewHandler(services)
	handlers.InitRouter(e)

	// start server
	go func() {
		if err := server.RunServer(e); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.StopServer(e, ctx); err != nil {
		e.Logger.Fatal(err)
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
