package models

import (
	"time"
)

type Shader interface {
	GetID() int
	GetOwnerID() int
	GetName() string
	GetGLSL() string
	GetTimeCreated() time.Time
	GetTimeUpdated() time.Time

	SetName(new_name string)
	SetProgram(new_src string) error
	SetOwner(new_owner *User) error
}

type shader struct {
	id          int
	owner_id    int
	name        string
	glsl        string
	createdTime time.Time
	updatedTime time.Time
}

func NewShader(id, owner_id int, name, source string, createdTime, updatedTime time.Time) Shader {
	return &shader{
		id:          id,
		owner_id:    owner_id,
		name:        name,
		glsl:        source,
		createdTime: createdTime,
		updatedTime: updatedTime,
	}
}

func (s *shader) GetID() int {
	return s.id
}

func (s *shader) GetOwnerID() int {
	return s.owner_id
}

func (s *shader) GetName() string {
	return s.name
}

func (s *shader) GetGLSL() string {
	return s.glsl
}

func (s *shader) GetTimeCreated() time.Time {
	return s.createdTime
}

func (s *shader) GetTimeUpdated() time.Time {
	return s.updatedTime
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

func (s *shader) SetOwner(new_owner *User) error {
	s.owner_id = new_owner.ID
	s.updateTime()
	return nil
}

func (s *shader) updateTime() {
	s.updatedTime = time.Now()
}
