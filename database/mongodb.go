package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	databaseName   = ""
	collectionName = ""
	dbUri          = "mongodb+srv://tccAdmin:pagamentoDigitalTCC123@pgmntdigital.hzuww3o.mongodb.net/?retryWrites=true&w=majority" // USE TO CONNECT ON SERVER.GO
)

type MongoConnection struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	Database   *mongo.Database
}

func Connect(uri string) (MongoConnection, context.Context, context.CancelFunc, error) {
	mongoConnection := MongoConnection{}

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	mongoConnection.Client = client
	mongoConnection.Database = client.Database(databaseName)

	return mongoConnection, ctx, cancel, err
}

func Close(mongoConnection MongoConnection, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := mongoConnection.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Ping(mongoConnection MongoConnection, ctx context.Context) error {

	if err := mongoConnection.Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
