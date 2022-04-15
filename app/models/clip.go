package models

import (
	"context"

	Config "../../config"
	Types "../../types"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOneClip(clip *Types.Clip) *mongo.InsertOneResult {

	clipCollection := Config.DBConnection.Collection("clip")

	res, err := clipCollection.InsertOne(context.TODO(), clip)

	if err != nil {
		panic(err)
	}

	return res
}

func FetchClips() {

}

func FetchClipById(id string) {

}
