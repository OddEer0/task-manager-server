package models

type Column struct {
	Id        string `json:"id" validate:"require,uuidv4"`
	Name      string `json:"name" validate:"require,min=3,max=50"`
	Bg        string `json:"bg" validate:"require"`
	Color     string `json:"color" validate:"require"`
	Order     int    `json:"order" validate:"require"`
	ProjectId string `json:"projectId" validate:"require,uuidv4"`
}
