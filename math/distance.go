package math

import "math"

func DistanceSingleAxis(originX, destinationX float64) float64 {
	return math.Sqrt(Square(originX - destinationX))
}

func DistanceTwoAxis(originX, originY, destinationX, destinationY float64) float64 {
	return math.Sqrt(Square(originX-destinationX) + Square(originY-destinationY))
}
