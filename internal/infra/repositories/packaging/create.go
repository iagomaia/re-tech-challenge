package repositories

import (
	"context"
	"log"
	"net/http"

	models "github.com/iagomaia/re-tech-challenge/internal/domain/models/packaging"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
	factories "github.com/iagomaia/re-tech-challenge/internal/factories/clients"
	service "github.com/iagomaia/re-tech-challenge/internal/services/dependencies/packaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ service.ICreatePackagingRepo = (*CreatePackagingRepo)(nil)
)

type CreatePackagingRepo struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *CreatePackagingRepo) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(PackagingCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *CreatePackagingRepo) WithCtx(ctx context.Context) service.ICreatePackagingRepo {
	return &CreatePackagingRepo{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *CreatePackagingRepo) Create(dto *service.CreatePackagingDto) (*models.Packaging, error) {
	defer r.session.EndSession(r.ctx)
	dbe := mapCreatePackagingDtoToDbe(dto)
	result, err := r.collection.InsertOne(r.ctx, dbe)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error inserting packaging into DB",
			OriginalError: err,
		}
		return nil, cErr
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	dbe.Id = &id

	return mapDbeToModel(dbe), nil
}
