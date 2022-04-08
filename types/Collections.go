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
	Name       string             `bson:"text"`
	UserName   string             `bson:"text"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	IsVerified bool               `bson:"isVerified"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

/**
* The Device Collection Model
**/
type Device struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserId     primitive.ObjectID `bson:"userId"`
	Name       string             `bson:"name"`
	Model      string             `bson:"model"`
	UniqueID   string             `bson:"unique_id"`
	IsVerified bool               `bson:"is_verified"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	LastSyncAt time.Time          `bson:"last_sync_at"` //last time the device was synced
}

/**
* The Clip Collection Model - represents each item copied to the clipBoard
**/
type Clip struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"user_id"`
	DeviceId  primitive.ObjectID `bson:"device_id"`
	Content   string             `bson:"content"`
	IsSafe    bool               `bson:"is_safe"`    //whether the text might be a password (length 8-16)
	CreatedAt time.Time          `bson:"created_at"` //when it created on user device
	UpdatedAt time.Time          `bson:"updated_at"` //when it was written to server
}
