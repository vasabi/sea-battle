package v1

import (
	"fmt"
	"sea-battle/internal/pkg/game"
)

func getGameState() GameState {
	return GameState{
		ShipCount: game.GetShipCount(),
		Destroyed: game.GetDestroyed(),
		Knocked:   game.GetKnocked(),
		ShotCount: game.GetShotCount(),
	}
}

func clearGameField() (msg string) {
	if game.GetGameFieldSize() > 0 {
		size := game.GetGameFieldSize()
		msg = fmt.Sprintf("recreated game field with size %dX%d", size, size)
		game.CreateGameField(size)

		return
	}

	msg = "nothing to clear"
	return
}
