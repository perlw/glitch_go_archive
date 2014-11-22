package vector

import (
	"shmup/math/matrix"
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 1, Y: 2, Z: 3}
	tv := Vector{X: 2, Y: 4, Z: 6}
	result := Add(v1, v2)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestSub(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 1, Y: 2, Z: 3}
	tv := Vector{X: 0, Y: 0, Z: 0}
	result := Sub(v1, v2)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestMul(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 1, Y: 2, Z: 3}
	tv := Vector{X: 1, Y: 4, Z: 9}
	result := Mul(v1, v2)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestDiv(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3, W: 1}
	v2 := Vector{X: 1, Y: 2, Z: 3, W: 1}
	tv := Vector{X: 1, Y: 1, Z: 1, W: 1}
	result := Div(v1, v2)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestAddScalar(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	s := float32(2.0)
	tv := Vector{X: 3, Y: 4, Z: 5, W: 2}
	result := AddScalar(v, s)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestSubScalar(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	s := float32(2.0)
	tv := Vector{X: -1, Y: 0, Z: 1, W: -2}
	result := SubScalar(v, s)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestMulScalar(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	s := float32(2.0)
	tv := Vector{X: 2, Y: 4, Z: 6}
	result := MulScalar(v, s)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestDivScalar(t *testing.T) {
	v := Vector{X: 2, Y: 2, Z: 4}
	s := float32(2.0)
	tv := Vector{X: 1, Y: 1, Z: 2}
	result := DivScalar(v, s)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 4, Y: 5, Z: 6}
	tv1 := float32(31.09)
	tv2 := float32(32.01)
	result := Dot(v1, v2)

	if result <= tv1 || result >= tv2 {
		t.Errorf("Expected %f < x < %f, got %f", tv1, tv2, result)
	}
}

func TestCross(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 4, Y: 5, Z: 6}
	tv := Vector{X: -3, Y: 6, Z: -3}
	result := Cross(v1, v2)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}

func TestLength(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	tv1 := float32(3.74)
	tv2 := float32(3.75)
	result := v.Length()

	if result <= tv1 || result >= tv2 {
		t.Errorf("Expected %f < x < %f, got %f", tv1, tv2, result)
	}
}

func TestNormalize(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	tv1 := Vector{X: 0.26, Y: 0.53, Z: 0.80, W: -0.01}
	tv2 := Vector{X: 0.27, Y: 0.54, Z: 0.81, W: 0.01}
	result := Normalize(v)

	if (result.X <= tv1.X || result.X >= tv2.X) || (result.Y <= tv1.Y || result.Y >= tv2.Y) || (result.Z <= tv1.Z || result.Z >= tv2.Z) || (result.W <= tv1.W || result.W >= tv2.W) {
		t.Errorf("Expected %v < x < %v, got %v", tv1, tv2, result)
	}
}

func TestMultiplyMatrix(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	m := matrix.NewIdentityMatrix()
	tv := Vector{X: 1, Y: 2, Z: 3}
	result := MultiplyMatrix(v, m)

	if result != tv {
		t.Errorf("Expected %v, got %v", tv, result)
	}
}
