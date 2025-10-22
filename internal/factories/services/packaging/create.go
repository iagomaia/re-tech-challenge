package factories

import (
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	repoFactories "github.com/iagomaia/re-tech-challenge/internal/factories/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
)

var createPackagingUseCase usecases.ICreatePackaging

func GetCreatePackagingUseCase() usecases.ICreatePackaging {
	if createPackagingUseCase != nil {
		return createPackagingUseCase
	}

	createPackagingUseCase = &service.CreatePackaging{
		CreateRepo: repoFactories.GetCreatePackagingRepository(),
	}
	return createPackagingUseCase
}
