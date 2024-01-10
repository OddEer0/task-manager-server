package valuesobject

import (
	"encoding/json"
	"errors"
	"regexp"
)

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, e.Value)
	if !match {
		return errors.New("incorrect email")
	}
	return nil
}

func (e *Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Value)
}

func NewEmail(email string) (Email, error) {
	result := Email{email}
	if err := result.Validate(); err != nil {
		return Email{}, err
	}
	return Email{email}, nil
}
