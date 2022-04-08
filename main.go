package main

import (
	"os"

	Controllers "./app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
