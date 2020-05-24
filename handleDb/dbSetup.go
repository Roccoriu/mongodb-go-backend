package handleDb

import (
	"context"
	"fmt"
	"github.com/roccoriu/restapi/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func SetupClient() *mongo.Client {
	setting := settings.GetSettings()
	client, err := mongo.NewClient(options.Client().ApplyURI(setting.ConnectionString))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfuly connected to MongoDB cluster")
	return client
}

func SetupContext(client *mongo.Client) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return ctx
}

func GetCollection(client *mongo.Client, collection string) *mongo.Collection {
	setting := settings.GetSettings()
	fmt.Printf("Using db %s -> collection %s\n", setting.DBName, collection)

	return client.Database(setting.DBName).Collection(collection)
}
