package services_test

import (
	"context"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
	domainDtosMocks "github.com/iagomaia/re-tech-challenge/mocks/domain/dtos/packaging"
	repositoriesMocks "github.com/iagomaia/re-tech-challenge/mocks/infra/repositories/packaging"
)

type CreatePackagingSutTypes struct {
	UseCase  *service.CreatePackaging
	MockRepo *repositoriesMocks.CreatePackagingMockRepo
}

func GetCreatePackagingSutDependencies() *CreatePackagingSutTypes {
	mockRepo := &repositoriesMocks.CreatePackagingMockRepo{}
	mockRepo.Init()

	useCase := &service.CreatePackaging{
		CreateRepo: mockRepo,
	}

	return &CreatePackagingSutTypes{
		UseCase:  useCase,
		MockRepo: mockRepo,
	}
}

func Test_Create(t *testing.T) {
	t.Run("should create a packaging calling the repository", func(t *testing.T) {
		sut := GetCreatePackagingSutDependencies()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).Create(domainDtosMocks.GetCreatePackagingDomainDtoMock())

		assert.Nil(t, err, "should not return error")
		assert.Equal(t, 1, sut.MockRepo.CreateMethodCalledTimes, "create packaging repository not called")
	})
	t.Run("should return error if failed to create packaging", func(t *testing.T) {
		sut := GetCreatePackagingSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.MockRepo.CreateMethodError = cErr

		packaging, err := sut.UseCase.WithCtx(
			context.Background(),
		).Create(domainDtosMocks.GetCreatePackagingDomainDtoMock())

		assert.NotNil(t, err, "should return error")
		assert.Nil(t, packaging, "should return nil packaging")
		assert.Equal(t, 1, sut.MockRepo.CreateMethodCalledTimes, "create packaging repository not called")
	})
}
