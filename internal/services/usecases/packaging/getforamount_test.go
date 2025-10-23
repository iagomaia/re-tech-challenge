package services_test

import (
	"context"
	"errors"
	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/re-tech-challenge/internal/domain/usecases/packaging"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
	repositoriesMocks "github.com/iagomaia/re-tech-challenge/mocks/infra/repositories/packaging"
)

type GetPackagingForAmountSutTypes struct {
	UseCase  *service.GetPacksForAmount
	MockRepo *repositoriesMocks.GetPackagingMockRepo
}

func GetGetPackagingForAmountSutTypes() *GetPackagingForAmountSutTypes {
	mockRepo := &repositoriesMocks.GetPackagingMockRepo{}
	mockRepo.Init()

	useCase := &service.GetPacksForAmount{
		GetRepo: mockRepo,
	}

	return &GetPackagingForAmountSutTypes{
		UseCase:  useCase,
		MockRepo: mockRepo,
	}
}

var mockPacks = []*models.Packaging{
	{
		Id:   "5",
		Size: 5000,
	},
	{
		Id:   "4",
		Size: 2000,
	},
	{
		Id:   "3",
		Size: 1000,
	},
	{
		Id:   "2",
		Size: 500,
	},
	{
		Id:   "1",
		Size: 250,
	},
}

func Test_GetForAmount(t *testing.T) {
	tests := []struct {
		Name          string
		Amount        int
		ExpectedPacks *usecases.GetForAmountResponse
		ExpectedError error
		MockPacks     []*models.Packaging
		MockError     error
	}{
		{
			Name:   "should get correct number of packs for 1 item",
			Amount: 1,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     250,
						Quantity: 1,
					},
				},
				PackQuantity: 1,
				TotalAmount:  250,
				LeftAmount:   249,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 250",
			Amount: 250,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     250,
						Quantity: 1,
					},
				},
				PackQuantity: 1,
				TotalAmount:  250,
				LeftAmount:   0,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 251",
			Amount: 251,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     500,
						Quantity: 1,
					},
				},
				PackQuantity: 1,
				TotalAmount:  500,
				LeftAmount:   249,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 501",
			Amount: 501,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     500,
						Quantity: 1,
					},
					{
						Size:     250,
						Quantity: 1,
					},
				},
				PackQuantity: 2,
				TotalAmount:  750,
				LeftAmount:   249,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 750",
			Amount: 750,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     500,
						Quantity: 1,
					},
					{
						Size:     250,
						Quantity: 1,
					},
				},
				PackQuantity: 2,
				TotalAmount:  750,
				LeftAmount:   0,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 4500",
			Amount: 4500,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     2000,
						Quantity: 2,
					},
					{
						Size:     500,
						Quantity: 1,
					},
				},
				PackQuantity: 3,
				TotalAmount:  4500,
				LeftAmount:   0,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:   "should get correct number of packs for 12001",
			Amount: 12001,
			ExpectedPacks: &usecases.GetForAmountResponse{
				Packs: []*usecases.Packs{
					{
						Size:     5000,
						Quantity: 2,
					},
					{
						Size:     2000,
						Quantity: 1,
					},
					{
						Size:     250,
						Quantity: 1,
					},
				},
				PackQuantity: 4,
				TotalAmount:  12250,
				LeftAmount:   249,
			},
			ExpectedError: nil,
			MockPacks:     mockPacks,
		},
		{
			Name:          "should return error if no packs available",
			Amount:        12001,
			ExpectedPacks: nil,
			ExpectedError: utils.CustomError{
				Status:        http.StatusBadRequest,
				Message:       "No packagings available",
				OriginalError: nil,
			},
			MockPacks: []*models.Packaging{},
			MockError: nil,
		},
		{
			Name:          "should return error if repository returns error",
			Amount:        12001,
			ExpectedPacks: nil,
			ExpectedError: errors.New("some repo error"),
			MockPacks:     []*models.Packaging{},
			MockError:     errors.New("some repo error"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sut := GetGetPackagingForAmountSutTypes()
			sut.MockRepo.GetAllSortedBySizeReturn = test.MockPacks
			sut.MockRepo.GetAllSortedBySizeError = test.MockError
			response, err := sut.UseCase.WithCtx(
				context.Background(),
			).GetForAmount(test.Amount)
			if test.ExpectedError == nil {
				assert.Nil(t, err, "should not return error")
				assert.ElementsMatch(t, test.ExpectedPacks.Packs, response.Packs, "should return correct packs")
			} else {
				assert.NotNil(t, err, "should return error")
				assert.Equal(t, test.ExpectedError, err, "errors should match")
			}

		})
	}
}
