package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchBasicProfile(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "")
}

func FetchUsageQuota(c *gin.Context) {

}

func UpgradeToPremium(c *gin.Context) {

}
