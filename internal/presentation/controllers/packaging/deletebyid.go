package controllers

import (
	"net/http"

	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	presentation "github.com/iagomaia/re-tech-challenge/internal/presentation/protocols"
)

type DeletePackagingByIdController struct {
	UseCase usecases.IDeletePackaging
}

func (c *DeletePackagingByIdController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	id := req.Params["id"]
	err := c.UseCase.WithCtx(req.Ctx).DeleteById(id)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusNoContent,
		Body:   struct{}{},
	}
	return resp, nil
}
