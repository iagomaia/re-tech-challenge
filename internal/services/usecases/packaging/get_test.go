package services_test

import (
	"context"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	service "github.com/iagomaia/re-tech-challenge/internal/services/usecases/packaging"
	repositoriesMocks "github.com/iagomaia/re-tech-challenge/mocks/infra/repositories/packaging"
)

type GetPackagingSutTypes struct {
	UseCase  *service.GetPackaging
	MockRepo *repositoriesMocks.GetPackagingMockRepo
}

func GetGetPackagingSutTypes() *GetPackagingSutTypes {
	mockRepo := &repositoriesMocks.GetPackagingMockRepo{}
	mockRepo.Init()

	useCase := &service.GetPackaging{
		GetRepo: mockRepo,
	}

	return &GetPackagingSutTypes{
		UseCase:  useCase,
		MockRepo: mockRepo,
	}
}

func Test_GetAll(t *testing.T) {
	t.Run("should get all packagings calling the repository", func(t *testing.T) {
		sut := GetGetPackagingSutTypes()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).GetAll()

		assert.Nil(t, err, "should not return error")
		assert.Equal(t, 1, sut.MockRepo.GetAllMethodCalledTimes, "get packaging repository not called")
	})
	t.Run("should return error if failed to get packagings", func(t *testing.T) {
		sut := GetGetPackagingSutTypes()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.MockRepo.GetAllMethodError = cErr

		packaging, err := sut.UseCase.WithCtx(
			context.Background(),
		).GetAll()

		assert.NotNil(t, err, "should return error")
		assert.Nil(t, packaging, "should return nil slice of packagings")
		assert.Equal(t, 1, sut.MockRepo.GetAllMethodCalledTimes, "get packaging repository not called")
	})
}
