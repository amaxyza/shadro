package main

import (
	"github.com/amaxyza/shadro/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// static files served under /static/*
	r.Static("/static", "./static")

	// templates
	r.LoadHTMLGlob("templates/*")

	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsersHandler)
		api.GET("/users/:id", controllers.GetUserWithID)

		api.POST("/login", controllers.PostLoginHandler)
		api.POST("/signup", controllers.PostCreateUserHandler)
	}

	r.Run(":8081")
}
