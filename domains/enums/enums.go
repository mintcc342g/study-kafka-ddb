package enums

type UserID string

type BandID int32

type BandPosition int32

const (
	BandPositionNone BandPosition = iota
	BandPositionVocal
	BandPositionGuitar
	BandPositionBass
	BandPositionDrum
	BandPositionKeyboard
	BandPositionDigitalInstrument
	BandPositionStringInstrument
	BandPositionBrassInstrument
	BandPositionWoodwindInstrument
	BandPositionOthers
)

type Genre int32

const (
	GenreNone Genre = iota
	GenreRock
	GenrePop
	GenreJazz
	GenreElectronic
)

type PostType int32

const (
	PostTypeNone PostType = iota
	PostTypeWanted
	PostTypeResume
)
