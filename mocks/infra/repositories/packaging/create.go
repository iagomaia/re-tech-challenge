package repositoriesmocks

import (
	"context"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	mocks "github.com/iagomaia/re-tech-challenge/mocks/domain/models"
)

var (
	_ service.ICreatePackagingRepo = (*CreatePackagingMockRepo)(nil)
)

type CreatePackagingMockRepo struct {
	CreateMethodCalledTimes int
	CreateMethodReturn      *models.Packaging
	CreateMethodError       error
}

func (r *CreatePackagingMockRepo) Init() {
	r.CreateMethodCalledTimes = 0
	r.CreateMethodError = nil
	r.CreateMethodReturn = nil
}

func (r *CreatePackagingMockRepo) WithCtx(_ context.Context) service.ICreatePackagingRepo {
	return r
}

func (r *CreatePackagingMockRepo) Create(_ *service.CreatePackagingDto) (*models.Packaging, error) {
	r.CreateMethodCalledTimes++
	if r.CreateMethodError != nil {
		return nil, r.CreateMethodError
	}
	if r.CreateMethodReturn != nil {
		return r.CreateMethodReturn, nil
	}
	return mocks.GetPackagingModelMock(), nil
}
