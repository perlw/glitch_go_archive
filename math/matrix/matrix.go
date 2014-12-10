package matrix

import "math"

type Matrix []float64

func identiyOrOutMatrix(out []Matrix) Matrix {
	var mat Matrix

	if len(out) > 0 {
		mat = out[0]
	} else {
		mat = NewIdentityMatrix()
	}

	return mat
}

func NewMatrix() Matrix {
	return Matrix{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

func NewIdentityMatrix() Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func NewPerspectiveMatrix(fov, ratio, nearZ, farZ float64) Matrix {
	fovRadii := (fov / 2.0) * (math.Pi / 360.0)
	f := 1.0 / (math.Tan(fovRadii))
	zDiff := farZ - nearZ

	mat := NewMatrix()
	mat[0] = f / ratio
	mat[5] = f
	mat[10] = -(farZ + nearZ) / zDiff
	mat[11] = -(2.0 * farZ * nearZ) / zDiff
	mat[14] = -1.0
	mat[15] = 0.0

	return mat
}

func NewTranslationMatrix(x, y, z float64, out ...Matrix) Matrix {
	mat := identiyOrOutMatrix(out)

	mat[12] = x
	mat[13] = y
	mat[14] = z

	return mat
}

func NewRotationMatrix(x, y, z, rot float64, out ...Matrix) Matrix {
	mat := identiyOrOutMatrix(out)
	c := math.Cos(rot)
	s := math.Sin(rot)
	t := 1.0 - c

	mat[0] = (t * (x * x)) + c
	mat[1] = (t * (x * y)) - (s * z)
	mat[2] = (t * (x * z)) + (s * y)

	mat[4] = (t * (x * y)) + (s * z)
	mat[5] = (t * (y * y)) + c
	mat[6] = (t * (y * z)) - (s * x)

	mat[8] = (t * (x * z)) - (s * y)
	mat[9] = (t * (y * z)) + (s * x)
	mat[10] = (t * (z * z)) + c

	return mat
}

func Multiply(m1, m2 Matrix, out ...Matrix) Matrix {
	mat := identiyOrOutMatrix(out)

	mat[0] = (m1[0] * m2[0]) + (m1[1] * m2[4]) + (m1[2] * m2[8]) + (m1[3] * m2[12])
	mat[1] = (m1[0] * m2[1]) + (m1[1] * m2[5]) + (m1[2] * m2[9]) + (m1[3] * m2[13])
	mat[2] = (m1[0] * m2[2]) + (m1[1] * m2[6]) + (m1[2] * m2[10]) + (m1[3] * m2[14])
	mat[3] = (m1[0] * m2[3]) + (m1[1] * m2[7]) + (m1[2] * m2[11]) + (m1[3] * m2[15])

	mat[4] = (m1[4] * m2[0]) + (m1[5] * m2[4]) + (m1[6] * m2[8]) + (m1[7] * m2[12])
	mat[5] = (m1[4] * m2[1]) + (m1[5] * m2[5]) + (m1[6] * m2[9]) + (m1[7] * m2[13])
	mat[6] = (m1[4] * m2[2]) + (m1[5] * m2[6]) + (m1[6] * m2[10]) + (m1[7] * m2[14])
	mat[7] = (m1[4] * m2[3]) + (m1[5] * m2[7]) + (m1[6] * m2[11]) + (m1[7] * m2[15])

	mat[8] = (m1[8] * m2[0]) + (m1[9] * m2[4]) + (m1[10] * m2[8]) + (m1[11] * m2[12])
	mat[9] = (m1[8] * m2[1]) + (m1[9] * m2[5]) + (m1[10] * m2[9]) + (m1[11] * m2[13])
	mat[10] = (m1[8] * m2[2]) + (m1[9] * m2[6]) + (m1[10] * m2[10]) + (m1[11] * m2[14])
	mat[11] = (m1[8] * m2[3]) + (m1[9] * m2[7]) + (m1[10] * m2[11]) + (m1[11] * m2[15])

	mat[12] = (m1[12] * m2[0]) + (m1[13] * m2[4]) + (m1[14] * m2[8]) + (m1[15] * m2[12])
	mat[13] = (m1[12] * m2[1]) + (m1[13] * m2[5]) + (m1[14] * m2[9]) + (m1[15] * m2[13])
	mat[14] = (m1[12] * m2[2]) + (m1[13] * m2[6]) + (m1[14] * m2[10]) + (m1[15] * m2[14])
	mat[15] = (m1[12] * m2[3]) + (m1[13] * m2[7]) + (m1[14] * m2[11]) + (m1[15] * m2[15])

	return mat
}

func (m *Matrix) Translate(x, y, z float64) {
	mat := NewMatrix()
	copy(mat, (*m))
	Multiply(mat, NewTranslationMatrix(x, y, z), (*m))
}

func (m *Matrix) Rotate(x, y, z, rot float64) {
	mat := NewMatrix()
	copy(mat, (*m))
	Multiply(mat, NewRotationMatrix(x, y, z, rot), (*m))
}

func (m Matrix) ToGL() []float32 {
	return []float32{
		float32(m[0]), float32(m[1]), float32(m[2]), float32(m[3]),
		float32(m[4]), float32(m[5]), float32(m[6]), float32(m[7]),
		float32(m[8]), float32(m[9]), float32(m[10]), float32(m[11]),
		float32(m[12]), float32(m[13]), float32(m[14]), float32(m[15]),
	}
}
