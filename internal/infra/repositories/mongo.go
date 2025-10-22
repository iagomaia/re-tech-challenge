package repositories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName = "re-tech-challenge"
)

type MongoClient struct {
	client *mongo.Client
}

func (m *MongoClient) initClient() {
	conString := os.Getenv("DB_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conString))
	if err != nil {
		log.Fatalf("unable to connect to mongo, error: %v\n", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("unable to reach mongo, error: %v\n", err)
	}
	m.client = client
}

func (m *MongoClient) GetCollection(collectionName string) (mongo.Session, *mongo.Collection, error) {
	if m.client == nil {
		m.initClient()
	}
	session, err := m.client.StartSession()
	if err != nil {
		return nil, nil, err
	}
	return session, session.Client().Database(DatabaseName).Collection(collectionName), nil
}
