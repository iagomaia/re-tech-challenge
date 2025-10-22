package usecases

import (
	"context"
)

type IDeletePackaging interface {
	DeleteById(id string) error
	WithCtx(ctx context.Context) IDeletePackaging
}
