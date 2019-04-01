package helper

import "math"

// PiOver2 ...
const PiOver2 float64 = (math.Pi / 2.0)

// PiOver4 ...
const PiOver4 float64 = (math.Pi / 4.0)

// TwoPi ...
const TwoPi float64 = (math.Pi * 2.0)

// Clamp ...
func Clamp(value, min, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

// Lerp ...
func Lerp(norm, min, max float64) float64 {
	return (max-min)*norm + min
}

// ToDegrees ...
func ToDegrees(radians float64) float64 {
	val := 57.295779513082320876798154814105 // 180/pi
	return radians * val
}

// ToRadians ...
func ToRadians(degrees float64) float64 {
	val := 0.017453292519943295769236907684886 // 180/pi
	return degrees * val
}
