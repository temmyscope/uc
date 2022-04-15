package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Model "github.com/temmyscope/uc/app/models"
	Types "github.com/temmyscope/uc/types"
)

//save single clip
func SaveClip(c *gin.Context) {
	var clip Types.Clip

	if err := c.BindJSON(&clip); err != nil {
		return
	}

	data := Model.CreateOneClip(clip)

	c.IndentedJSON(http.StatusCreated, data)
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
