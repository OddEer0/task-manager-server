package valuesobject

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	BirthdayIsZero       = "birthday is zero"
	BirthdayIncorrectMax = "birthday max value after 12 year"
	BirthdayIncorrectMin = "birthday min value 1900 year"
)

var (
	MinBirthdayTime = time.Date(1900, 1, 10, 0, 0, 0, 0, time.UTC)
	MaxBirthdayTime = time.Now().AddDate(-12, 0, 0)
)

type Birthday struct {
	Value time.Time
}

func (b *Birthday) Validate() error {
	if b.Value.IsZero() {
		return errors.New(BirthdayIsZero)
	}
	if b.Value.Before(MinBirthdayTime) {
		return errors.New(BirthdayIncorrectMin)
	}
	if b.Value.After(MaxBirthdayTime) {
		return errors.New(BirthdayIncorrectMax)
	}
	return nil
}

func (b *Birthday) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Value.Format("2006-01-02T15:04:05Z07:00"))
}

func NewBirthday(date time.Time) (Birthday, error) {
	result := Birthday{date}
	if err := result.Validate(); err != nil {
		return Birthday{}, err
	}
	return result, nil
}
