package primitives

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func (u Vector) RGBA() (r, g, b, a uint32) {
	// Sqrt() for gamma-2 correction
	r = uint32(math.Sqrt(u.X) * 0xffff)
	g = uint32(math.Sqrt(u.Y) * 0xffff)
	b = uint32(math.Sqrt(u.Z) * 0xffff)
	a = 0xffff
	return
}

var UnitVector = Vector{1, 1, 1}

func VectorInUnitSphere(rand *rand.Rand) Vector {
	for {
		r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
		p := r.MultiplyScalar(2.0).Subtract(UnitVector)
		if p.DotProduct(p) >= 1.0 {
			return p
		}
	}
}

func (u Vector) Add(v Vector) Vector {
	return Vector{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

func (u Vector) Subtract(v Vector) Vector {
	return Vector{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

func (u Vector) DotProduct(v Vector) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func (u Vector) Multiply(v Vector) Vector {
	return Vector{u.X * v.X, u.Y * v.Y, u.Z * v.Z}
}

func (u Vector) Length() float64 {
	return math.Sqrt(u.DotProduct(u))
}

func (u Vector) Normalise() Vector {
	l := u.Length()
	return Vector{u.X / l, u.Y / l, u.Z / l}
}

func (u Vector) AddScalar(a float64) Vector {
	return Vector{u.X + a, u.Y + a, u.Z + a}
}

func (u Vector) MultiplyScalar(a float64) Vector {
	return Vector{u.X * a, u.Y * a, u.Z * a}
}

func (u Vector) DivideScalar(a float64) Vector {
	return Vector{u.X / a, u.Y / a, u.Z / a}
}

func (u Vector) Refract(normal Vector, ni_over_nt float64) (bool, Vector) {
	uv := u.Normalise()
	un := normal.Normalise()
	vdotn := uv.DotProduct(un)

	// 1 - (n1/n2)^2 [1-(v.n)^2]
	discriminant := 1 - (ni_over_nt * ni_over_nt * (1 - vdotn*vdotn))

	if discriminant > 0 { // we have refraction
		// n1/n2 (v - (v.n)N) - rt(discriminant) N
		first_term := uv.Subtract(un.MultiplyScalar(vdotn)).MultiplyScalar(ni_over_nt)
		second_term := un.MultiplyScalar(math.Sqrt(discriminant))
		return true, first_term.Subtract(second_term)
	}

	return false, Vector{}

}

func (u Vector) Reflect(n Vector) Vector {
	b := 2 * u.DotProduct(n)
	return u.Subtract(n.MultiplyScalar(b))
}
