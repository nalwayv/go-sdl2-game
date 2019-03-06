package vec2d

import (
	"math"

	"../../gologger"
)

// Vector2D ...
type Vector2D struct {
	x float64
	y float64
}

// NewVector2d ...
func NewVector2d(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

// GetX ...
func (v Vector2D) GetX() float64 {
	return v.x
}

// GetY ...
func (v Vector2D) GetY() float64 {
	return v.y
}

// SetX ...
func (v *Vector2D) SetX(newX float64) {
	v.x = newX
}

// SetY ...
func (v *Vector2D) SetY(newY float64) {
	v.y = newY
}

// Length ...
func Length(vec Vector2D) float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y)
}

// Add ...
func Add(v1, v2 Vector2D) *Vector2D {
	x := v1.x + v2.x
	y := v1.y + v2.y

	return &Vector2D{x, y}
}

// Sub ...
func Sub(v1, v2 Vector2D) *Vector2D {
	x := v1.x - v2.x
	y := v1.y - v2.y

	return &Vector2D{x, y}
}

// Scale ...
func Scale(vec Vector2D, scale float64) *Vector2D {
	x := vec.x * scale
	y := vec.y * scale

	return &Vector2D{x, y}
}

// Divide ...
func Divide(vec Vector2D, scale float64) *Vector2D {
	// divide by zero error
	if scale == 0.0 {
		gologger.SLogger.Fatalln("tried to divide by zero error")
	}

	x := vec.x / scale
	y := vec.y / scale

	return &Vector2D{x, y}
}

// Normalize ...
func Normalize(vec Vector2D) *Vector2D {
	len := Length(vec)

	x := vec.x
	y := vec.y

	if len > 0 {
		x *= 1 / len
		y *= 1 / len
	}

	return &Vector2D{x, y}
}
