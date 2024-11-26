package storage

import "github.com/WeisseNacht18/gophermart/internal/config"

//тут у нас будет храниться storage в общем и целом

func NewStorage(config config.Config) {
	NewJWTStorage()
	NewDatabaseStorage(config.DatabaseUri)
}

func NewMockStorage() {
	NewJWTStorage()
	NewMockDatabaseStorage()
}
