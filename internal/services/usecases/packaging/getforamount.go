package services

import (
	"context"
	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	"net/http"
)

var (
	_ usecases.IGetPacks = (*GetPacksForAmount)(nil)
)

type GetPacksForAmount struct {
	GetRepo service.IGetPackagingRepo
	ctx     context.Context
}

func (c *GetPacksForAmount) WithCtx(ctx context.Context) usecases.IGetPacks {
	return &GetPacksForAmount{
		GetRepo: c.GetRepo,
		ctx:     ctx,
	}
}

func (c *GetPacksForAmount) GetForAmount(amount int) (response *usecases.GetForAmountResponse, err error) {
	availablePacks, err := c.GetRepo.WithCtx(c.ctx).GetAllSortedBySize()
	if err != nil {
		return
	}

	if len(availablePacks) == 0 {
		err = utils.CustomError{
			Status:        http.StatusBadRequest,
			Message:       "No packagings available",
			OriginalError: nil,
		}
		return
	}

	usablePacks := getUsablePacks(availablePacks, amount)

	if usablePacks[0].Size == amount {
		response = &usecases.GetForAmountResponse{
			Packs: []*usecases.Packs{
				{
					Size:     usablePacks[0].Size,
					Quantity: 1,
				},
			},
			PackQuantity: 1,
			TotalAmount:  amount,
			LeftAmount:   0,
		}
		return
	}

	packs := getPacks(amount, usablePacks, []*usecases.Packs{})
	totalQty := calculateTotalQty(packs)
	totalAmount := calculateTotalAmount(packs)
	response = &usecases.GetForAmountResponse{
		Packs:        packs,
		PackQuantity: totalQty,
		TotalAmount:  totalAmount,
		LeftAmount:   totalAmount - amount,
	}
	return
}

func getUsablePacks(availablePacks []*models.Packaging, amount int) []*models.Packaging {
	var usablePacks []*models.Packaging
	for i, _ := range availablePacks {
		// if we find an exact match
		if availablePacks[i].Size == amount {
			usablePacks = append(usablePacks, availablePacks[i])
			break
		}

		if availablePacks[i].Size < amount {
			if i == 0 {
				// all packagings are usable in the order
				usablePacks = availablePacks
				break
			} else {
				// we get the first pack that's bigger than the amount and all that are smaller
				// this way we can validate if it's worth to send a single pack bigger than the amount
				// or a combination of the smaller ones
				usablePacks = availablePacks[i-1:]
				break
			}
		}

		// at last, if we are at the last smaller packaging, and it's still bigger than the order amount,
		// we'll save to use just this one
		if i == len(availablePacks)-1 && availablePacks[i].Size > amount {
			usablePacks = append(usablePacks, availablePacks[i])
			break
		}
	}
	return usablePacks
}

func getPacks(amount int, usablePacks []*models.Packaging, currentPacks []*usecases.Packs) []*usecases.Packs {
	if len(usablePacks) == 1 {
		qty := 1
		if amount > usablePacks[0].Size {
			if amount%usablePacks[0].Size != 0 {
				qty = amount/usablePacks[0].Size + 1
			} else {
				qty = amount / usablePacks[0].Size
			}
		}
		pack := &usecases.Packs{
			Size:     usablePacks[0].Size,
			Quantity: qty,
		}
		newPacks := append(currentPacks, pack)
		return newPacks
	}

	if usablePacks[0].Size == amount {
		pack := &usecases.Packs{
			Size:     usablePacks[0].Size,
			Quantity: 1,
		}
		return append(currentPacks, pack)
	}

	if usablePacks[0].Size > amount {
		firstOptionPack := &usecases.Packs{
			Size:     usablePacks[0].Size,
			Quantity: 1,
		}
		firstOption := append(currentPacks, firstOptionPack)

		secondOption := getPacks(amount, usablePacks[1:], currentPacks)

		firstOptionTotalAmount := calculateTotalAmount(firstOption)
		secondOptionTotalAmount := calculateTotalAmount(secondOption)

		if secondOptionTotalAmount < firstOptionTotalAmount {
			return secondOption
		}

		// if the first option total is less than the second option, we should use it
		// if they are equal, the first option will use fewer packages
		return firstOption
	}

	remaining := amount % usablePacks[0].Size
	if remaining == 0 {
		pack := &usecases.Packs{
			Size:     usablePacks[0].Size,
			Quantity: amount / usablePacks[0].Size,
		}
		return append(currentPacks, pack)
	}

	firstOptionQty := amount / usablePacks[0].Size
	firstOptionPack := &usecases.Packs{
		Size:     usablePacks[0].Size,
		Quantity: firstOptionQty,
	}
	newPacks := append(currentPacks, firstOptionPack)
	firstOption := getPacks(remaining, usablePacks[1:], newPacks)

	secondOptionPack := &usecases.Packs{
		Size:     usablePacks[0].Size,
		Quantity: firstOptionQty + 1,
	}
	secondOption := append(currentPacks, secondOptionPack)

	firstOptionTotalAmount := calculateTotalAmount(firstOption)
	secondOptionTotalAmount := calculateTotalAmount(secondOption)

	if secondOptionTotalAmount < firstOptionTotalAmount {
		return secondOption
	}
	if secondOptionTotalAmount > firstOptionTotalAmount {
		return firstOption
	}

	// we'll only use the total quantity of packages if both options
	// result in the same number of items, in which case
	// we'll favor the option with fewer packages

	firstOptionTotalQty := calculateTotalQty(firstOption)
	secondOptionTotalQty := calculateTotalQty(secondOption)

	if firstOptionTotalQty <= secondOptionTotalQty {
		return firstOption
	}

	return secondOption
}

func calculateTotalAmount(packs []*usecases.Packs) int {
	var totalAmount int
	for _, pack := range packs {
		totalAmount += pack.Size * pack.Quantity
	}
	return totalAmount
}

func calculateTotalQty(packs []*usecases.Packs) int {
	var totalQty int
	for _, pack := range packs {
		totalQty += pack.Quantity
	}
	return totalQty
}
