package trigonometry

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func GetOppositeLineValue(origin, destination math.Location) float64 {
	quadrant := GetRelativeQuadrant(origin, destination)

	if quadrant == PLUS_PLUS || quadrant == MINUS_MINUS {
		return math.DistanceSingleAxis(origin.X, destination.X)
	}
	if quadrant == PLUS_MINUS || quadrant == MINUS_PLUS {
		return math.DistanceSingleAxis(origin.Y, destination.Y)
	}
	panic("WHAT4")
}
