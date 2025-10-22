package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PackagingCollection = "packaging"
)

type PackagingDbe struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	Size      int                 `bson:"size"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt *time.Time          `bson:"updatedAt"`
	DeletedAt *time.Time          `bson:"deletedAt"`
}
