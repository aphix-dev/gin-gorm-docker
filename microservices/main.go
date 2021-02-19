package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saneduck/gin_rummy/controllers"
	"github.com/saneduck/gin_rummy/middleware"
	"github.com/saneduck/gin_rummy/models"
)

func main() {
	//r := gin.Default()
	r := gin.New()

	db := models.SetupModels()
	r.Use(gin.Recovery(), middleware.Logger())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", controllers.Get)
	r.GET("/:spec", controllers.GetSpecific)
	r.POST("/", controllers.Post)

	r.Run(":9090")
}
