package repositories

import (
	"time"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

func mapCreatePackagingDtoToDbe(dto *service.CreatePackagingDto) *PackagingDbe {
	return &PackagingDbe{
		Id:        nil,
		Size:      dto.Size,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func mapDbeToModel(dbe *PackagingDbe) *models.Packaging {
	return &models.Packaging{
		Id:        dbe.Id.Hex(),
		Size:      dbe.Size,
		CreatedAt: dbe.CreatedAt,
		UpdatedAt: dbe.UpdatedAt,
		DeletedAt: dbe.DeletedAt,
	}
}
