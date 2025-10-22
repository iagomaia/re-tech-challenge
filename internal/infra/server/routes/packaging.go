package routes

import (
	"github.com/go-chi/chi/v5"
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/controllers/packaging"
	"github.com/iagomaia/re-tech-challenge/internal/infra/adapters"
	contractsutils "github.com/iagomaia/re-tech-challenge/internal/presentation/contracts"
	"github.com/iagomaia/re-tech-challenge/internal/presentation/contracts/packaging"
)

func GetPackagingRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", adapters.AdaptRoute[contracts.CreatePackagingRequest](factories.GetCreatePackagingController(), nil))
	r.Get("/", adapters.AdaptRoute[contractsutils.EmptyRequest](factories.GetGetAllPackagingController(), nil))
	r.Get("/amount/{amount}", adapters.AdaptRoute[contractsutils.EmptyRequest](factories.GetGetForAmountController(), &[]string{"amount"}))
	r.Delete("/{id}", adapters.AdaptRoute[contractsutils.EmptyRequest](factories.GetDeletePackagingByIdController(), &[]string{"id"}))
	return r
}
