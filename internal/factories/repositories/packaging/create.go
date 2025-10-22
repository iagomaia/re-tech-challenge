package factories

import (
	packagingrepo "github.com/iagomaia/re-tech-challenge/internal/infra/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var createPackagingRepository service.ICreatePackagingRepo

func GetCreatePackagingRepository() service.ICreatePackagingRepo {
	if createPackagingRepository == nil {
		createPackagingRepository = new(packagingrepo.CreatePackagingRepo)
		createPackagingRepository.Init()
	}
	return createPackagingRepository
}
