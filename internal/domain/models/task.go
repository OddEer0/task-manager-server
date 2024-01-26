package models

type Task struct {
	Id          string  `json:"id" validate:"require,uuidv4"`
	Name        string  `json:"name" validate:"require,min=3,max=127"`
	Description *string `json:"description" validate:"omitempty,min=3,max=255"`
	Priority    *string `json:"priority" validate:"omitempty,isPriority"`
	Order       int     `json:"order" validate:"required"`
	ColumnId    string  `json:"columnId" validate:"required,uuidv4"`
}
