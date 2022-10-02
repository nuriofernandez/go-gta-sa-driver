package trigonometry

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func GetOffSetGrades(origin, destination math.Location) int {
	quadrant := GetRelativeQuadrant(origin, destination)
	if quadrant == CENTER {
		return 0
	}
	if quadrant == PLUS_PLUS {
		return 0
	}
	if quadrant == PLUS_MINUS {
		return 90
	}
	if quadrant == MINUS_MINUS {
		return 180
	}
	if quadrant == MINUS_PLUS {
		return 270
	}
	panic("WHAT2")
}
