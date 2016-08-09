package pt

import "math"

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
	Box      Box
}

func NewSphere(center Vector, radius float64, material Material) Shape {
	min := Vector{center.X - radius, center.Y - radius, center.Z - radius}
	max := Vector{center.X + radius, center.Y + radius, center.Z + radius}
	box := Box{min, max}
	return &Sphere{center, radius, material, box}
}

func (s *Sphere) Compile() {
}

func (s *Sphere) BoundingBox() Box {
	return s.Box
}

func (s *Sphere) Intersect(r Ray) Hit {
	to := r.Origin.Sub(s.Center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > 0 {
			return Hit{s, t1, nil}
		}
		// t2 := -b + d
		// if t2 > 0 {
		// 	return Hit{s, t2}
		// }
	}
	return NoHit
}

func (s *Sphere) ColorAt(p Vector) Color {
	if s.Material.Texture == nil {
		return s.Material.Color
	}
	p = p.Sub(s.Center)
	u := math.Atan2(p.Z, p.X)
	v := math.Atan2(p.Y, Vector{p.X, 0, p.Z}.Length())
	u = (u + math.Pi) / (2 * math.Pi)
	v = (v + math.Pi/2) / math.Pi
	return s.Material.Texture.Sample(u, v)
}

func (s *Sphere) MaterialAt(p Vector) Material {
	return s.Material
}

func (s *Sphere) NormalAt(p Vector) Vector {
	return p.Sub(s.Center).Normalize()
}
