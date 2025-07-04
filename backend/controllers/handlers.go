package controllers

import (
	"net/http"
	"strconv"

	"github.com/amaxyza/shadro/models"
	"github.com/amaxyza/shadro/services"
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

	if err := c.ShouldBind(&ru); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "unable to serialize form input",
		})
		return
	}

	_, err := services.ValidateUser(ru.Name, ru.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "unable to validate user",
		})
		return
	}

	c.Status(200)
}

func PostCreateUserHandler(c *gin.Context) {
	var ru requested_user

	if err := c.ShouldBind(&ru); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to serialize form input",
		})
		return
	}

	user, err := services.AddUser(ru.Name, ru.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to hash password.",
		})
		return
	}

	c.JSON(http.StatusCreated, models.Publicize(user))
}

func GetUsersHandler(c *gin.Context) {
	c.JSON(200, services.GetUserList())
}

func GetUserWithID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "unable to convert id to integer",
		})
	}

	user, err := services.GetUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot find user with requested id",
		})
	}

	c.JSON(200, models.Publicize(user))
}
