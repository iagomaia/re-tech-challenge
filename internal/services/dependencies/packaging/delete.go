package services

import (
	"context"
)

type IDeletePackagingRepo interface {
	DeleteById(id string) error
	WithCtx(ctx context.Context) IDeletePackagingRepo
	Init()
}
