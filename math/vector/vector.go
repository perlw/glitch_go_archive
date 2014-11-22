package vector

import (
	"math"
	"shmup/math/matrix"
)

type Vector struct {
	X, Y, Z, W float32
}

func Add(v1, v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z, W: v1.W + v2.W}
}

func Sub(v1, v2 Vector) Vector {
	return Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z, W: v1.W - v2.W}
}

func Mul(v1, v2 Vector) Vector {
	return Vector{X: v1.X * v2.X, Y: v1.Y * v2.Y, Z: v1.Z * v2.Z, W: v1.W * v2.W}
}

func Div(v1, v2 Vector) Vector {
	return Vector{X: v1.X / v2.X, Y: v1.Y / v2.Y, Z: v1.Z / v2.Z, W: v1.W / v2.W}
}

func AddScalar(v Vector, s float32) Vector {
	return Vector{X: v.X + s, Y: v.Y + s, Z: v.Z + s, W: v.W + s}
}

func SubScalar(v Vector, s float32) Vector {
	return Vector{X: v.X - s, Y: v.Y - s, Z: v.Z - s, W: v.W - s}
}

func MulScalar(v Vector, s float32) Vector {
	return Vector{X: v.X * s, Y: v.Y * s, Z: v.Z * s, W: v.W * s}
}

func DivScalar(v Vector, s float32) Vector {
	return Vector{X: v.X / s, Y: v.Y / s, Z: v.Z / s, W: v.W / s}
}

func Dot(v1, v2 Vector) float32 {
	return (v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z) + (v1.W * v2.W)
}

func Cross(v1, v2 Vector) Vector {
	v := Vector{}
	v.X = (v1.Y * v2.Z) - (v1.Z * v2.Y)
	v.Y = (v1.Z * v2.X) - (v1.X * v2.Z)
	v.Z = (v1.X * v2.Y) - (v1.Y * v2.X)
	return v
}

func Normalize(v Vector) Vector {
	return DivScalar(v, v.Length())
}

func MultiplyMatrix(v Vector, m matrix.Matrix) Vector {
	result := Vector{}

	result.X = m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W
	result.Y = m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W
	result.Z = m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W
	result.W = m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W

	return result
}

func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W))))
}
