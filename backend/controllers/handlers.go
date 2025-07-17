package controllers

import (
	"fmt"
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

func LogoutHandler(c *gin.Context) {
	//Overwrite current cookie with bad one
	c.SetCookie("token", "", -1, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func GetMeHandler(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "missing jwt or cookie",
		})
		return
	}

	username, id, err := services.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "invalid token",
		})
		return
	}

	fmt.Println("DEBUG: printing user with name - " + username)

	c.JSON(200, gin.H{"username": username, "id": id})
}

func PostLoginHandler(c *gin.Context) {
	var ru requested_user

	if err := c.ShouldBind(&ru); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "unable to serialize form input",
		})
		return
	}

	id, err := services.ValidateUser(ru.Name, ru.Password)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status": "unable to validate user",
		})
		return
	}

	// Create token and set cookie
	token_str, err := services.CreateToken(id, ru.Name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.SetCookie("token", token_str, 3600*24*30, "/", "", true, true)

	c.JSON(200, gin.H{"success": true, "id": id})
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
			"error": err.Error(),
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
