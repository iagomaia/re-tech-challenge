package factories

import "github.com/iagomaia/re-tech-challenge/internal/infra/repositories"

func GetMongoClient() *repositories.MongoClient {
	return new(repositories.MongoClient)
}
