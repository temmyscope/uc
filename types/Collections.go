package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
* The User Collection Model
**/
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	UserName   string             `json:"user_name" bson:"user_name"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	IsVerified bool               `json:"is_verified,omitempty" bson:"is_verified"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

/**
* The Device Collection Model
**/
type Device struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name       string             `json:"name" bson:"name"`
	Model      string             `json:"model" bson:"model"`
	UniqueID   string             `json:"unique_id" bson:"unique_id"`
	IsVerified bool               `json:"is_verified,omitempty" bson:"is_verified"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	LastSyncAt time.Time          `bson:"last_sync_at"` //last time the device was synced
}

/**
* The Clip Collection Model - represents each item copied to the clipBoard
**/
type Clip struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeviceId  primitive.ObjectID `json:"device_id" bson:"device_id"`
	Content   string             `json:"content" bson:"content"`
	IsSafe    bool               `json:"is_safe" bson:"is_safe"`                           //whether the text might be a password (length 8-16)
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`           //when it created on user device
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"` //when it was written to server
}
