package mock

import (
	"fmt"
	"strings"
	"time"

	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/google/uuid"
)

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

type MockedUser struct {
	User                           models.User
	RequiredErrUser                models.User
	MinErrUser                     models.User
	MaxErrUser                     models.User
	IdErrUser                      models.User
	EmailErrUser                   models.User
	LinkErrUser                    models.User
	RoleErrUser                    models.User
	DateCreatedAndUpdatedAtErrUser models.User
}

func NewMockUser() *MockedUser {
	pass, err := valuesobject.NewPassword("CorrectPassword1234")
	if err != nil {
		return nil
	}
	user := models.User{
		Id:             uuid.New().String(),
		Nick:           "Eer0",
		Email:          "eer0@gmail.com",
		FirstName:      "Marlen",
		LastName:       "Karimov",
		ActivationLink: fmt.Sprintf("https://%s", uuid.New().String()),
		Role:           constants.User,
	}

	errUser := user
	errUser.Id = ""
	errUser.Nick = ""
	errUser.Email = ""
	errUser.FirstName = ""
	errUser.LastName = ""
	errUser.ActivationLink = ""
	errUser.Role = ""

	user.UpdatedAt = time.Now().AddDate(0, 0, -2)
	user.CreatedAt = time.Now().AddDate(0, 0, -1)
	user.Password = pass

	errUser2 := user
	errUser2.Nick = "mi"
	errUser2.FirstName = "ma"
	minStr := "mu"
	errUser2.LastName = "mo"
	errUser2.SubTitle = &minStr
	errUser2.BanReason = &minStr

	errUser3 := user
	long50str := "lolkeklolkeklolkeklolkeklolkeklolkeklolkeklolkeklolkeklolkeklolkek"
	long500str := strings.Repeat(long50str, 10)
	long255str := strings.Repeat(long50str, 6)
	errUser3.Nick = long50str
	errUser3.FirstName = long50str
	errUser3.LastName = long50str
	errUser3.SubTitle = &long500str
	errUser3.BanReason = &long255str

	errUser4 := user
	invalidId := "invalid-id"
	errUser4.Id = invalidId

	errUser5 := user
	invalidEmail := "not-emaIL"
	errUser5.Email = invalidEmail

	errUser6 := user
	invalidLink := "no-link"
	errUser6.Avatar = &invalidLink
	errUser6.ActivationLink = invalidLink

	errUser7 := user
	invalidRole := "not-role"
	errUser7.Role = invalidRole

	errUser8 := user
	invalidCreatedAndUpdatedDate := time.Now().AddDate(0, 0, 1)
	errUser8.CreatedAt = invalidCreatedAndUpdatedDate
	errUser8.UpdatedAt = invalidCreatedAndUpdatedDate

	return &MockedUser{
		User:                           user,
		RequiredErrUser:                errUser,
		MinErrUser:                     errUser2,
		MaxErrUser:                     errUser3,
		IdErrUser:                      errUser4,
		EmailErrUser:                   errUser5,
		LinkErrUser:                    errUser6,
		RoleErrUser:                    errUser7,
		DateCreatedAndUpdatedAtErrUser: errUser8,
	}
}
