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

type DeletePackagingSutTypes struct {
	UseCase  *service.DeletePackaging
	MockRepo *repositoriesMocks.DeletePackagingMockRepo
}

func GetDeletePackagingSutDependencies() *DeletePackagingSutTypes {
	mockRepo := &repositoriesMocks.DeletePackagingMockRepo{}
	mockRepo.Init()

	useCase := &service.DeletePackaging{
		DeleteRepo: mockRepo,
	}

	return &DeletePackagingSutTypes{
		UseCase:  useCase,
		MockRepo: mockRepo,
	}
}

func Test_DeleteById(t *testing.T) {
	t.Run("should delete a packaging calling the repository", func(t *testing.T) {
		sut := GetDeletePackagingSutDependencies()
		err := sut.UseCase.WithCtx(
			context.Background(),
		).DeleteById("some-id")

		assert.Nil(t, err, "should not return error")
		assert.Equal(t, 1, sut.MockRepo.DeleteByIdMethodCalledTimes, "delete packaging repository not called")
	})
	t.Run("should return error if failed to delete packaging", func(t *testing.T) {
		sut := GetDeletePackagingSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.MockRepo.DeleteByIdMethodError = cErr

		err := sut.UseCase.WithCtx(
			context.Background(),
		).DeleteById("some-id")

		assert.NotNil(t, err, "should return error")
		assert.Equal(t, 1, sut.MockRepo.DeleteByIdMethodCalledTimes, "delete packaging repository not called")
	})
}
