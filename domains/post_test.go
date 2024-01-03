package domains

import (
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils"
	"study-kafka-ddb/utils/deftype"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeMessage(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"success":                        testSuccess,
		"fails due to invalid post type": testInvalidPostTypeError,
	}

	for scenario, testFn := range testCases {
		t.Run(scenario, testFn)
	}
}

func testSuccess(t *testing.T) {
	post := &Post{ID: 1, Type: []enums.PostType{
		enums.PostTypeWanted,
		enums.PostTypeResume,
	}[utils.RandIntFromTo(0, 1)]}
	_, err := post.MakeMessage()
	assert.Nil(t, err)
}

func testInvalidPostTypeError(t *testing.T) {
	post := &Post{ID: 1, Type: []enums.PostType{
		enums.PostTypeNone,
		enums.PostType(utils.RandIntFromTo(3, 100)),
	}[utils.RandIntFromTo(0, 1)]}
	_, err := post.MakeMessage()
	assert.Equal(t, deftype.ErrInvalidRequestData, err)
}

func TestIsExpired(t *testing.T) {
	type testCase struct {
		scenario string
		post     *Post
		expected bool
	}
	testCases := []*testCase{
		{
			scenario: "expired",
			post: &Post{
				CreatedAt: time.Now().Add(-validityPeriodOfPost - 24),
			},
			expected: true,
		},
		{
			scenario: "not expired",
			post: &Post{
				CreatedAt: time.Now().Add(-2 * 24 * time.Hour),
			},
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.post.IsExpired())
		})
	}
}
