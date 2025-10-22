package mocks

import (
	"time"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
)

func GetPackagingModelMock() *models.Packaging {
	return &models.Packaging{
		Id:        "some-id",
		Size:      1,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}
