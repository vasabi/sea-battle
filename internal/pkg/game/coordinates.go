package game

import (
	"fmt"
	"strconv"
	"strings"
)

// convertCoordinates - converting head and tail coordinates of the ship to index (12D -> 12,3).
func convertCoordinates(coordinates string) (x, y int, err error) {
	switch {
	case len(coordinates) == 2:
		x = int(coordinates[0] - 49)
		y = int(strings.ToLower(coordinates)[1] - 97)
	case len(coordinates) == 3:
		x, err = strconv.Atoi(strconv.Itoa(int(coordinates[0]-48)) + strconv.Itoa(int(coordinates[1]-49)))
		y = int(strings.ToLower(coordinates)[2] - 97)
	default:
		err = fmt.Errorf("invalid coordinates %s", coordinates)
	}

	return
}

// validateCoordinates - returns error if the ship cannot be created.
func validateCoordinates(xHead, yHead, xTail, yTail int) (err error) {
	if xHead < 0 || xHead >= len(gameField) ||
		xTail < 0 || xTail >= len(gameField) ||
		yHead < 0 || yHead >= len(gameField) ||
		yTail < 0 || yTail >= len(gameField) {
		err = fmt.Errorf(
			"some coordinates between %d%s and %d%s outside the playing field",
			xHead+1,
			strings.ToUpper(string(int32(yHead+97))),
			xTail+1,
			strings.ToUpper(string(int32(yTail+97))),
		)
	}

	if xTail < xHead || yTail < yHead {
		err = fmt.Errorf(
			"the ship %d%s %d%s is not square or rectangular",
			xHead+1,
			strings.ToUpper(string(int32(yHead+97))),
			xTail+1,
			strings.ToUpper(string(int32(yTail+97))),
		)
	}

	return
}
