package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//save single clip
func SaveClip(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "")
}

//save and sync multiple clips upwards from client
func SyncClipsFromClient(c *gin.Context) {

}

//download and sync clips downwards/clientside from server
func SyncClipsFromServer(c *gin.Context) {

}

//fetch clips
func FetchClips(c *gin.Context) {

}

//fetch sensitive clip
func FetchSpecialClips(c *gin.Context) {

}

//delete clip
func DeleteClip(c *gin.Context) {

}
