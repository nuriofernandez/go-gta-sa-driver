package driver

import "github.com/go-vgo/robotgo"

func MoveForward() {
	robotgo.KeyToggle("w", "down")
}
func StopMovingForward() {
	robotgo.KeyToggle("w", "up")
}

func MoveBack() {
	robotgo.KeyToggle("s", "down")
}
func StopMovingBack() {
	robotgo.KeyToggle("s", "up")
}

func MoveLeft() {
	robotgo.KeyToggle("a", "down")
}
func StopMovingLeft() {
	robotgo.KeyToggle("a", "up")
}

func MoveRight() {
	robotgo.KeyToggle("d", "down")
}
func StopMovingRight() {
	robotgo.KeyToggle("d", "up")
}

func Brake() {
	robotgo.KeyToggle(" ", "down")
}
func StopBraking() {
	robotgo.KeyToggle(" ", "up")
}

func StopMoving() {
	StopBraking()
	StopMovingRight()
	StopMovingLeft()
	StopMovingBack()
	StopMovingForward()
}
