package mock

//type MockedUser struct {
//	Users                   []*models.User
//	FullUser                *models.User
//	MinUser                 *models.User
//	BannedUser              *models.User
//	AdminUser               *models.User
//	Registration1User       *dto.RegistrationInputDto
//	Registration1UserResult *authUsecase.AuthResult
//}
//
//func NewMockUser() *MockedUser {
//	return &MockedUser{
//		Users:      getMockUsers(),
//		FullUser:   getMockUsers()[1],
//		MinUser:    getMockUsers()[0],
//		BannedUser: getMockUsers()[2],
//		AdminUser:  getMockUsers()[3],
//		Registration1User: &dto.RegistrationInputDto{
//			Nick:      "Eer0",
//			Password:  "LeagueOfLegends5757",
//			Email:     "Lolkek@gmail.com",
//			FirstName: "Marlen",
//			LastName:  "Karimov",
//		},
//		Registration1UserResult: &authUsecase.AuthResult{
//			User: &dto.ResponseUserDto{
//				Id:        "not",
//				Nick:      "Eer0",
//				Email:     "Lolkek@gmail.com",
//				FirstName: "Marlen",
//				LastName:  "Karimov",
//				Role:      constants.User,
//			},
//			Tokens: tokenService.JwtTokens{AccessToken: "dsads", RefreshToken: "dsadsad"},
//		},
//		//&aggregate.UserAggregate{User: models.User{
//		//	Id:             "xz",
//		//	Nick:           "Eer0",
//		//	Password:       "LeagueOfLegends5757",
//		//	Email:          valuesobject.Email{Value: "Lolkek@gmail.com"},
//		//	FirstName:      "Marlen",
//		//	LastName:       "Karimov",
//		//	ActivationLink: "https://fjdskfjdlsfpdaksoad.com",
//		//	Role:           constants.User,
//		//},
//		//},
//	}
//}
//
//func getMockUsers() []*models.User {
//	secondSubTitle := "Wc3 pro gamer"
//	secondAvatar := "Wc3 pro gamer"
//	secondBirthday := time.Now()
//	thirdSubTitle := "Toxic my profession"
//	thirdAvatar := "Wc3 pro gamer"
//	thirdBirthday := time.Now()
//	thirdBanReason := "toxic!!"
//
//	return []*models.User{
//		{
//			Id:             "first",
//			Nick:           "Singer",
//			Password:       "dsadsagwfsadasvdasgfdasfdafda",
//			Email:          valuesobject.Email{Value: "singer@gmail.com"},
//			FirstName:      "John",
//			LastName:       "Martin",
//			SubTitle:       nil,
//			Avatar:         nil,
//			Birthday:       nil,
//			Role:           constants.User,
//			ActivationLink: "https://fakeapi.com",
//			IsActivate:     true,
//			IsBanned:       false,
//			BanReason:      nil,
//			CreatedAt:      time.Now(),
//			UpdatedAt:      time.Now(),
//		},
//		{
//			Id:             "second",
//			Nick:           "Player",
//			Password:       "gsgkll;sfogzfdlkl;fdsl;fmd;f",
//			Email:          valuesobject.Email{Value: "player@gmail.com"},
//			FirstName:      "Foggy",
//			LastName:       "Happy",
//			SubTitle:       &secondSubTitle,
//			Avatar:         &secondAvatar,
//			Birthday:       &valuesobject.Birthday{Value: secondBirthday},
//			Role:           constants.User,
//			ActivationLink: "https://fakeapi.com",
//			IsActivate:     true,
//			IsBanned:       false,
//			BanReason:      nil,
//			CreatedAt:      time.Now(),
//			UpdatedAt:      time.Now(),
//		},
//		{
//			Id:             "third",
//			Nick:           "toxic",
//			Password:       "gsgkll;sfogzfdlkl;fdsl;fmd;f",
//			Email:          valuesobject.Email{Value: "toxic@gmail.com"},
//			FirstName:      "Moon",
//			LastName:       "Romantic",
//			SubTitle:       &thirdSubTitle,
//			Avatar:         &thirdAvatar,
//			Birthday:       &valuesobject.Birthday{Value: thirdBirthday},
//			Role:           constants.User,
//			ActivationLink: "https://fakeapi.com",
//			IsActivate:     true,
//			IsBanned:       true,
//			BanReason:      &thirdBanReason,
//			CreatedAt:      time.Now(),
//			UpdatedAt:      time.Now(),
//		},
//		{
//			Id:             "four",
//			Nick:           "boss",
//			Password:       "gsgkll;sfogzfdlkl;fdsl;fmd;f",
//			Email:          valuesobject.Email{Value: "boss@gmail.com"},
//			FirstName:      "Ricardo",
//			LastName:       "Milos",
//			SubTitle:       nil,
//			Avatar:         nil,
//			Birthday:       nil,
//			Role:           constants.Admin,
//			ActivationLink: "https://fakeapi.com",
//			IsActivate:     true,
//			IsBanned:       false,
//			BanReason:      nil,
//			CreatedAt:      time.Now(),
//			UpdatedAt:      time.Now(),
//		},
//	}
//}
