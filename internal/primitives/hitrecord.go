package primitives

import (
	"math"
)

type HitRecord struct {
	Time          float64
	Point, Normal Vector
	Material
}

type Hitable interface {
	Hit(r Ray, tMin, tMax float64) (bool, HitRecord)
}

type Sphere struct {
	Center Vector
	Radius float64
	Material
}

func (s *Sphere) Hit(r Ray, tMin, tMax float64) (bool, HitRecord) {
	a := r.Direction.DotProduct(r.Direction)
	// r_0 - c
	oc := r.Origin.Subtract(s.Center)

	// 2u * (r_0 -c)
	b := 2 * r.Direction.DotProduct(oc)

	// [||r_o - c||^2 - ||R||^2]
	c := oc.DotProduct(oc) - s.Radius*s.Radius

	// b^2 - 4ac
	discriminant := b*b - 4*a*c

	hit := HitRecord{Material: s.Material}

	if discriminant > 0 {
		// find where the sphere was hit
		// NOTE: always evaluate the smallest t first
		t := (-b - math.Sqrt(discriminant)) / (2 * a)
		if t < tMax && t > tMin {
			hit.Time = t
			hit.Point = r.Point(t)
			hit.Normal = hit.Point.Subtract(s.Center).DivideScalar(s.Radius)
			return true, hit

		}
		t = (-b + math.Sqrt(discriminant)) / (2 * a)
		if t < tMax && t > tMin {
			hit.Time = t
			hit.Point = r.Point(t)
			hit.Normal = hit.Point.Subtract(s.Center).DivideScalar(s.Radius)
			return true, hit

		}

	}
	return false, HitRecord{}
}
