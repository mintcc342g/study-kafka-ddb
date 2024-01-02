package domains

import (
	"regexp"
	"strings"
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils"
	"study-kafka-ddb/utils/deftype"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        enums.UserID // NOTE: uuid
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string // NOTE: unique
	Password  string
}

func NewUser() *User {
	return &User{}
}

func (r *User) SignUp(name, email, password string) deftype.Error {
	if !isValidName(name) || !ValidateUserInput(email, password) {
		return deftype.ErrInvalidRequestData
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		zap.S().Error("fail to generate UUID", err)
		return deftype.ErrInternalServerError
	}

	r.Password, err = hashPassword(password)
	if err != nil {
		zap.S().Error("fail to hash password", err)
		return deftype.ErrInternalServerError
	}

	r.ID = enums.UserID(uuid)
	r.Name = name
	r.Email = email
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	return nil
}

func (r *User) SignIn(email, password string) deftype.Error {
	if !strings.EqualFold(r.Email, email) {
		return deftype.ErrInvalidRequestData
	}

	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(password)); err != nil {
		return deftype.ErrInvalidRequestData
	}

	return nil
}

func hashPassword(password string) (string, deftype.Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", deftype.ErrInternalServerError
	}

	return string(hash), nil
}

func ValidateUserInput(email, password string) bool {
	return isValidEmail(email) && isValidPassword(password)
}

func isValidName(name string) bool {
	return name != "" && len(name) <= 50
}

func isValidEmail(email string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z0-9]{1,150}@[a-zA-Z0-9]{1,150}\.[a-zA-Z]{2,}$`, email)

	return ok && err == nil
}

func isValidPassword(password string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z\d]{3,8}$`, password)

	return ok && err == nil
}

func (r *User) SeekPosition(contents string, position enums.BandPosition, genre enums.Genre) (*Post, deftype.Error) {
	if !isValidContents(contents) {
		return nil, deftype.ErrInvalidRequestData
	}

	return NewResumePost(r.ID, contents, position, genre), nil
}
