package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
* The User Collection Model
**/
type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	name       string             `bson:"text"`
	email      string             `bson:"email"`
	password   string             `bson:"password"`
	isVerified bool               `bson:"is_verified"`
	createdAt  time.Time          `bson:"created_at"`
	updatedAt  time.Time          `bson:"updated_at"`
}

/**
* The Device Collection Model
**/
type Device struct {
	ID         primitive.ObjectID `bson:"_id"`
	userId     primitive.ObjectID `bson:"user_id"`
	name       string             `bson:"name"`
	model      string             `bson:"model"`
	uniqueID   string             `bson:"unique_id"`
	isVerified bool               `bson:"is_verified"`
	createdAt  time.Time          `bson:"created_at"`
	updatedAt  time.Time          `bson:"updated_at"`
	lastSyncAt time.Time          `bson:"last_sync_at"` //last time the device was synced
}

/**
* The Clip Collection Model - represents each item copid to the clipBoard
**/
type Clip struct {
	ID        primitive.ObjectID `bson:"_id"`
	userId    primitive.ObjectID `bson:"user_id"`
	deviceId  primitive.ObjectID `bson:"device_id"`
	content   string             `bson:"content"`
	isSafe    bool               `bson:"is_safe"`    //whether the text might be a password (length 8-16)
	createdAt time.Time          `bson:"created_at"` //when it created on user device
	updatedAt time.Time          `bson:"updated_at"` //when it was written to server
}
