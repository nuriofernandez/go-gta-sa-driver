package math

import "github.com/xXNurioXx/go-gta-sa-driver/math/trigonometry"

func GetDirectionTo(origin, destination Location, currentRotationAngle float64) MovingDirection {
	desiredAngle := trigonometry.GetAngleTo(origin, destination)
	closeDistance := trigonometry.CalculateAngleDistance(desiredAngle, currentRotationAngle)

	if closeDistance <= 5 && closeDistance >= -5 {
		return FORWARD
	}
	if closeDistance < 0 {
		return RIGHT
	}
	return LEFT
}
