package controllers

import (
	"net/http"

	"github.com/amaxyza/shadro/backend/models"
	"github.com/amaxyza/shadro/backend/services"
	"github.com/gin-gonic/gin"
)

type requested_user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PingPongGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostLoginHandler(c *gin.Context) {
	var ru requested_user

	if err := c.BindJSON(&ru); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "unable to serialize form input",
		})
	}

	_, err := services.ValidateUser(ru.Name, ru.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "unable to validate user",
		})
	}

	c.Status(200)
}

func PostCreateUserHandler(c *gin.Context) {
	var requested_user requested_user

	if err := c.BindJSON(&requested_user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to serialize form input",
		})
	}
	user, err := services.AddUser(requested_user.Name, requested_user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to hash password.",
		})
	}

	c.JSON(http.StatusCreated, models.Publicize(user))
}
