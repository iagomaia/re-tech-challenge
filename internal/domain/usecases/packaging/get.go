package usecases

import (
	"context"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
)

type IGetPackagings interface {
	GetAll() ([]*models.Packaging, error)
	WithCtx(ctx context.Context) IGetPackagings
}
