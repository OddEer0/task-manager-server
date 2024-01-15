package vo_tests

import (
	"errors"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

// TODO - сделать тесты на маршалинг json
func TestPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		error    error
	}{
		{
			name:     "Should correct create values object",
			password: "mysupperpupper1234",
			error:    nil,
		},
		{
			name:     "Should error len < 8",
			password: "small1",
			error:    errors.New(valuesobject.PasswordMinLength),
		},
		{
			name:     "Should error len > 30",
			password: "veryveryveryveryveryveryverylong123456789",
			error:    errors.New(valuesobject.PasswordMaxLength),
		},
		{
			name:     "Should error incorrect password without number",
			password: "notnumber",
			error:    errors.New(valuesobject.PasswordInvalid),
		},
		{
			name:     "Should error incorrect password without symbol",
			password: "123456789",
			error:    errors.New(valuesobject.PasswordInvalid),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := valuesobject.NewPassword(tc.password)
			assert.Equal(t, err, tc.error)
			if tc.error == nil {
				err2 := bcrypt.CompareHashAndPassword([]byte(result.Value), []byte(tc.password))
				assert.Equal(t, nil, err2)
			} else {
				assert.Equal(t, result, valuesobject.Password{})
			}
		})
	}
}
