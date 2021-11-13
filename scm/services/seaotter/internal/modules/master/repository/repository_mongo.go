package repository

import (
	"monorepo/services/seaotter/pkg/constant"

	"go.mongodb.org/mongo-driver/mongo"
)

type masterRepoMongo struct {
	readDB, writeDB *mongo.Database
	collection      string
}

type MasterRepoMongo interface{}

// NewMasterRepoMongo mongo repo constructor
func NewMasterRepoMongo(readDB, writeDB *mongo.Database) MasterRepoMongo {
	return &masterRepoMongo{
		readDB:     readDB,
		writeDB:    writeDB,
		collection: constant.TableMasterSOPrefix,
	}
}
