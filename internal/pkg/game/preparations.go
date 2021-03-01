package game

import (
	"fmt"
	"strings"
)

var gameField [][]int

// CreateGameField creates new game field.
func CreateGameField(size int) {
	resetStats()

	gameField = make([][]int, size, size)

	for i := 0; i < size; i++ {
		gameField[i] = make([]int, size, size)

		for j := 0; j < size; j++ {
			gameField[i][j] = Empty
		}
	}
}

// GetGameFieldSize returns the game field size
func GetGameFieldSize() int {
	return len(gameField)
}

// CreateFleet creates one or more possible ships at the game field.
func CreateFleet(coordinates string) (err error) {
	if err = isGameFieldCreatedAndEmpty(); err != nil {
		return
	}

	coordinatesPairArr := strings.Split(coordinates, ",")

	for _, pair := range coordinatesPairArr {
		coordinatesArr := strings.Split(pair, " ")

		if len(coordinatesArr) != 2 {
			err = fmt.Errorf("the ship should be a pair of coordinates instead of %v", coordinatesArr)
			CreateGameField(GetGameFieldSize())
			return
		}

		if err = createShip(coordinatesArr[0], coordinatesArr[1]); err != nil {
			CreateGameField(GetGameFieldSize())
			return
		}
	}

	return
}

func createShip(head, tail string) (err error) {
	var xHead, yHead, xTail, yTail int

	if xHead, yHead, err = convertCoordinates(head); err != nil {
		err = fmt.Errorf("left top coordinates. %w", err)
		return
	}

	if xTail, yTail, err = convertCoordinates(tail); err != nil {
		err = fmt.Errorf("right bottom coordinates. %w", err)
		return
	}

	if err = validateCoordinates(xHead, yHead, xTail, yTail); err != nil {
		return
	}

	if err = checkShipPlacement(xHead, yHead, xTail, yTail); err != nil {
		return
	}

	if err = checkEmptyCellsAroundShip(xHead, yHead, xTail, yTail); err != nil {
		err = fmt.Errorf(
			"the ship %d%s %d%s cannot touch another ship: %w",
			xHead+1,
			strings.ToUpper(string(int32(yHead+97))),
			xTail+1,
			strings.ToUpper(string(int32(yTail+97))),
			err,
		)

		return
	}

	for i := xHead; i <= xTail; i++ {
		for j := yHead; j <= yTail; j++ {
			gameField[i][j] = ShipPart
		}
	}

	shipCount++
	printGameField()
	return
}

func isGameFieldCreatedAndEmpty() (err error) {
	if len(gameField) == 0 {
		err = fmt.Errorf("the playing field has not yet been created")
		return
	} else {
		for i := 0; i < len(gameField); i++ {
			for j := 0; j < len(gameField); j++ {
				if gameField[i][j] != Empty {
					err = fmt.Errorf("the game has already started")
					return
				}
			}
		}
	}

	return
}

func checkShipPlacement(xHead, yHead, xTail, yTail int) (err error) {
checkShipPlacement:
	for i := xHead; i <= xTail; i++ {
		for j := yHead; j <= yTail; j++ {
			if gameField[i][j] != Empty {
				err = fmt.Errorf(
					"cell %d%s is already taken by another ship",
					i+1,
					strings.ToUpper(string(int32(j+97))))
				break checkShipPlacement
			}
		}
	}

	return
}

func checkEmptyCellsAroundShip(xHead, yHead, xTail, yTail int) (err error) {
	for i := xHead - 1; i <= xTail+1; i++ {
		for j := yHead - 1; j <= yTail+1; j++ {
			err = fmt.Errorf("cell %d%s", i+1, strings.ToUpper(string(int32(j+97))))

			if i >= 0 && i <= xHead && j >= 0 && j < len(gameField) && gameField[i][j] != Empty {
				return
			}
			if i < len(gameField) && i >= xTail && j >= 0 && j < len(gameField) && gameField[i][j] != Empty {
				return
			}
			if j >= 0 && j <= yHead && i >= 0 && i < len(gameField) && gameField[i][j] != Empty {
				return
			}
			if j < len(gameField) && j >= yTail && i >= 0 && i < len(gameField) && gameField[i][j] != Empty {
				return
			}
		}
	}

	err = nil
	return
}
