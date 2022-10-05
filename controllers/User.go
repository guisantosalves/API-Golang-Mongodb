package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Createuser(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	req, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	insertedId := req.InsertedID

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"insertedId": insertedId,
		},
	}

	return res, nil
}

func Getusers(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// it will be run after all the other codes in this function
	defer cur.Close(ctx)

	// all the documents
	var users []bson.M

	for cur.Next(ctx) {

		var user bson.M

		if err = cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	res := map[string]interface{}{}

	res = map[string]interface{}{
		"data": users,
	}

	return res, nil
}

func UpdateOne(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	filtering := bson.M{"nome": data["nome"]}
	fields := bson.M{"$set": data}

	result, err := collection.UpdateOne(ctx, filtering, fields)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": result,
	}

	return res, nil

}

func Deleteuser(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	//to delete with the ID we need get the id from r.body and convert to string
	// not a map

	//{"user_id": "foksadofgksadfksdafkosakdf"}
	result, err := collection.DeleteOne(ctx, bson.M{"nome": data["nome"]})
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": result.DeletedCount,
	}

	return res, nil
}
