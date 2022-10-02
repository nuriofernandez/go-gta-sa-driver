# GTA SA Vehicle driver

A simple and easy to use GTA SA vehicle driver. 🚗

## Usage example

```go
func main() {
	driver := driver.New()
	driver.DriveToPos(-2137, -351)
	driver.DriveToPos(-2073, -356)
	driver.DriveToPos(-1991, -374)
	driver.DriveToPos(-1943, -412)
	driver.DriveToPos(-1926, -460)
	driver.DriveToPos(-1920, -502)
	driver.DriveToPos(-1914, -578)
	driver.DriveToPos(-1913, -745)
	driver.DriveToPos(-1913, -919)
	driver.DriveToPos(-1910, -1039)
	driver.DriveToPos(-1908, -1132)
	driver.DriveToPos(-1907, -1305)
}
```

## Video demonstration

<img src="./.github/documentation/self-driving-car-demo.gif" alt="Car driving thru checkpoints" width="100%"/>

## Clarifications

At the previous example there were very few checkpoints.

In order to generate a smooth drive the amount of checkpoints must be significantly higher than the one used for this example. 
This can be achieved by using a pathfinder or something similar.

The purpose of this repository is only to take control of the car. Not to map the game's map or know where to drive to.