package trigonometry

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func GetHypotenuseLineValue(origin, destination math.Location) float64 {
	return math.DistanceTwoAxis(origin.X, origin.Y, destination.X, destination.Y)
}
