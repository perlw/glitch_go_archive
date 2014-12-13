package physics

import (
	"github.com/doxxan/glitch/math/vector"
	"math"
)

type AABB struct {
	Min, Max vector.Vec2
}

func AABBvsAABB(m *Manifold) bool {
	a := m.b1
	b := m.b2

	normal := vector.Sub2(b.Position, a.Position)

	abox := a.Box
	bbox := b.Box

	aXExtent := (abox.Max.X - abox.Min.X) / 2.0
	bXExtent := (bbox.Max.X - bbox.Min.X) / 2.0
	aYExtent := (abox.Max.Y - abox.Min.Y) / 2.0
	bYExtent := (bbox.Max.Y - bbox.Min.Y) / 2.0

	// Calculate overlap
	xOverlap := aXExtent + bXExtent - math.Abs(normal.X)
	yOverlap := aYExtent + bYExtent - math.Abs(normal.Y)
	if xOverlap <= 0.0 || yOverlap <= 0.0 {
		return false
	}

	// Find axis of least penetration
	if xOverlap < yOverlap {
		// Point towards b knowing that normal points from a to b
		if normal.X < 0.0 {
			m.normal = vector.Vec2{X: -1.0, Y: 0.0}
		} else {
			m.normal = vector.Vec2{X: 1.0, Y: 0.0}
		}
		m.penetration = xOverlap
	} else {
		// Point towards b knowing that normal points from a to b
		if normal.Y < 0.0 {
			m.normal = vector.Vec2{X: 0.0, Y: -1.0}
		} else {
			m.normal = vector.Vec2{X: 0.0, Y: 1.0}
		}
		m.penetration = yOverlap
	}

	return true
}

type Body struct {
	Position vector.Vec2
	Velocity vector.Vec2

	restitution     float64
	staticFriction  float64
	dynamicFriction float64

	Force vector.Vec2

	mass    float64
	invMass float64

	Box AABB
}

func NewBody(mass float64) *Body {
	b := Body{}

	b.mass = mass
	if mass == 0.0 {
		b.invMass = 0.0
	} else {
		b.invMass = 1.0 / mass
	}

	b.restitution = 0.2
	b.staticFriction = 0.5
	b.dynamicFriction = 0.3

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
	gravityDt := vector.MulScalar2(vector.Vec2{X: 0.0, Y: -9.8}, 1.0/60.0)
	if (rv.X*rv.X + rv.Y*rv.Y) < (gravityDt.X*gravityDt.X+gravityDt.Y*gravityDt.Y)+0.0001 {
		e = 0.0
	}

	// Impulse scalar
	j := -(1.0 + e) * velAlongNormal
	j /= a.invMass + b.invMass

	// Apply impulse
	impulse := vector.MulScalar2(m.normal, j)
	a.Velocity.Sub(vector.MulScalar2(impulse.Neg(), a.invMass))
	b.Velocity.Add(vector.MulScalar2(impulse, b.invMass))

	// Friction impulse
	rv = vector.Sub2(b.Velocity, a.Velocity)

	t := vector.Sub2(rv, vector.MulScalar2(m.normal, vector.Dot2(rv, m.normal)))
	t = vector.Normalize2(t)

	// Tangent magnitude
	jt := -vector.Dot2(rv, t)
	jt /= a.invMass + b.invMass

	// Skip tiny friction impulses
	if math.Abs(jt) < 0.0001 {
		return
	}

	// Coulumb's law
	sf := math.Sqrt(a.staticFriction * b.staticFriction)
	df := math.Sqrt(a.dynamicFriction * b.dynamicFriction)
	var tImpulse vector.Vec2
	if math.Abs(jt) < j*sf {
		tImpulse = vector.MulScalar2(t, jt)
	} else {
		tImpulse = vector.MulScalar2(vector.MulScalar2(t, -j), df)
	}

	a.Velocity.Sub(vector.MulScalar2(tImpulse.Neg(), a.invMass))
	b.Velocity.Add(vector.MulScalar2(tImpulse, b.invMass))
}

func (m *Manifold) PositionalCorrection() {
	percent := 0.2
	slop := 0.01

	a := m.b1
	b := m.b2

	correction := m.normal
	correction.MulScalar(math.Max(m.penetration-slop, 0.0) / (a.invMass + b.invMass))
	correction.MulScalar(percent)

	a.Position.Sub(vector.MulScalar2(correction, a.invMass))
	b.Position.Add(vector.MulScalar2(correction, b.invMass))
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
	w.contacts = nil
	for index, b1 := range w.bodies {
		for _, b2 := range w.bodies[index+1:] {
			if b1.invMass == 0.0 && b2.invMass == 0 {
				continue
			}

			m := &Manifold{
				b1: b1,
				b2: b2,
			}
			if AABBvsAABB(m) {
				w.contacts = append(w.contacts, m)
			}
		}
	}

	// Integrate forces
	for _, body := range w.bodies {
		if body.invMass == 0.0 {
			continue
		}

		acceleration := vector.Add2(vector.MulScalar2(body.Force, body.invMass), w.gravity)
		body.Velocity.Add(vector.MulScalar2(acceleration, w.dt/2.0))
		// Angular Velocity
	}

	// Collisions
	for t := 0; t < w.iterations; t++ {
		for _, manifold := range w.contacts {
			manifold.ResolveCollision()
		}
	}

	// Integrate Velocity
	for _, body := range w.bodies {
		if body.invMass == 0.0 {
			continue
		}

		body.Position.Add(vector.MulScalar2(body.Velocity, w.dt))
		// Orientation

		acceleration := vector.Add2(vector.MulScalar2(body.Force, body.invMass), w.gravity)
		body.Velocity.Add(vector.MulScalar2(acceleration, w.dt/2.0))

		body.Force = vector.Vec2{X: 0.0, Y: 0.0}
	}

	// Correct positions
	for _, manifold := range w.contacts {
		manifold.PositionalCorrection()
	}
}

func (w *World) AddBody(body *Body) {
	w.bodies = append(w.bodies, body)
}
