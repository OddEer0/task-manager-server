package mock

type MockedUser struct{}

func NewMockUser() *MockedUser {
	return &MockedUser{}
}
