package internal

import "algo-archive/internal/service"

func Initialize() {
	// migrate database if needed

	service.Initialize()
}
