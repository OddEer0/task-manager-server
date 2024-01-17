package mock

//type MockedToken struct {
//	Tokens         []*models.Token
//	FullUserToken  *models.Token
//	MinUserToken   *models.Token
//	AdminUserToken *models.Token
//}
//
//type customClaims struct {
//	JwtUserData dto.GenerateTokenDto
//	jwt.StandardClaims
//}
//
//func NewMockToken() *MockedToken {
//	mockUser := NewMockUser()
//	tokens := createMockTokensWithUser(mockUser.Users)
//
//	return &MockedToken{
//		Tokens:         tokens,
//		FullUserToken:  tokens[1],
//		MinUserToken:   tokens[0],
//		AdminUserToken: tokens[3],
//	}
//}
//
//func createMockTokensWithUser(users []*models.User) []*models.Token {
//	tokens := make([]*models.Token, len(users))
//	for _, user := range users {
//		cfg, _ := config.NewConfig()
//		refreshDuration, _ := time.ParseDuration(cfg.RefreshTokenTime)
//		refreshClaims := customClaims{
//			JwtUserData:    dto.GenerateTokenDto{Id: user.Id, Roles: user.Role},
//			StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(refreshDuration).Unix()},
//		}
//		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
//		tokens = append(tokens, &models.Token{
//			Id:    user.Id,
//			Value: refreshToken.Signature,
//		})
//	}
//	tokens = lo.Filter(tokens, func(item *models.Token, index int) bool {
//		if index == 3 {
//			return false
//		}
//		return true
//	})
//	return tokens
//}
