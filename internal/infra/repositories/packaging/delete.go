package repositories

import (
	"context"
	"errors"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/clients"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var (
	_ service.IDeletePackagingRepo = (*DeletePackagingRepo)(nil)
)

type DeletePackagingRepo struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *DeletePackagingRepo) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(PackagingCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *DeletePackagingRepo) WithCtx(ctx context.Context) service.IDeletePackagingRepo {
	return &DeletePackagingRepo{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *DeletePackagingRepo) DeleteById(id string) error {
	defer r.session.EndSession(r.ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Invalid ID format",
			OriginalError: err,
		}
		return cErr
	}
	filter := bson.D{
		{Key: "_id", Value: objId},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deletedAt", Value: time.Now()},
		}},
	}

	result := r.collection.FindOneAndUpdate(r.ctx, filter, update)
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		cErr := utils.CustomError{
			Status:        http.StatusNotFound,
			Message:       "Packaging not found",
			OriginalError: result.Err(),
		}
		return cErr
	}

	if result.Err() != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error deleting packaging",
			OriginalError: result.Err(),
		}
		return cErr
	}

	return nil
}
