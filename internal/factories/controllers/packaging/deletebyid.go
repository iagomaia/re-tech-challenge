package factories

import (
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/services/packaging"
	controllers "github.com/iagomaia/re-tech-challenge/internal/presentation/controllers/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

var deletePackagingByIdController presentation.IHandler

func GetDeletePackagingByIdController() presentation.IHandler {
	if deletePackagingByIdController != nil {
		return deletePackagingByIdController
	}

	deletePackagingByIdController = &controllers.DeletePackagingByIdController{UseCase: factories.GetDeletePackagingUseCase()}
	return deletePackagingByIdController
}
