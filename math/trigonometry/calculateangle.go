package trigonometry

import "math"

func CalculateAngleDistance(angle1, angle2 float64) float64 {
	diff := math.Mod(angle2-angle1+180, 360) - 180
	if diff < -180 {
		return diff + 360
	}
	return diff
}
