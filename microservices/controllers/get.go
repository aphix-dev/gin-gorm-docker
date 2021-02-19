package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	r := rand.Intn(100)
	c.JSON(http.StatusOK, gin.H{
		"num": r,
	})
}

func GetSpecific(c *gin.Context) {
	// pull the request out of the uri
	spec := c.Param("spec")
	c.JSON(http.StatusOK, gin.H{
		"num": spec,
	})
}
