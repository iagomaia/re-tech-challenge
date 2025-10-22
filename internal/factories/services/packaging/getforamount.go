package factories

import (
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	repoFactories "github.com/iagomaia/re-tech-challenge/internal/factories/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
)

var getForAmountPackagingUseCase usecases.IGetPacks

func GetGetForAmountPackagingUseCase() usecases.IGetPacks {
	if getForAmountPackagingUseCase != nil {
		return getForAmountPackagingUseCase
	}

	getForAmountPackagingUseCase = &service.GetPacksForAmount{
		GetRepo: repoFactories.GetGetPackagingRepository(),
	}
	return getForAmountPackagingUseCase
}
