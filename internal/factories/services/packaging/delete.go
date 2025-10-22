package factories

import (
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	repoFactories "github.com/iagomaia/re-tech-challenge/internal/factories/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
)

var deletePackagingUseCase usecases.IDeletePackaging

func GetDeletePackagingUseCase() usecases.IDeletePackaging {
	if deletePackagingUseCase != nil {
		return deletePackagingUseCase
	}

	deletePackagingUseCase = &service.DeletePackaging{
		DeleteRepo: repoFactories.GetDeletePackagingRepository(),
	}
	return deletePackagingUseCase
}
