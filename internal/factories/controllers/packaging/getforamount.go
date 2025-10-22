package factories

import (
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/services/packaging"
	controllers "github.com/iagomaia/re-tech-challenge/internal/presentation/controllers/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

var getForAmountController presentation.IHandler

func GetGetForAmountController() presentation.IHandler {
	if getForAmountController != nil {
		return getForAmountController
	}

	getForAmountController = &controllers.GetPackagingsForAmountController{UseCase: factories.GetGetForAmountPackagingUseCase()}
	return getForAmountController
}
