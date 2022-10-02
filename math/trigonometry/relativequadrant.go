package trigonometry

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func GetRelativeQuadrant(origin, destination math.Location) PositionQuadrant {
	if origin.X == destination.X && origin.Y == destination.Y {
		return PLUS_PLUS
	}
	if destination.X >= origin.X && destination.Y > origin.Y {
		return PLUS_PLUS
	}
	if destination.X > origin.X && destination.Y <= origin.Y {
		return PLUS_MINUS
	}
	if destination.X < origin.X && destination.Y <= origin.Y {
		return MINUS_MINUS
	}
	if destination.X <= origin.X && destination.Y != origin.Y {
		return MINUS_PLUS
	}
	panic("WHAT")
}
