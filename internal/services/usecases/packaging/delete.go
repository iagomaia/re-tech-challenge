package services

import (
	"context"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var (
	_ usecases.IDeletePackaging = (*DeletePackaging)(nil)
)

type DeletePackaging struct {
	DeleteRepo service.IDeletePackagingRepo
	ctx        context.Context
}

func (c *DeletePackaging) WithCtx(ctx context.Context) usecases.IDeletePackaging {
	return &DeletePackaging{
		DeleteRepo: c.DeleteRepo,
		ctx:        ctx,
	}
}

func (c *DeletePackaging) DeleteById(id string) error {
	return c.DeleteRepo.WithCtx(c.ctx).DeleteById(id)
}
