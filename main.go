package main

import (
	"context"
	"fmt"
	"log"
	"os"

	Controllers "./app/controllers"
	Types "./types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func initDBConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("clipboarddb").Collection("tasks")
}

type ClipBoard []Types.Clip

var clips = []Types.Clip{
	{UserId: "2", DeviceId: "3", Content: "The lyrics I choosse", CreatedAt: "2022-04-05 09:15"},
}

func main() {
	godotenv.Load()

	router := gin.Default()

	authRoute := router.Group("/api/auth")

	//register
	authRoute.POST("/register", Controllers.Register)

	//login
	authRoute.POST("/login", Controllers.Login)

	//forgot password
	authRoute.POST("/forgot-password", Controllers.ForgotPassword)

	//reset password
	authRoute.POST("/reset-password", Controllers.ResetPassword)

	//device-login using otp
	authRoute.POST("/device-login", Controllers.DeviceLogin)

	//save single clip

	//save and sync multiple clips upwards from client

	//fetch clips

	//fetch sensitive clip

	//download and sync clips downwards/clientside from server

	//delete clip

	//fetch devices

	//fetch specific device

	//delete device

	//logout

	router.Run(os.Getenv("PORT"))

}
