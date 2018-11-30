package main

import (
	"database/sql"
	"sync"

	"github.com/sadewoeko/scheduler/config"
	"github.com/sadewoeko/scheduler/driver"
	"github.com/sadewoeko/scheduler/duler"
)

type DepedencyInjector interface {
	InjectDuler() duler.DulerController
}

type depedency struct{}

func (d *depedency) InjectDuler() duler.DulerController {

	sqlConn, _ := sql.Open("mysql", config.MYSQL_HOST)
	DBHandler := &driver.MySQLHandler{}
	DBHandler.Conn = sqlConn

	dulerRepository := &duler.DulerRepository{DBHandler}
	dulerService := &duler.DulerService{dulerRepository}
	dulerController := duler.DulerController{dulerService}

	return dulerController
}

func Depedency() DepedencyInjector {

	var (
		d    *depedency
		once sync.Once
	)

	if d == nil {
		once.Do(func() {
			d = &depedency{}
		})
	}
	return d
}
