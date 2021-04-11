package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	systemDelivery "university/app/system/delivery"
	systemRepo "university/app/system/repo"
	systemUseCase "university/app/system/usecase"

	deptDelivery "university/app/department/delivery"
	deptRepo "university/app/department/repo"
	deptUseCase "university/app/department/usecase"

	teacherDelivery "university/app/teacher/delivery"
	teacherRepo "university/app/teacher/repo"
	teacherUseCase "university/app/teacher/usecase"

	"university/infrastructure/config"
	"university/infrastructure/db"
	"university/infrastructure/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Serve() {
	// load application configuration
	if err := config.Load(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// connect to database
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
	sysRepo := systemRepo.NewSystemRepository(db)
	deptRepo := deptRepo.NewDeptRepository(db)
	teacherRepo := teacherRepo.NewTeacherRepository(db)

	// use cases
	sysUseCase := systemUseCase.NewSystemUsecase(sysRepo)
	deptUseCase := deptUseCase.NewDeptUsecase(deptRepo)
	teacherUseCase := teacherUseCase.NewTeacherUsecase(teacherRepo)

	// delivery
	systemDelivery.NewSystemHandler(e, sysUseCase)
	deptDelivery.NewDeptHandler(e, deptUseCase)
	teacherDelivery.NewTeacherHandler(e, teacherUseCase)

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
