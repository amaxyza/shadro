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
type programInput struct {
	Owner_ID     int    `json:"owner_id"`
	Program_Name string `json:"program_name`
	Source       string `json:"source"`
}

type glslProgram struct {
	ID       int       `json:"id"`
	Owner_id int       `json:"owner_id"`
	Name     string    `json:"name"`
	Glsl     string    `json:"glsl"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// TODO: Fix programs in response JSON being empty.
func GetAllUserProgramsHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "unable to convert id to integer"},
		)
		return
	}

	programs, err := services.GetAllUserPrograms(id)

	if err != nil {
		c.AbortWithStatusJSON(
			400, gin.H{"error": "unable to create program list of user with id " + string(id)},
		)
		return
	}

	c.JSON(200, programs)
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
	var programInput programInput

	if err := c.ShouldBind(&programInput); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": "unable to bind data to expected input"})
		return
	}

	shaderprogram, err := services.CreateProgram(
		programInput.Owner_ID,
		programInput.Program_Name,
		programInput.Source,
	)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": err.Error()})
		return
	}

	c.JSON(201, shaderprogram)
}
