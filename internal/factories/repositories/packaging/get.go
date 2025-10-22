package factories

import (
	packagingrepo "github.com/iagomaia/re-tech-challenge/internal/infra/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var getPackagingRepository service.IGetPackagingRepo

func GetGetPackagingRepository() service.IGetPackagingRepo {
	if getPackagingRepository == nil {
		getPackagingRepository = new(packagingrepo.GetPackagingRepo)
		getPackagingRepository.Init()
	}
	return getPackagingRepository
}
