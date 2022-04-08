package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//register
func Register(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "")
}

//login
func Login(c *gin.Context) {

}

//forgot password
func ForgotPassword(c *gin.Context) {

}

//reset password
func ResetPassword(c *gin.Context) {

}

//special device login [username & otp]
func DeviceLogin(c *gin.Context) {

}
