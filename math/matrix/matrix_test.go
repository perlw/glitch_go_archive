package matrix

import "testing"

func floatInRange(val, target float32) bool {
	return (val > (target-0.01) && val < (target+0.01))
}

func TestIdentity(t *testing.T) {
	m := NewIdentityMatrix()
	tm := Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}

func TestPerspective(t *testing.T) {
	m := NewPerspectiveMatrix(45, 1.33, 1.0, 100.0)
	tm := Matrix{
		3.77, 0, 0, 0,
		0, 5.02, 0, 0,
		0, 0, -1.02, -2.02,
		0, 0, -1, 0,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}

func TestTranslation(t *testing.T) {
	m := NewTranslationMatrix(1, 2, 3)
	tm := Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}

func TestRotation(t *testing.T) {
	m := NewRotationMatrix(1, 0, 0, 1)
	tm := Matrix{
		1, 0, 0, 0,
		0, 0.54, -0.84, 0,
		0, 0.84, 0.54, 0,
		0, 0, 0, 1,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}

func TestMatrixTranslate(t *testing.T) {
	m := NewIdentityMatrix()
	m.Translate(1, 2, 3)
	tm := Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}

func TestMatrixRotate(t *testing.T) {
	m := NewIdentityMatrix()
	m.Rotate(1, 0, 0, 1)
	tm := Matrix{
		1, 0, 0, 0,
		0, 0.54, -0.84, 0,
		0, 0.84, 0.54, 0,
		0, 0, 0, 1,
	}

	for i := 0; i < 16; i++ {
		if !floatInRange(m[i], tm[i]) {
			t.Errorf("Expected %v, got %v", tm, m)
			break
		}
	}
}
