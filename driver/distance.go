package driver

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func (driver *Driver) getDistanceTo(destination math.Location) float64 {
	origin := driver.getLocation()
	return math.DistanceTwoAxis(origin.X, origin.Y, destination.X, destination.Y)
}
