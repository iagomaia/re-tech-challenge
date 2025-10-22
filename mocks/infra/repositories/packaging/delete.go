package repositoriesmocks

import (
	"context"

	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var (
	_ service.IDeletePackagingRepo = (*DeletePackagingMockRepo)(nil)
)

type DeletePackagingMockRepo struct {
	DeleteByIdMethodCalledTimes int
	DeleteByIdMethodError       error
}

func (r *DeletePackagingMockRepo) Init() {
	r.DeleteByIdMethodCalledTimes = 0
	r.DeleteByIdMethodError = nil
}

func (r *DeletePackagingMockRepo) WithCtx(_ context.Context) service.IDeletePackagingRepo {
	return r
}

func (r *DeletePackagingMockRepo) DeleteById(_ string) error {
	r.DeleteByIdMethodCalledTimes++
	return r.DeleteByIdMethodError
}
