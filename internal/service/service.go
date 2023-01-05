package service

import (
	"algo-archive/internal/core"
	"algo-archive/internal/dao"
)

var ds core.DataService

func Initialize() {
	ds = dao.DataService()
}
