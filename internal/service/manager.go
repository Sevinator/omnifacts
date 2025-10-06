package service

import (
	"omnifactory/internal/storage"
	// "log/slog"
)

type ManagerService struct {
	storages []storage.Storage
}

func New() ManagerService {
	manager := ManagerService{}

	return manager
}
