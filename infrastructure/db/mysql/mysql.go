package mysql

import (
	"fmt"
	"time"
	"university/infrastructure/config"
	"university/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlClient struct {
	*gorm.DB
}

// db is the mysql instance
var db *mysqlClient

// Get returns the default mysqlClient currently in use
func Get() *mysqlClient {
	return db
}

// Connect database, must call once before server boot to Get() the db instance
func Connect() error {
	cnfg := config.Get().Database
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		cnfg.Username, cnfg.Password, cnfg.Host, cnfg.Port, cnfg.Name,
	)

	// setting log level of Databases
	logLevel := logger.Error
	if cnfg.Debug {
		logLevel = logger.Info
	}

	// open connection to mysql db
	instance, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{
		CreateBatchSize: cnfg.BatchSize,
		Logger:          logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return err
	}

	// connection pool settings
	dbInstance, err := instance.DB()
	if err != nil {
		return err
	}

	if cnfg.MaxLifeTime != 0 {
		dbInstance.SetConnMaxLifetime(cnfg.MaxLifeTime * time.Second)
	}
	if cnfg.MaxIdleConn != 0 {
		dbInstance.SetMaxIdleConns(cnfg.MaxIdleConn)
	}
	if cnfg.MaxOpenConn != 0 {
		dbInstance.SetMaxOpenConns(cnfg.MaxOpenConn)
	}

	// Model migration
	if err = instance.AutoMigrate(
		&model.Club{},
		&model.Department{},
		&model.Teacher{},
		&model.Student{},
	); err != nil {
		return err
	}

	db = &mysqlClient{DB: instance}
	return nil
}
