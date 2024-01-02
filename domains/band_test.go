package domains

import (
	"strings"
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils/deftype"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenPosition(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"success":                          testOpenPositionSuccess,
		"fails due to lack of permissions": testOpenPositionPermissionError,
		"fails due to invalid contents":    testOpenPositionInvalidContentsError,
	}

	for scenario, testFn := range testCases {
		t.Run(scenario, testFn)
	}
}

func testOpenPositionSuccess(t *testing.T) {
	// request
	user := &User{ID: enums.UserID("tomori")}
	position := enums.BandPositionVocal
	contents := "Will you play in this band for the rest of your life?"
	bandID := enums.BandID(1)

	// result
	post := NewWantedPost(user.ID, contents, position, bandID)

	// test
	band := &Band{ID: bandID, ReaderID: user.ID}
	res, err := band.OpenPosition(user, position, contents)
	assert.Nil(t, err)
	assert.Equal(t, res.BandID, post.BandID)
	assert.Equal(t, res.Position, post.Position)
	assert.Equal(t, res.IsOpened, post.IsOpened)
}

func testOpenPositionPermissionError(t *testing.T) {
	// request
	user := &User{ID: enums.UserID("ra-na")}
	position := enums.BandPositionGuitar
	contents := "I am searching for a place where I belong."
	bandID := enums.BandID(1)

	// result
	err := deftype.ErrUnauthorized

	// test
	band := &Band{ID: bandID, ReaderID: enums.UserID("tomori")}
	_, res := band.OpenPosition(user, position, contents)
	assert.Equal(t, err, res)
}

func testOpenPositionInvalidContentsError(t *testing.T) {
	// request
	user := &User{ID: enums.UserID("tomori")}
	position := enums.BandPositionVocal
	contents := strings.Repeat("band-aid", 110)
	bandID := enums.BandID(1)

	// result
	err := deftype.ErrInvalidRequestData

	// test
	band := &Band{ID: bandID, ReaderID: user.ID}
	_, res := band.OpenPosition(user, position, contents)
	assert.Equal(t, err, res)
}
