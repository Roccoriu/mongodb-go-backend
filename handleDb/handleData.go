package handleDb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func SaveUser(db *mongo.Database, ctx context.Context, data bson.D) {
	userCollection := db.Collection("users")
	var err error
	var userResult *mongo.InsertOneResult

	if data == nil {
		fmt.Println("No or invalid data passed")
		return
	}

	if userResult, err = userCollection.InsertOne(ctx, data); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Added Record wih ID: ")
	fmt.Println(userResult.InsertedID)
}

func FetchUsers(db *mongo.Database, ctx context.Context) []bson.M {
	userCollection := db.Collection("users")
	var users []bson.M
	var cursor *mongo.Cursor
	var err error

	if cursor, err = userCollection.Find(ctx, bson.M{}); err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user bson.M
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

/*func FetchUser(db *mongo.Database, ctx context.Context) bson.M {

}*/

func SaveOrder(db *mongo.Database, ctx context.Context) {
	//orderCollection := goBackend.Collection("orders")

}
