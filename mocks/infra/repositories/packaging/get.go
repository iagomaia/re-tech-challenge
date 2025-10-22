package repositoriesmocks

import (
	"context"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	mocks "github.com/iagomaia/re-tech-challenge/mocks/domain/models"
)

var (
	_ service.IGetPackagingRepo = (*GetPackagingMockRepo)(nil)
)

type GetPackagingMockRepo struct {
	GetAllMethodCalledTimes             int
	GetAllSortedBySizeMethodCalledTimes int
	GetAllMethodReturn                  []*models.Packaging
	GetAllSortedBySizeReturn            []*models.Packaging
	GetAllMethodError                   error
	GetAllSortedBySizeError             error
}

func (r *GetPackagingMockRepo) Init() {
	r.GetAllMethodCalledTimes = 0
	r.GetAllSortedBySizeMethodCalledTimes = 0
	r.GetAllMethodReturn = []*models.Packaging{}
	r.GetAllSortedBySizeReturn = []*models.Packaging{}
	r.GetAllMethodError = nil
	r.GetAllSortedBySizeError = nil
}

func (r *GetPackagingMockRepo) WithCtx(_ context.Context) service.IGetPackagingRepo {
	return r
}

func (r *GetPackagingMockRepo) GetAll() ([]*models.Packaging, error) {
	r.GetAllMethodCalledTimes++
	if r.GetAllMethodError != nil {
		return nil, r.GetAllMethodError
	}
	if r.GetAllMethodReturn != nil {
		return r.GetAllMethodReturn, nil
	}
	return []*models.Packaging{mocks.GetPackagingModelMock()}, nil
}

func (r *GetPackagingMockRepo) GetAllSortedBySize() ([]*models.Packaging, error) {
	r.GetAllSortedBySizeMethodCalledTimes++
	if r.GetAllSortedBySizeError != nil {
		return nil, r.GetAllSortedBySizeError
	}
	if r.GetAllSortedBySizeReturn != nil {
		return r.GetAllSortedBySizeReturn, nil
	}
	return []*models.Packaging{mocks.GetPackagingModelMock()}, nil
}
