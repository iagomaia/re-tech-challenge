package controllers

import (
	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	domain "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	contracts "github.com/iagomaia/re-tech-challenge/internal/presentation/contracts/packaging"
)

func mapPackagingModelToObject(model *models.Packaging) *contracts.Packaging {
	return &contracts.Packaging{
		Id:   model.Id,
		Size: model.Size,
	}
}

func mapPackagingModelListToObjectList(models []*models.Packaging) []*contracts.Packaging {
	var packagings = make([]*contracts.Packaging, len(models))
	for i, model := range models {
		packagings[i] = mapPackagingModelToObject(model)
	}
	return packagings
}

func mapDomainPacksResponseToContractResponse(packs *domain.GetForAmountResponse) *contracts.GetForAmountResponse {
	packsResponse := make([]*contracts.Packs, len(packs.Packs))
	for i, p := range packs.Packs {
		packsResponse[i] = &contracts.Packs{
			Size:     p.Size,
			Quantity: p.Quantity,
		}
	}
	return &contracts.GetForAmountResponse{
		Packs:        packsResponse,
		PackQuantity: packs.PackQuantity,
		TotalAmount:  packs.TotalAmount,
		LeftAmount:   packs.LeftAmount,
	}
}
