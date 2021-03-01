package game

var (
	shipCount, destroyed, knocked, shotCount int
)

func GetShipCount() int {
	return shipCount
}

func GetShotCount() int {
	return shotCount
}

func GetDestroyed() int {
	return destroyed
}

func GetKnocked() int {
	return knocked
}

func resetStats() {
	shipCount = 0
	destroyed = 0
	knocked = 0
	shotCount = 0
}
