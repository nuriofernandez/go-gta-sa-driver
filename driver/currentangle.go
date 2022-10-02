package driver

import (
	"github.com/xXNurioXx/go-gta-sa-driver/math/trigonometry"
)

func (driver *Driver) getCurrentAngle() float64 {
	return trigonometry.CalculateAngleByLookAndGrad(
		float64(driver.gtaApi.GetPedVehiclePositionAngleYLook()),
		float64(driver.gtaApi.GetPedVehiclePositionAngleXLook()),
		float64(driver.gtaApi.GetPedVehiclePositionAngleXGrad()),
		float64(driver.gtaApi.GetPedVehiclePositionAngleZGrad()),
	)
}
