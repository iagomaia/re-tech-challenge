package factories

import (
	packagingrepo "github.com/iagomaia/re-tech-challenge/internal/infra/repositories/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
)

var deletePackagingRepository service.IDeletePackagingRepo

func GetDeletePackagingRepository() service.IDeletePackagingRepo {
	if deletePackagingRepository == nil {
		deletePackagingRepository = new(packagingrepo.DeletePackagingRepo)
		deletePackagingRepository.Init()
	}
	return deletePackagingRepository
}
