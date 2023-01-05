package jinzhu

import (
	"algo-archive/internal/conf"
	"algo-archive/internal/core"
	"log"
)

type dataServant struct {
	core.UserManageService
}

var (
	_ core.DataService = (*dataServant)(nil)
)

func NewDataService() core.DataService {
	db := conf.InitGormDB()
	log.Println("Init gorm")
	// TODO: use cache.

	ds := &dataServant{
		UserManageService: newUserManageService(db),
	}
	return ds
}

func (s *dataServant) Name() string {
	return "Gorm"
}
