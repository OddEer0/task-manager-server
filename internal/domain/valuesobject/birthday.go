package valuesobject

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	MinTime = time.Date(1900, 1, 10, 0, 0, 0, 0, time.UTC)
	MaxTime = time.Now().AddDate(-12, 0, 0)
)

type Birthday struct {
	Value time.Time
}

func (b *Birthday) Validate() error {
	if b.Value.IsZero() {
		return errors.New("birthday is zero")
	}
	if b.Value.Before(MinTime) || b.Value.After(MaxTime) {
		return errors.New("invalid time")
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
