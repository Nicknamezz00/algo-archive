package dao

import (
	"algo-archive/internal/core"
	"algo-archive/internal/dao/jinzhu"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	ds   core.DataService
	once sync.Once
)

func DataService() core.DataService {
	once.Do(func() {
		// use gorm as orm for sql database
		ds = jinzhu.NewDataService()
		logrus.Infof("use gorm as data service")
	})
	return ds
}
