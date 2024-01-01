package dto

type CreateUserDto struct {
	Name           string `json:"name" validate:"required,min=3"`
	Password       string `json:"password" validate:"required,min=8"`
	Email          string `json:"email" validate:"required, min=6"`
	FirstName      string `json:"firstName" validate:"required, min=3"`
	LastName       string `json:"lastName" validate:"required, min=3"`
	ActivationLink string `json:"activationLink"`
	Role           string `json:"role"`
}
