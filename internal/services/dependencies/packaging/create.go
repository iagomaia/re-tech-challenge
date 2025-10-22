package services

import (
	"context"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
)

type CreatePackagingDto struct {
	Size int
}

type ICreatePackagingRepo interface {
	Create(dto *CreatePackagingDto) (*models.Packaging, error)
	WithCtx(ctx context.Context) ICreatePackagingRepo
	Init()
}
