package dto

type LoginInputDto struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

type RegistrationInputDto struct {
	Nick      string `json:"nick" validate:"required,min=3,max=50"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,min=3"`
	LastName  string `json:"lastName" validate:"required,min=3"`
}
