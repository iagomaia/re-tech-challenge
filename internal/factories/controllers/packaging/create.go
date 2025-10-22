package factories

import (
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/services/packaging"
	controllers "github.com/iagomaia/re-tech-challenge/internal/presentation/controllers/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

var createPackagingController presentation.IHandler

func GetCreatePackagingController() presentation.IHandler {
	if createPackagingController != nil {
		return createPackagingController
	}

	createPackagingController = &controllers.CreatePackagingController{UseCase: factories.GetCreatePackagingUseCase()}
	return createPackagingController
}
