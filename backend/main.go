package main

import (
	"fmt"

	"github.com/amaxyza/shadro/controllers"
	"github.com/amaxyza/shadro/services"
	"github.com/gin-gonic/gin"
)

func main() {
	err := services.Connect()
	defer services.ClosePool()

	if err != nil {
		fmt.Errorf("couldn't connect to postgresql renderer db")
		return
	}

	r := gin.Default()

	// static files served under /static/*
	r.Static("/static", "./static")

	// templates
	r.LoadHTMLGlob("templates/*")

	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsersHandler)
		api.GET("/users/:id", controllers.GetUserWithID)
		//api.GET("/programs")

		api.POST("/login", controllers.PostLoginHandler)
		api.POST("/signup", controllers.PostCreateUserHandler)
	}

	r.Run(":8081")
}
