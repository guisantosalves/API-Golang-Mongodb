package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// like the model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"FirstName,omitempty" bson:"FirstName,omitempty"`
	LastName  string             `json:"LastName,omitempty" bson:"LastName,omitempty"`
}

var client *mongo.Client

func main() {

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		// force err == nil
		log.Fatal("Error loading .env file")
	}

	apiString := os.Getenv("API_KEY")

	// conectou, tive que gerar nova senha no atlas
	client, err := mongo.NewClient(options.Client().ApplyURI(apiString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// qualquer instrução precedida pela palavra-chave defer não é invocada até o final da função na qual a defer tiver sido usada
	defer client.Disconnect(ctx)

	// listing the databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil { //força a ser nulo
		log.Fatal(err)
	}

	fmt.Println(databases)

}
