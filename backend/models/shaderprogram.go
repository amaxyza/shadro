package models

import (
	"time"
)

type Shader interface {
	Name() string
	Date() time.Time
	Author_ID() int
	GLSL() string

	SetName(new_name string)
	SetProgram(new_src string) error
	SetAuthor(new_author *User) error
}

type shader struct {
	name        string    `json:"name"`
	createdTime time.Time `json:"created"`
	updatedTime time.Time `json:"updated"`
	author_id   int       `json:"author_id`
	glsl        string    `json:"glsl"`
}

func CreateShader(name string, author *User, program_src string) (Shader, error) {
	return &shader{
		name:        name,
		createdTime: time.Now(),
		author_id:   author.ID,
		glsl:        program_src,
	}, nil
}

func (s *shader) Name() string {
	s.updateTime()
	return s.name
}

func (s *shader) Date() time.Time {
	s.updateTime()
	return s.createdTime
}

func (s *shader) Author_ID() int {
	s.updateTime()
	return s.author_id
}

func (s *shader) GLSL() string {
	s.updateTime()
	return s.glsl
}

func (s *shader) SetName(new_name string) {
	s.name = new_name
	s.updateTime()
}

func (s *shader) SetProgram(new_src string) error {
	s.glsl = new_src
	s.updateTime()
	return nil
}

func (s *shader) SetAuthor(new_author *User) error {
	s.author_id = new_author.ID
	s.updateTime()
	return nil
}

func (s *shader) updateTime() {
	s.updatedTime = time.Now()
}
