package enums

type UserID int64

type BandID int64

type BandRole int32

const (
	BandRoleNone BandRole = iota
	BandRoleVocal
	BandRoleGuitar
	BandRoleBass
	BandRoleDrum
	BandRoleKeyboard
	BandRoleDigitalInstrument
	BandRoleStringInstrument
	BandRoleBrassInstrument
	BandRoleWoodwindInstrument
	BandRoleOthers
)

type Genre int32

const (
	GenreNone Genre = iota
	GenreRock
	GenrePop
	GenreJazz
	GenreElectronic
)
