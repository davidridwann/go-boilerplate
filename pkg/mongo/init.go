package mongo

import (
	"context"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://davidridwan:davidridwan123@cluster0.dzpwrtb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Err().Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Err().Fatal(err)
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Err().Fatal(err)
	}
	log.Std().Infoln("Mongo DB Connected")
	return client
}

// DB Client instance
var DB *mongo.Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("wlb_test_log").Collection(collectionName)
	return collection
}
