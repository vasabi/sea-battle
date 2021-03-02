package game

import (
	"fmt"
	"strings"
)

func Shot(coord string) (destroy, knock, end bool, err error) {
	var x, y int

	defer printGameField()

	if shipCount-destroyed == 0 {
		err = fmt.Errorf("lay down your arms the war is over")
		return
	}

	if x, y, err = convertCoordinates(coord); err != nil {
		return
	}

	if x < 0 || x > len(gameField) || y < 0 || y > len(gameField) {
		err = fmt.Errorf(
			"coordinates between %d%s outside the playing field",
			x+1,
			strings.ToUpper(string(int32(y+97))),
		)
		return
	}

	if gameField[x][y] == Miss || gameField[x][y] == DamagedShipPart {
		err = fmt.Errorf("a shot has already been fired at these coordinates: %s", coord)
		return
	}

	if gameField[x][y] == Empty {
		gameField[x][y] = Miss
	} else if gameField[x][y] == ShipPart {
		destroy, knock, end = handleHit(x, y)
	}

	shotCount++
	return
}

func handleHit(x, y int) (destroy, knock, end bool) {
	var xHead, xTail, yHead, yTail, damaged int

leftTop:
	for i := x - 1; ; i-- {
		if i < 0 || gameField[i][y] == Empty || gameField[i][y] == Miss {
			xHead = i + 1
			for j := y - 1; ; j-- {
				if j < 0 || gameField[x][j] == Empty || gameField[x][j] == Miss {
					yHead = j + 1
					break leftTop
				}
			}
		}
	}

rightBottom:
	for i := x + 1; ; i++ {
		if i >= len(gameField) || gameField[i][y] == Empty || gameField[i][y] == Miss {
			xTail = i - 1
			for j := y; ; j++ {
				if j >= len(gameField) || gameField[x][j] == Empty || gameField[x][j] == Miss {
					yTail = j - 1
					break rightBottom
				}
			}
		}
	}

	if gameField[x][y] == ShipPart {
		gameField[x][y] = DamagedShipPart
		knock = true
	}

	for i := xHead; i <= xTail; i++ {
		for j := yHead; j <= yTail; j++ {
			if gameField[i][j] == DamagedShipPart {
				damaged++
			}
		}
	}

	if damaged == 1 {
		knocked++
	}

	if damaged == (xTail+1-xHead)*(yTail+1-yHead) {
		destroy = true
		destroyed++
	}

	if shipCount-destroyed == 0 {
		end = true
	}

	return
}
