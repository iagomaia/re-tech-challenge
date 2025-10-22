package services

import (
	"context"
	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var (
	_ usecases.IGetPackagings = (*GetPackaging)(nil)
)

type GetPackaging struct {
	GetRepo service.IGetPackagingRepo
	ctx     context.Context
}

func (c *GetPackaging) WithCtx(ctx context.Context) usecases.IGetPackagings {
	return &GetPackaging{
		GetRepo: c.GetRepo,
		ctx:     ctx,
	}
}

func (c *GetPackaging) GetAll() ([]*models.Packaging, error) {
	return c.GetRepo.WithCtx(c.ctx).GetAll()
}
