package models

import (
	"context"
	"encoding/json"
	"fmt"

	Config "../../config"
	Types "../../types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConnection *mongo.Database = Config.ConnectDB()

func CreateOneClip(clip Types.Clip) primitive.ObjectID {
	collection := dbConnection.Collection("clip")

	res, err := collection.InsertOne(context.TODO(), bson.M(clip))

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", "Document Created", res.InsertedID)

	//id :=

	return res.InsertedID
}

func FetchClips() {

}

func FetchClipById(id string) {

}
