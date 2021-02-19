package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saneduck/gin_rummy/models"
)

func Post(c *gin.Context) {
	u := models.NewUser()
	err := c.BindJSON(u)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})

	fmt.Println(u)
}
