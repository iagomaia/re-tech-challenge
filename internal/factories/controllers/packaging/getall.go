package factories

import (
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/services/packaging"
	controllers "github.com/iagomaia/re-tech-challenge/internal/presentation/controllers/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

var getAllPackagingController presentation.IHandler

func GetGetAllPackagingController() presentation.IHandler {
	if getAllPackagingController != nil {
		return getAllPackagingController
	}

	getAllPackagingController = &controllers.GetAllPackagingController{UseCase: factories.GetGetAllPackagingUseCase()}
	return getAllPackagingController
}
