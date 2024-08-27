package db

import (
	"context"
	"fmt"
	"time"

	"github.com/coderajay94/go-mongo-with-streams/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "users"
	collection = "userAccount"
)

type (
	MongoDatabase interface{
		SaveUserData(ctx context.Context, userData model.UserAccountRequest) error
	}

	mongoDatabase struct {
		client         *mongo.Client
		collection *mongo.Collection
	}
)

func (mdb mongoDatabase) NewClient(host, username, password string) (MongoDatabase, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s", username, password, host)

	clientOptions := options.Client().ApplyURI(connectionString).SetTimeout(5*time.Second)
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err !=nil {
		return nil, err
	}
	database := mongoClient.Database(database)
	return &mongoDatabase{mongoClient, database.Collection(collection)}, nil
}

func (mdb mongoDatabase) Close(ctx context.Context)error{
	return mdb.client.Disconnect(ctx)
}

func (mdb mongoDatabase)SaveUserData(ctx context.Context, userData model.UserAccountRequest) error{
	_, err := mdb.collection.InsertOne(context.TODO(), userData)
	if err != nil{
		fmt.Println("Error inserting user into mongo db:", err)
	}
	return err
}