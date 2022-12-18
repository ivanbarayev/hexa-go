package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	"time"
)

var (
	db  *mongo.Client
	Ctx context.Context
	err error
)

// NewMongoDB Return new MongoDB client
func NewMongoDB(cfg *config.Config) *mongo.Client {
	println("Driver MongoDB Initialized")
	ConnStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&authSource=admin", cfg.MongoDB.USER, cfg.MongoDB.PASS, cfg.MongoDB.HOST, cfg.MongoDB.DEFAULT_DB)
	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	db, err = mongo.Connect(Ctx, options.Client().ApplyURI(ConnStr))
	if err != nil {
		println(err.Error())
	} else {
		print("conn ok")
	}
	return db
}
