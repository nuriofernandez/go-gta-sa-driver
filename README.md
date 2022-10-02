# GTA SA Vehicle driver

A simple and easy to use GTA SA vehicle driver. 🚗

## Usage example

```go
func main() {
	driver := driver.New()
	driver.DriveToPos(math.Location{X: -2137, Y: -351})
	driver.DriveToPos(math.Location{X: -2073, Y: -356})
	driver.DriveToPos(math.Location{X: -1991, Y: -374})
	driver.DriveToPos(math.Location{X: -1943, Y: -412})
	driver.DriveToPos(math.Location{X: -1926, Y: -460})
	driver.DriveToPos(math.Location{X: -1920, Y: -502})
	driver.DriveToPos(math.Location{X: -1914, Y: -578})
	driver.DriveToPos(math.Location{X: -1913, Y: -745})
	driver.DriveToPos(math.Location{X: -1913, Y: -919})
	driver.DriveToPos(math.Location{X: -1910, Y: -1039})
	driver.DriveToPos(math.Location{X: -1908, Y: -1132})
	driver.DriveToPos(math.Location{X: -1907, Y: -1305})
}
```

## Video demonstration

<img src="https://i.imgur.com/3qYsubs.gif" alt="Car driving thru checkpoints" width="100%"/>

## Clarifications

At the previous example there were very few checkpoints.

In order to generate a smooth drive the amount of checkpoints must be significantly higher than the one used for this example. 
This can be achieved by using a pathfinder or something similar.

The purpose of this repository is only to take control of the car. Not to map the game's map or know where to drive to.