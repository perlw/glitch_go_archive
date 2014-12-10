package vector

import (
	"github.com/doxxan/glitch/math/matrix"
	"math"
)

type Vector4 struct {
	X, Y, Z, W float64
}

func Add4(v1, v2 Vector4) Vector4 {
	return Vector4{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z, W: v1.W + v2.W}
}

func Sub4(v1, v2 Vector4) Vector4 {
	return Vector4{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z, W: v1.W - v2.W}
}

func Mul4(v1, v2 Vector4) Vector4 {
	return Vector4{X: v1.X * v2.X, Y: v1.Y * v2.Y, Z: v1.Z * v2.Z, W: v1.W * v2.W}
}

func Div4(v1, v2 Vector4) Vector4 {
	return Vector4{X: v1.X / v2.X, Y: v1.Y / v2.Y, Z: v1.Z / v2.Z, W: v1.W / v2.W}
}

func AddScalar4(v Vector4, s float64) Vector4 {
	return Vector4{X: v.X + s, Y: v.Y + s, Z: v.Z + s, W: v.W + s}
}

func SubScalar4(v Vector4, s float64) Vector4 {
	return Vector4{X: v.X - s, Y: v.Y - s, Z: v.Z - s, W: v.W - s}
}

func MulScalar4(v Vector4, s float64) Vector4 {
	return Vector4{X: v.X * s, Y: v.Y * s, Z: v.Z * s, W: v.W * s}
}

func DivScalar4(v Vector4, s float64) Vector4 {
	return Vector4{X: v.X / s, Y: v.Y / s, Z: v.Z / s, W: v.W / s}
}

func Dot4(v1, v2 Vector4) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z) + (v1.W * v2.W)
}

func Cross4(v1, v2 Vector4) Vector4 {
	v := Vector4{}
	v.X = (v1.Y * v2.Z) - (v1.Z * v2.Y)
	v.Y = (v1.Z * v2.X) - (v1.X * v2.Z)
	v.Z = (v1.X * v2.Y) - (v1.Y * v2.X)
	return v
}

func Normalize4(v Vector4) Vector4 {
	return DivScalar4(v, v.Length())
}

func MultiplyMatrix4(v Vector4, m matrix.Matrix) Vector4 {
	result := Vector4{}

	result.X = m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W
	result.Y = m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W
	result.Z = m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W
	result.W = m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W

	return result
}

func (v Vector4) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W))
}

func (v Vector4) ToGL() []float32 {
	return []float32{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}
