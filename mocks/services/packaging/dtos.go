package packaging

import (
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

func GetCreatePackagingDataDtoMock() *service.CreatePackagingDto {
	return &service.CreatePackagingDto{
		Size: 1,
	}
}
