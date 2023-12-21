package domains

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils"
	"study-kafka-ddb/utils/deftype"
	"time"

	"go.uber.org/zap"
)

type User struct {
	ID        enums.UserID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
}

func NewUser() *User {
	return &User{}
}

func (r *User) SignUp(name, email, password string) deftype.Error {
	if !ValidateUserInput(name, email, password) {
		return deftype.ErrInvalidRequestData
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		zap.S().Error("fail to generate UUID", err)
		return deftype.ErrInternalServerError
	}

	r.ID = enums.UserID(uuid)
	r.Name = name
	r.Email = email
	r.Password = hashPassword(password)
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	return nil
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}

func ValidateUserInput(name, email, password string) bool {
	return isValidName(name) && isValidEmail(email) && isValidPassword(password)
}

func isValidName(name string) bool {
	return len(name) <= 50
}

func isValidEmail(email string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z0-9]{1,150}@[a-zA-Z0-9]{1,150}\.[a-zA-Z]{2,}$`, email)

	return ok || err == nil
}

func isValidPassword(password string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z\d]{8,}$`, password)

	return ok || err == nil
}
