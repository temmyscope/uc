package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//fetch devices
func FetchDevices(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "")
}

//fetch specific device
func FetchDeviceById(c *gin.Context) {

}

//delete device
func DeleteDeviceById(c *gin.Context) {

}
