package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	systemDelivery "kabikha/app/system/delivery"
	systemRepo "kabikha/app/system/repository"
	systemUseCase "kabikha/app/system/usecase"
	"kabikha/infrastructure/config"
	"kabikha/infrastructure/db"
	"kabikha/infrastructure/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Serve() {
	// load application configuration
	if err := config.Load(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// connect to postgres DB
	if err := db.Connect(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	// http server setup
	e := echo.New()

	// fetch infra and routes, todo: add contexts
	if err := middlewares.Attach(e); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
	db := db.Get().DB

	// repository
	sysRepo := systemRepo.NewPgSystemRepository(db)

	// use cases
	sysUseCase := systemUseCase.NewSystemUsecase(sysRepo)

	// delivery
	systemDelivery.NewSystemHandler(e, sysUseCase)

	// start http server
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Get().App.Port)))
	}()

	// graceful shutdown setup
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	logrus.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	logrus.Infof("server shutdowns gracefully")
}
