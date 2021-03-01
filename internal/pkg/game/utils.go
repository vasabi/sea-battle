package game

import (
	"fmt"
	"strconv"
	"strings"
)

func printGameField() {
	var (
		horizontalCoordinates, line string
		i, j                        int
	)

	for ind := range gameField {
		horizontalCoordinates += strings.ToUpper(string(int32(ind + 97)))
		line += "_"
	}

	fmt.Println("\n" + horizontalCoordinates)
	fmt.Println(line)

	for i = 0; i < len(gameField); i++ {
		var horizontalString string

		for j = 0; j < len(gameField); j++ {
			horizontalString += strconv.Itoa(gameField[i][j])
		}

		fmt.Println(horizontalString + "| " + strconv.Itoa(i+1))
	}

	fmt.Println()
}
