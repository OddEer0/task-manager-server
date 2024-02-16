package models

type Project struct {
	Id          string  `json:"id" validate:"required,uuidv4"`
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description *string `json:"description" validate:"omitempty,min=3,max=255"`
	Bg          string  `json:"bg" validate:"required"`
	Color       string  `json:"color" validate:"required"`
	Order       int     `json:"order" validate:"required"`
	UserId      string  `json:"projectId" validate:"required,uuidv4"`
}
