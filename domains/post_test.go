package domains

import (
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils"
	"study-kafka-ddb/utils/deftype"
	"testing"

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
