package controllers

import (
	contracts "github.com/iagomaia/re-tech-challenge/internal/presentation/contracts/packaging"
	"net/http"

	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

type CreatePackagingController struct {
	UseCase usecases.ICreatePackaging
}

func (c *CreatePackagingController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	reqBody := req.Body.(*contracts.CreatePackagingRequest)

	dto := &usecases.CreatePackagingDto{
		Size: reqBody.Size,
	}

	pack, err := c.UseCase.WithCtx(req.Ctx).Create(dto)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusCreated,
		Body:   mapPackagingModelToObject(pack),
	}
	return resp, nil
}
