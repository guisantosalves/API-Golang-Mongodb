package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/guisantosalves/mongodb-go/controllers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// like the model
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"FirstName,omitempty" bson:"FirstName,omitempty"`
	Email string             `json:"LastName,omitempty" bson:"LastName,omitempty"`
}

var client *mongo.Client

func main() {

	// w -> response, r -> resquest
	handledefault := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "eaeae")
	}

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		// force err == nil
		log.Fatal("Error loading .env file")
	}

	// listing the databases
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil { //força a ser nulo
	// 	log.Fatal(err)
	// }

	// fmt.Println(databases)

	http.HandleFunc("/", handledefault)
	http.HandleFunc("/api/v1/user", requestHandler)
	http.ListenAndServe(":8080", nil)
}

// w -> response, r  -> request
func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	// making the key -> string and value -> any
	response := map[string]interface{}{}

	ctx := context.Background()

	apiString := os.Getenv("API_KEY")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(apiString))
	if err != nil {
		log.Println(err.Error())
	}

	collection := client.Database("store").Collection("custumer")

	data := map[string]interface{}{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)
	// := is for declaration + assignment, whereas = is for assignment only
	switch r.Method {
	case "POST":
		response, err = controllers.Createuser(collection, ctx, data)
		// case "GET":
		// 	response, err = getRecords(collection, ctx, data)
		// case "PUT":
		// 	response, err = updateRecord(collection, ctx, data)
		// case "DELETE":
		// 	response, err = deleteRecord(collection, ctx, data)
	}
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
	// qualquer instrução precedida pela palavra-chave defer não é invocada até o final da função na qual a defer tiver sido usada
	defer client.Disconnect(ctx)
}
