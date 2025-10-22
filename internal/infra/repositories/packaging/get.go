package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/clients"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ service.IGetPackagingRepo = (*GetPackagingRepo)(nil)
)

type GetPackagingRepo struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *GetPackagingRepo) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(PackagingCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *GetPackagingRepo) WithCtx(ctx context.Context) service.IGetPackagingRepo {
	return &GetPackagingRepo{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *GetPackagingRepo) GetAll() ([]*models.Packaging, error) {
	defer r.session.EndSession(r.ctx)
	filter := bson.D{
		{Key: "deletedAt", Value: nil},
	}
	result, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error searching for packagings on the DB",
			OriginalError: err,
		}
		return nil, cErr
	}

	var packagings []*models.Packaging
	for result.Next(r.ctx) {
		var dbe PackagingDbe
		err := result.Decode(&dbe)
		if err != nil {
			cErr := utils.CustomError{
				Status:        http.StatusInternalServerError,
				Message:       "Error decoding DB result",
				OriginalError: err,
			}
			return nil, cErr
		}
		model := mapDbeToModel(&dbe)
		packagings = append(packagings, model)
	}
	err = result.Close(r.ctx)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error decoding packagins",
			OriginalError: err,
		}
		return nil, cErr
	}
	return packagings, nil
}

func (r *GetPackagingRepo) GetAllSortedBySize() ([]*models.Packaging, error) {
	defer r.session.EndSession(r.ctx)
	filter := bson.D{
		{Key: "deletedAt", Value: nil},
	}
	sort := &options.FindOptions{
		Sort: bson.D{
			{Key: "size", Value: -1},
		},
	}
	result, err := r.collection.Find(r.ctx, filter, sort)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error searching for packagings on the DB",
			OriginalError: err,
		}
		return nil, cErr
	}

	var packagings []*models.Packaging
	for result.Next(r.ctx) {
		var dbe PackagingDbe
		err := result.Decode(&dbe)
		if err != nil {
			cErr := utils.CustomError{
				Status:        http.StatusInternalServerError,
				Message:       "Error decoding DB result",
				OriginalError: err,
			}
			return nil, cErr
		}
		model := mapDbeToModel(&dbe)
		packagings = append(packagings, model)
	}
	err = result.Close(r.ctx)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error decoding packagins",
			OriginalError: err,
		}
		return nil, cErr
	}
	return packagings, nil
}
