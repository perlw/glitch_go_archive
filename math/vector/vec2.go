package vector

import "math"

type Vec2 struct {
	X, Y float64
}

func Add2(v1, v2 Vec2) Vec2 {
	return Vec2{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func Sub2(v1, v2 Vec2) Vec2 {
	return Vec2{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func Mul2(v1, v2 Vec2) Vec2 {
	return Vec2{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func Div2(v1, v2 Vec2) Vec2 {
	return Vec2{X: v1.X / v2.X, Y: v1.Y / v2.Y}
}

func AddScalar2(v Vec2, s float64) Vec2 {
	return Vec2{X: v.X + s, Y: v.Y + s}
}

func SubScalar2(v Vec2, s float64) Vec2 {
	return Vec2{X: v.X - s, Y: v.Y - s}
}

func MulScalar2(v Vec2, s float64) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func DivScalar2(v Vec2, s float64) Vec2 {
	return Vec2{X: v.X / s, Y: v.Y / s}
}

func Dot2(v1, v2 Vec2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Cross2(v1, v2 Vec2) Vec2 {
	return Vec2{
		X: v1.Y*v2.X - v1.X*v2.Y,
		Y: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Normalize2(v Vec2) Vec2 {
	if l := v.Length(); l > 0.0001 {
		return DivScalar2(v, l)
	}

	return v
}

func (v *Vec2) Add(v1 Vec2) {
	v.X += v1.X
	v.Y += v1.Y
}

func (v *Vec2) Sub(v1 Vec2) {
	v.X -= v1.X
	v.Y -= v1.Y
}

func (v *Vec2) Mul(v1 Vec2) {
	v.X *= v1.X
	v.Y *= v1.Y
}

func (v *Vec2) Div(v1 Vec2) {
	v.X /= v1.X
	v.Y /= v1.Y
}

func (v *Vec2) AddScalar(s float64) {
	v.X += s
	v.Y += s
}

func (v *Vec2) SubScalar(s float64) {
	v.X -= s
	v.Y -= s
}

func (v *Vec2) MulScalar(s float64) {
	v.X *= s
	v.Y *= s
}

func (v *Vec2) DivScalar(s float64) {
	v.X /= s
	v.Y /= s
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Neg() Vec2 {
	return Vec2{X: -v.X, Y: -v.Y}
}

func (v Vec2) ToGL() []float32 {
	return []float32{float32(v.X), float32(v.Y)}
}
