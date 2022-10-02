package driver

import (
	"github.com/xXNurioXx/go-gta-sa-driver/driver/movement"
	"github.com/xXNurioXx/go-gta-sa-driver/math"
	"github.com/xXNurioXx/go-gta-sa-driver/math/trigonometry"
)

func (driver *Driver) DriveTo(destination math.Location) {
	//gameUtils->waitInVehicle();
	origin := driver.getLocation()

	currentAngle := driver.getCurrentAngle()

	movement.MoveForward()
	movingTo := trigonometry.GetDirectionTo(origin, destination, currentAngle)

	if movingTo == math.FORWARD {
		movement.StopMovingLeft()
		movement.StopMovingRight()
	} else if movingTo == math.LEFT {
		movement.StopMovingRight()
		movement.MoveLeft()
	} else {
		movement.StopMovingLeft()
		movement.MoveRight()
	}
}
