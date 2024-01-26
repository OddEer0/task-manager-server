package models

type Tag struct {
	Id        string `json:"id" validate:"required,uuidv4"`
	Name      string `json:"name" validate:"required,min=3,max=30"`
	Color     string `json:"color" validate:"required"`
	Bg        string `json:"bg" validate:"required"`
	ProjectId string `json:"projectId" validate:"required,uuidv4"`
}
