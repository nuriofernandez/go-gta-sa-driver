package trigonometry

import "math"

func CalculateAngleByLookAndGrad(angleyLook, anglexLook, anglexGrad, anglezGrad float64) float64 {
	if (angleyLook >= 0 && anglexLook >= 0) || (angleyLook < 0 && anglexLook > 0) {
		return math.Acos(anglexGrad/math.Cos(math.Asin(anglezGrad))) * 180.0 / 3.1415
	}
	return 360 - math.Acos(anglexGrad/math.Cos(math.Asin(anglezGrad)))*180.0/3.1415
}
