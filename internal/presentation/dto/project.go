package dto

type (
	CreateProjectDto struct {
		Name  string `json:"name" validate:"required,min=3,max=50"`
		Color string `json:"color" validate:"required"`
		Bg    string `json:"bg" validate:"required"`
	}
)
