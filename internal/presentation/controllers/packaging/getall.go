package controllers

import (
	"net/http"

	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

type GetAllPackagingController struct {
	UseCase usecases.IGetPackagings
}

func (c *GetAllPackagingController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	packs, err := c.UseCase.WithCtx(req.Ctx).GetAll()
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusOK,
		Body:   mapPackagingModelListToObjectList(packs),
	}
	return resp, nil
}
