package vo_tests

import (
	"errors"
	"fmt"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TODO - сделать тесты на маршалинг json
func TestBirthday(t *testing.T) {
	testCases := []struct {
		name  string
		date  time.Time
		error error
	}{
		{
			name:  "Should correct create values object",
			date:  time.Now().AddDate(-19, 0, 0),
			error: nil,
		},
		{
			name:  "Should error with birthday time small 12 year ago",
			date:  time.Now().AddDate(-10, 0, 0),
			error: errors.New(valuesobject.BirthdayIncorrectMax),
		},
		{
			name:  fmt.Sprintf("Should error with birthday small %d year", valuesobject.MinBirthdayTime.Year()),
			date:  valuesobject.MinBirthdayTime.AddDate(-1, 0, 0),
			error: errors.New(valuesobject.BirthdayIncorrectMin),
		},
		{
			name:  "Should error with zero date",
			date:  time.Time{},
			error: errors.New(valuesobject.BirthdayIsZero),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := valuesobject.NewBirthday(tc.date)
			assert.Equal(t, err, tc.error)
			if tc.error == nil {
				assert.Equal(t, result, valuesobject.Birthday{Value: tc.date})
			} else {
				assert.Equal(t, result, valuesobject.Birthday{})
			}
		})
	}
}
