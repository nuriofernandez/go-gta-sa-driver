package trigonometry

import (
	math2 "github.com/xXNurioXx/go-gta-sa-driver/math"
	"math"
)

func GetAngleTo(origin, destination math2.Location) float64 {
	adjacent := GetAdjacentLineValue(origin, destination)
	hypotenuse := GetHypotenuseLineValue(origin, destination)

	angle := math.Acos(adjacent/hypotenuse) * 180.0 / math.Pi
	return angle + float64(GetOffSetGrades(origin, destination))
}
