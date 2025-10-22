package services

import (
	"context"
	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var (
	_ usecases.ICreatePackaging = (*CreatePackaging)(nil)
)

type CreatePackaging struct {
	CreateRepo service.ICreatePackagingRepo
	ctx        context.Context
}

func (c *CreatePackaging) WithCtx(ctx context.Context) usecases.ICreatePackaging {
	return &CreatePackaging{
		CreateRepo: c.CreateRepo,
		ctx:        ctx,
	}
}

func (c *CreatePackaging) Create(dto *usecases.CreatePackagingDto) (*models.Packaging, error) {
	dataDto := mapDomainDtoToRepoDto(dto)
	return c.CreateRepo.WithCtx(c.ctx).Create(dataDto)
}

func mapDomainDtoToRepoDto(dto *usecases.CreatePackagingDto) *service.CreatePackagingDto {
	return &service.CreatePackagingDto{
		Size: dto.Size,
	}
}
