package services

import (
	"context"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
)

type IGetPackagingRepo interface {
	GetAll() ([]*models.Packaging, error)
	GetAllSortedBySize() ([]*models.Packaging, error)
	WithCtx(ctx context.Context) IGetPackagingRepo
	Init()
}
