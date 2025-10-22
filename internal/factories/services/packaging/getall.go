package factories

import (
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	repoFactories "github.com/iagomaia/re-tech-challenge/internal/factories/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
)

var getAllPackagingUseCase usecases.IGetPackagings

func GetGetAllPackagingUseCase() usecases.IGetPackagings {
	if getAllPackagingUseCase != nil {
		return getAllPackagingUseCase
	}

	getAllPackagingUseCase = &service.GetPackaging{
		GetRepo: repoFactories.GetGetPackagingRepository(),
	}
	return getAllPackagingUseCase
}
