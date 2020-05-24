package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/roccoriu/restapi/handleDb"
	"github.com/roccoriu/restapi/helpers"
	"github.com/roccoriu/restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	// connect to the database
	client := handleDb.SetupClient()
	ctx := handleDb.SetupContext(client)
	collection := handleDb.GetCollection(client, "users")

	w.Header().Set("Content-Type", "application/json")

	var users []models.User

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user models.User

		if err := cur.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	// connect to the database
	client := handleDb.SetupClient()
	ctx := handleDb.SetupContext(client)
	collection := handleDb.GetCollection(client, "users")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	name, _ := params["name"]

	filter := bson.M{"name": name}

	var user models.User
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		helpers.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	// connect to the database
	client := handleDb.SetupClient()
	ctx := handleDb.SetupContext(client)
	collection := handleDb.GetCollection(client, "users")

	w.Header().Set("Content-Type", "application/json")

	var newUser models.User

	_ = json.NewDecoder(r.Body).Decode(&newUser)

	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		helpers.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	/*	// connect to the database
		client := handleDb.SetupClient()
		ctx := handleDb.SetupContext(client)
		collection := handleDb.GetCollection(client, "users")

		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		var user models.User*/
}

func delUser(w http.ResponseWriter, r *http.Request) {

}
