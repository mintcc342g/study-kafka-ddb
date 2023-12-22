package domains

import (
	"study-kafka-ddb/utils/deftype"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	type testCase struct {
		scenario string

		name     string
		email    string
		password string

		result deftype.Error
	}
	testCases := []testCase{
		{
			scenario: "success",
			name:     "tomori",
			email:    "confeito1122@mygo.com",
			password: "qwer1234",
			result:   nil,
		},
		{
			scenario: "fail due to an invalid name",
			name:     "",
			result:   deftype.ErrInvalidRequestData,
		},
		{
			scenario: "fail due to an invalid email",
			name:     "anon",
			email:    "anontokyo",
			result:   deftype.ErrInvalidRequestData,
		},
		{
			scenario: "fail due to an invalid password",
			name:     "rana",
			email:    "matcha@mygo.com",
			password: "a",
			result:   deftype.ErrInvalidRequestData,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			user := NewUser()
			res := user.SignUp(tc.name, tc.email, tc.password)
			assert.Equal(t, tc.result, res)
		})
	}
}
