package controllers

import (
	"strconv"
	"time"

	"github.com/amaxyza/shadro/services"
	"github.com/gin-gonic/gin"
)

/*
api.GET("/programs/:id", controllers.GetProgramHandler)
api.DELETE("/programs/:id", controllers.DeleteProgramHandler)
api.POST("/programs", controllers.PostProgramHandler)
*/
type glslProgram struct {
	ID       int       `json:"id"`
	Owner_id int       `json:"owner_id"`
	Name     string    `json:"name"`
	Glsl     string    `json:"glsl"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func GetProgramHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "unable to convert id to integer"},
		)
		return
	}

	program, err := services.GetProgram(id)

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "couldnt find program with request id"},
		)
		return
	}

	c.JSON(200, glslProgram{
		ID:       program.GetID(),
		Owner_id: program.GetOwnerID(),
		Name:     program.GetName(),
		Glsl:     program.GetGLSL(),
		Created:  program.GetTimeCreated(),
		Updated:  program.GetTimeUpdated(),
	})
}

func DeleteProgramHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "unable to convert id to integer"},
		)
		return
	}

	err = services.DeleteProgram(id)

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "program doesnt exist."},
		)
		return
	}
}

func PostProgramHandler(c *gin.Context) {
	c.Status(200)
	// var gp glslProgram
	// if err := c.ShouldBind(&gp); err != nil {
	// 	c.AbortWithStatusJSON(
	// 		400, gin.H{"error": "program doesnt exist."},
	// 	)
	// 	return
	// }

	// services.CreateProgram(models.User {
	// 	ID: 3,
	// 	Name: "amax",
	// 	Password_Hash: "xd",
	// }, gp.Name, )
}
