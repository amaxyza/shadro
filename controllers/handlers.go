package controllers

import (
	"net/http"

	"github.com/amaxyza/shadro/models"
	"github.com/amaxyza/shadro/services"
	"github.com/gin-gonic/gin"
)

func PingPongGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostLoginHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	_, err := services.ValidateUser(name, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "user does not exist",
		})
	}

	c.Status(200)
}

func PostCreateUserHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	user, err := services.AddUser(name, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to hash password.",
		})
	}

	c.JSON(http.StatusCreated, models.Publicize(user))
}
