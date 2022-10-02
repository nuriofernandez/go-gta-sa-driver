package main

import (
	"fmt"
	"github.com/xXNurioXx/go-gta-sa-driver/driver"
	"github.com/xXNurioXx/go-gta-sa-driver/math"
)

func main() {
	driver := driver.New()
	gtaSa := driver.GtaSa

	vehicleLocation := gtaSa.GetVehicle().GetLocation()
	fmt.Printf("Vehicle location: '%f', '%f', '%f'\n", vehicleLocation.X, vehicleLocation.Y, vehicleLocation.Z)

	// Testing ciruit
	driver.DriveToPos(math.Location{
		X: -1278,
		Y: -292,
	})
	driver.DriveToPos(math.Location{
		X: -1398,
		Y: -154,
	})
	driver.DriveToPos(math.Location{
		X: -1282,
		Y: 9,
	})

	// Repeat
	driver.DriveToPos(math.Location{
		X: -1278,
		Y: -292,
	})
	driver.DriveToPos(math.Location{
		X: -1398,
		Y: -154,
	})
	driver.DriveToPos(math.Location{
		X: -1282,
		Y: 9,
	})

}
