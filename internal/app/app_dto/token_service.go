package appDto

type (
	GenerateTokenServiceDto struct {
		Id   string
		Role string
	}

	SaveTokenServiceDto struct {
		Id           string
		RefreshToken string
	}
)
