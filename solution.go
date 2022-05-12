package square

import (
	"math"
)

type figureType int

const (
	SidesTriangle figureType = 3
	SidesSquare   figureType = 4
	SidesCircle   figureType = 0
)

func CalcSquare(sideLen float64, sidesNum figureType) (res float64) {
	if sidesNum == SidesSquare {
		res = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		res = math.Pow(sideLen, 2) * math.Pi
	} else if sidesNum == SidesTriangle {
		res = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	} else {
		res = 0
	}

	return
}
