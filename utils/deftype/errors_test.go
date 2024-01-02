package deftype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	type testCase struct {
		scenario string

		err     Error
		compare Error

		expected bool
	}
	testCases := []*testCase{
		{
			scenario: "success",
			err:      ErrNotFound,
			compare:  ErrNotFound,
			expected: true,
		},
		{
			scenario: "fail",
			err:      ErrDuplicatedRequest,
			compare:  ErrNotFound,
			expected: false,
		},
		{
			scenario: "nil check",
			err:      ErrDuplicatedRequest,
			compare:  nil,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			result := tc.err.Equal(tc.compare)
			assert.Equal(t, tc.expected, result)
		})
	}
}
