package service

import (
	"algo-archive/internal/model"
	"algo-archive/pkg/errcode"
	"algo-archive/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"regexp"
	"strings"
	"unicode/utf8"
)

type RegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func ComparePassword(dbPassword, password, salt string) bool {
	return strings.Compare(dbPassword, util.EncodeMD5(util.EncodeMD5(password)+salt)) == 0
}

func ValidatePassword(password string) error {
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 32 {
		return errcode.PasswordLengthLimit
	}
	return nil
}

func ValidateUsername(username string) error {
	// validation
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return errcode.UsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errcode.UsernameCharacterLimit
	}

	// duplication
	user, _ := ds.GetUserByUsername(username)
	if user.Model != nil && user.ID > 0 {
		return errcode.UserAlreadyExist
	}

	return nil
}

func Register(username, password string) (*model.User, error) {
	password, salt := EncryptPasswordAndSalt(password)

	user := &model.User{
		Nickname: username,
		Username: username,
		Password: password,
		Status:   model.UserStatusNormal,
		Salt:     salt,
	}

	user, err := ds.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// EncryptPasswordAndSalt Encrypt password and generate salt
func EncryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = util.EncodeMD5(util.EncodeMD5(password) + salt)
	return password, salt
}

func DoLogin(ctx *gin.Context, param *AuthRequest) (*model.User, error) {
	user, err := ds.GetUserByUsername(param.Username)
	if err != nil {
		return nil, errcode.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		// TODO: Redis validate tries
		// Implement: redis.Get(ctx, LOGIN_ERR_KEY, user.ID) >= MAX_LOGIN_TRIES ?

		if ComparePassword(user.Password, param.Password, user.Salt) {
			if user.Status == model.UserStatusClosed {
				return nil, errcode.UserHasBeenBanned
			}

			// Clear redis login tries

			// All good
			return user, nil
		}

		// Clear redis login failed tries

		// password does not match
		return nil, errcode.UnauthorizedAuthFailed
	}

	return nil, errcode.UnauthorizedAuthNotExist
}
