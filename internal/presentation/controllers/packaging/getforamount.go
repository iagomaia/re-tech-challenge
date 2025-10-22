package controllers

import (
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	"net/http"
	"strconv"

	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

type GetPackagingsForAmountController struct {
	UseCase usecases.IGetPacks
}

func (c *GetPackagingsForAmountController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	strAmount := req.Params["amount"]
	amount, err := strconv.Atoi(strAmount)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusBadRequest,
			Message:       "Invalid amount",
			OriginalError: err,
		}
		return nil, cErr
	}
	packs, err := c.UseCase.WithCtx(req.Ctx).GetForAmount(amount)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusOK,
		Body:   mapDomainPacksResponseToContractResponse(packs),
	}
	return resp, nil
}
