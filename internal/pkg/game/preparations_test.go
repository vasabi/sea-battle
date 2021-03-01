package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckGameFieldAndCreateNewGameField(t *testing.T) {
	a := assert.New(t)

	a.Error(isGameFieldCreatedAndEmpty(), "game field must not be created")

	CreateGameField(11)
	a.NoError(isGameFieldCreatedAndEmpty(), "game field must be crated")

	a.Equal(11, len(gameField), "gameField length must be 16")
	a.Equal(11, len(gameField[0]), "gameField[0] length must be 16")

	for _, cell := range gameField[0] {
		a.Equal(0, cell, "all elements of gameField[0] must be 0")
	}
}

func TestFleetNotCreatedWhenAlreadyExists(t *testing.T) {
	a := assert.New(t)

	CreateGameField(3)
	gameField[1][1] = ShipPart
	a.EqualErrorf(CreateFleet("1A 2B, 3C 4D"), "the game has already started", "must return error")
}

func TestFleetNotCreatedWithBadCoordinates(t *testing.T) {
	a := assert.New(t)

	CreateGameField(3)
	a.EqualErrorf(CreateFleet("1A 2B, 3C 4C"), "the ship should be a pair of coordinates instead of [ 3C 4C]", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,4C 3C"), "the ship 4C 3C is not square or rectangular", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,3C 5C"), "some coordinates between 3C and 5C outside the playing field", "must return error")
	a.EqualErrorf(CreateFleet("123A 2B,3C 5C"), "left top coordinates. invalid coordinates 123A", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,3C 5er4C"), "right bottom coordinates. invalid coordinates 5er4C", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,2B 3C"), "cell 2B is already taken by another ship", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,3A 3C"), "the ship 3A 3C cannot touch another ship: cell 2A", "must return error")
	a.EqualErrorf(CreateFleet("1A 2B,3B 3D"), "some coordinates between 3B and 3D outside the playing field", "must return error")
	a.EqualErrorf(CreateFleet("1A 1D,3B 3C"), "some coordinates between 1A and 1D outside the playing field", "must return error")
}

func TestFleetCreatedNearWallsOK(t *testing.T) {
	a := assert.New(t)

	CreateGameField(3)
	a.NoError(CreateFleet("1A 1C,3B 3B"))
}
