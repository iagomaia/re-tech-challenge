package usecases

import (
	"context"
)

type GetForAmountResponse struct {
	Packs        []*Packs
	PackQuantity int
	TotalAmount  int
	LeftAmount   int
}

type Packs struct {
	Size     int
	Quantity int
}

type IGetPacks interface {
	GetForAmount(amount int) (*GetForAmountResponse, error)
	WithCtx(ctx context.Context) IGetPacks
}
