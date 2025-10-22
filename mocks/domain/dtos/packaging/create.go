package packaging

import (
	domain "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
)

func GetCreatePackagingDomainDtoMock() *domain.CreatePackagingDto {
	return &domain.CreatePackagingDto{
		Size: 1,
	}
}
