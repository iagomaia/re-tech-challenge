package services

import (
	"context"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	"math"
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

	availablePacksInt := make([]int, len(availablePacks))
	for i, pack := range availablePacks {
		availablePacksInt[i] = pack.Size
	}
	bestTotal, bestPackaging := getBestCombo(amount, availablePacksInt)

	packs := make([]*usecases.Packs, 0, len(bestPackaging))
	for k, v := range bestPackaging {
		packs = append(packs, &usecases.Packs{
			Size:     k,
			Quantity: v,
		})
	}

	totalQty := calculateTotalQty(packs)
	response = &usecases.GetForAmountResponse{
		Packs:        packs,
		PackQuantity: totalQty,
		TotalAmount:  bestTotal,
		LeftAmount:   bestTotal - amount,
	}
	return
}

func calculateTotalQty(packs []*usecases.Packs) int {
	var totalQty int
	for _, pack := range packs {
		totalQty += pack.Quantity
	}
	return totalQty
}

// Use Dynamic Programming search to find the best solution
func getBestCombo(amount int, packSizes []int) (bestTotal int, bestCombo map[int]int) {
	if amount <= 0 {
		return 0, map[int]int{}
	}

	// Add the biggest package size to amount to get a max possible total
	maxTotal := amount + packSizes[0]

	// dp[total] = minimum packs needed to reach this total
	dp := make([]int, maxTotal+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0 // 0 items, 0 packs

	// Track the last pack used for backtracking
	prev := make([]int, maxTotal+1)

	// Fill DP table
	for _, p := range packSizes {
		for total := p; total <= maxTotal; total++ {
			if dp[total-p]+1 < dp[total] {
				dp[total] = dp[total-p] + 1
				prev[total] = p
			}
		}
	}

	// Find the smallest total >= amount that is achievable
	bestTotal = -1
	for total := amount; total <= maxTotal; total++ {
		if dp[total] != math.MaxInt {
			bestTotal = total
			break
		}
	}

	if bestTotal == -1 {
		return -1, nil // sanity check for no solution
	}

	// Reconstruct which packs were used
	bestCombo = make(map[int]int)
	t := bestTotal
	for t > 0 {
		p := prev[t]
		if p == 0 {
			break
		}
		bestCombo[p]++
		t -= p
	}

	return bestTotal, bestCombo
}
