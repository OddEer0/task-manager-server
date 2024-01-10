package dto

type GenerateTokenDto struct {
	Id    string
	Roles string
}

type SaveTokenDto struct {
	Id           string
	RefreshToken string
}
