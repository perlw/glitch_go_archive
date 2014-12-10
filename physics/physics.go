package physics

import (
	"github.com/doxxan/glitch/math/vector"
	"math"
)

type AABB struct {
	min, max vector.Vec2
}

func AABBvsAABB(a, b AABB) bool {
	// Exit with no intersection if separated along axis
	if a.max.X < b.min.X || a.min.X > b.max.X {
		return false
	}
	if a.max.Y < b.min.Y || a.min.Y > b.max.Y {
		return false
	}

	// No separation found, overlapping
	return true
}

type Circle struct {
	radius   float64
	position vector.Vec2
}

func CirclevsCircle(a, b Circle) bool {
	r := a.radius + b.radius
	r *= r
	return (r < math.Pow(a.position.X+b.position.X, 2)+math.Pow(a.position.X+b.position.Y, 2))
}

type Body struct {
	velocity    vector.Vec2
	restitution float64

	mass    float64
	invMass float64
}

func NewBody(mass float64) *Body {
	b := Body{}

	b.mass = mass
	if mass == 0.0 {
		b.invMass = 0.0
	} else {
		b.invMass = 1.0 / mass
	}

	return &b
}

type Manifold struct {
	b1, b2      *Body
	penetration float64
	normal      vector.Vec2
}

func (m *Manifold) ResolveCollision() {
	a := m.b1
	b := m.b2

	rv := vector.Sub2(b.velocity, a.velocity)

	velAlongNormal := vector.Dot2(rv, m.normal)
	if velAlongNormal > 0.0 {
		return
	}

	// Restitution
	e := math.Min(a.restitution, b.restitution)

	// Impulse scalar
	j := -(1.0 + e) * velAlongNormal
	j /= a.invMass + b.invMass

	// Apply impulse
	impulse := vector.MulScalar2(m.normal, j)
	massSum := a.mass + b.mass
	ratioA := a.mass / massSum
	ratioB := b.mass / massSum
	a.velocity.Sub(vector.MulScalar2(impulse, ratioA))
	b.velocity.Add(vector.MulScalar2(impulse, ratioB))
}

func (m *Manifold) PositionalCorrection() {
	/*percent := 0.2
	  slop := 0.01

	  a := m.b1
	  b := m.b2*/

	/*Vec2 correction = max( penetration - k_slop, 0.0f ) / (A.inv_mass + B.inv_mass)) * percent * n
	  A.position -= A.inv_mass * correction
	  B.position += B.inv_mass * correction*/
}
