package aggregate_tests

import (
	"fmt"
	"github.com/OddEer0/task-manager-server/internal/common/constants"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"github.com/google/uuid"
	"strings"
	"time"
)

type mockUser struct {
	user                           models.User
	requiredErrUser                models.User
	minErrUser                     models.User
	maxErrUser                     models.User
	idErrUser                      models.User
	emailErrUser                   models.User
	linkErrUser                    models.User
	roleErrUser                    models.User
	dateCreatedAndUpdatedAtErrUser models.User
}

func newMockUser() mockUser {
	pass, err := valuesobject.NewPassword("CorrectPassword1234")
	if err != nil {
		return mockUser{}
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

	return mockUser{
		user:                           user,
		requiredErrUser:                errUser,
		minErrUser:                     errUser2,
		maxErrUser:                     errUser3,
		idErrUser:                      errUser4,
		emailErrUser:                   errUser5,
		linkErrUser:                    errUser6,
		roleErrUser:                    errUser7,
		dateCreatedAndUpdatedAtErrUser: errUser8,
	}
}
