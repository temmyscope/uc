package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	Controllers "./app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Clip struct {
	ID        string `json:"id"`
	UserId    string `json:"userId"`
	DeviceId  string `json:"device_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

var collection *mongo.Collection
var ctx = context.TODO()

type ClipBoard []Clip

var clips = []Clip{
	{ID: "1", UserId: "2", DeviceId: "3", Content: "The lyrics I choosse", CreatedAt: "2022-04-05 09:15"},
}

//route for fetching content to clipBoard
func fetchClip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ClipBoard Content:")
}

//route for adding content to clipBoard
func addClip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

// route for adding content to clipBoard
// (Run in background) for syncing offline clipboard with server
func addClips(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

// route for updating offline clipboard from server
func fetchClips(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/sync/clips", fetchClips).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func initDBConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("clipboarddb").Collection("tasks")
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

}
