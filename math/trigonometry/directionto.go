package trigonometry

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func GetDirectionTo(origin, destination math.Location, currentRotationAngle float64) math.MovingDirection {
	desiredAngle := GetAngleTo(origin, destination)
	closeDistance := CalculateAngleDistance(desiredAngle, currentRotationAngle)

	if closeDistance <= 5 && closeDistance >= -5 {
		return math.FORWARD
	}
	if closeDistance < 0 {
		return math.RIGHT
	}
	return math.LEFT
}
