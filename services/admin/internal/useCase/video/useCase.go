package video

import (
	log "architecture_go/pkg/type/logger"
	repositoryStorage "architecture_go/services/admin/internal/repository/storage/postgres"
	"architecture_go/services/admin/internal/useCase/adapters/storage"
	"go.uber.org/zap"
)

type UseCase struct {
	adapterStorage storage.Video
	options        Options
}

type Options struct{}

func New(storage *repositoryStorage.Repository, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
		log.Info("set new options", zap.Any("options", uc.options))
	}
}
