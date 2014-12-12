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

func CirclevsCircle(a, b *Circle) (bool, *Manifold) {
	// Vector from a to b
	normal := vector.Sub2(b.position, a.position)

	r := a.radius + b.radius
	r *= r

	// Initial check if circles are overlapping
	if math.Pow(a.position.X+b.position.X, 2)+math.Pow(a.position.X+b.position.Y, 2) > r {
		return false, nil
	}

	// We have overlap, compute manifold
	m := Manifold{}
	/*m.a = a
	m.b = b*/

	if d := normal.Length(); d != 0.0 {
		// Distance is difference between radius and distance
		m.penetration = r - d

		// Utilize d since we performed sqrt on it already
		// Points from a to b and is a unit vector
		m.normal = vector.DivScalar2(normal, d)
	} else {
		m.penetration = a.radius
		m.normal = vector.Vec2{X: 1.0, Y: 0.0}
	}

	return true, &m
}

type Body struct {
	Position    vector.Vec2
	Velocity    vector.Vec2
	restitution float64

	force vector.Vec2

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

	rv := vector.Sub2(b.Velocity, a.Velocity)

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
	a.Velocity.Sub(vector.MulScalar2(impulse, ratioA))
	b.Velocity.Add(vector.MulScalar2(impulse, ratioB))
}

func (m *Manifold) PositionalCorrection() {
	percent := 0.2
	slop := 0.01

	a := m.b1
	b := m.b2

	correction := m.normal
	correction.MulScalar(math.Max(m.penetration-slop, 0.0) / (a.invMass + b.invMass))
	correction.MulScalar(percent)
}

type World struct {
	bodies []*Body

	contacts []*Manifold

	gravity    vector.Vec2
	iterations int
	dt         float64
}

func NewWorld(gravity vector.Vec2, iterations int, dt float64) *World {
	return &World{
		gravity:    gravity,
		iterations: iterations,
		dt:         dt,
	}
}

func (w World) Step() {
	// Broadphase

	// Integrate forces
	for _, body := range w.bodies {
		if body.invMass == 0.0 {
			continue
		}

		acceleration := vector.Add2(vector.MulScalar2(body.force, body.invMass), w.gravity)
		body.Velocity.Add(vector.MulScalar2(acceleration, w.dt/2))
		// Angular Velocity
	}

	// Collisions

	// Integrate Velocity
	for _, body := range w.bodies {
		if body.invMass == 0.0 {
			continue
		}

		body.Position.Add(vector.MulScalar2(body.Velocity, w.dt))
		// Orientation
	}

	// Correct positions

}

func (w *World) AddBody(body *Body) {
	w.bodies = append(w.bodies, body)
}
