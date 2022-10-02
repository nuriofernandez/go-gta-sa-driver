package driver

import "github.com/xXNurioXx/go-gta-sa-driver/math"

func (driver *Driver) getLocation() math.Location {
	gtaSaLocation := driver.GtaSa.GetVehicle().GetLocation()
	return math.Location{
		X: float64(gtaSaLocation.X),
		Y: float64(gtaSaLocation.Y),
	}
}
