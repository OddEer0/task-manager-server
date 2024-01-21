package mock

import "github.com/OddEer0/task-manager-server/internal/domain/models"

type MockedToken struct {
	Tokens []*models.Token
}

func NewMockToken() *MockedToken {
	return &MockedToken{
		Tokens: []*models.Token{},
	}
}
