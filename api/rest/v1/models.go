package v1

import (
	"fmt"
	"sea-battle/internal/pkg/game"
)

type GameFieldReq struct {
	Size int `json:"range"`
}

func (f *GameFieldReq) CreateGameField() (msg string) {
	var size = 26

	// Check max range (the number of characters in the Latin alphabet).
	if f.Size <= 26 {
		size = f.Size
		msg = fmt.Sprintf("Your range is %d. Created game field with size %dX%d", f.Size, f.Size, f.Size)
	} else {
		msg = fmt.Sprintf(
			"Max size of the game field is 26X26. Your range is %d. Created game field with max size %dX%d",
			f.Size,
			size,
			size,
		)
	}

	game.CreateGameField(size)
	return
}

type CreateFleetReq struct {
	Coordinates string `json:"Coordinates"`
}

func (f *CreateFleetReq) CreateFleet() error {
	return game.CreateFleet(f.Coordinates)
}

type ShotReq struct {
	Coord string `json:"сoord"`
}

func (s *ShotReq) Shot() (ShotResultResp, error) {
	destroy, knock, end, err := game.Shot(s.Coord)
	return ShotResultResp{
		Destroy: destroy,
		Knock:   knock,
		End:     end,
	}, err
}

type ShotResultResp struct {
	Destroy bool `json:"destroy"`
	Knock   bool `json:"knock"`
	End     bool `json:"end"`
}

type GameState struct {
	ShipCount int `json:"ship_count"`
	Destroyed int `json:"destroyed"`
	Knocked   int `json:"knocked"`
	ShotCount int `json:"shot_count"`
}
