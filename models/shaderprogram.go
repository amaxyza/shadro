package models

type Shader struct {
	Author User   `json:"author"`
	GLSL   string `json:"glsl"`
}
