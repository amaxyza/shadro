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

	// // routes
	// r.GET("/ping", controllers.PingPongGet)

	// r.GET("/login", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// })

	// r.GET("/signup", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "signup.html", nil)
	// })

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })

	api := r.Group("/api")
	{
		api.POST("/login", controllers.PostLoginHandler)
		api.POST("/signup", controllers.PostCreateUserHandler)
	}

	r.Run(":8081")
}
