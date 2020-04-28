package primatives

import (
	"math"
	"math/rand"
)

type Material interface {
	Bounce(ray Ray, hit HitRecord, rand *rand.Rand) (bool, Ray)
	Color() Vector
}

type Lambertian struct {
	C Vector
}

func (l Lambertian) Bounce(input Ray, record HitRecord, rand *rand.Rand)(bool, Ray){
	direction := record.Normal.Add(VectorInUnitSphere(rand))
	return true, Ray{record.Point, direction}
}

func (l Lambertian) Color() Vector {
	return l.C
}


type Metal struct {
	C    Vector
	Fuzz float64
}

func (m Metal) Bounce(input Ray, record HitRecord, rand *rand.Rand)(bool, Ray){
	direction := input.Direction.Reflect(record.Normal)
	fuzzed := VectorInUnitSphere(rand).MultiplyScalar(m.Fuzz)
	bounced_ray := Ray{record.Point, direction.Add(fuzzed)}
	bounced := direction.DotProduct(record.Normal) > 0
	return bounced, bounced_ray
}

func (m Metal) Color() Vector {
	return m.C
}


type Dielectric struct {
	C               Vector
	RefractiveIndex float64
}


func (d Dielectric) Color() Vector {
	return d.C
}

func (d Dielectric) Bounce(input Ray, record HitRecord, rand *rand.Rand) (bool, Ray){
	var outwardNormal Vector
	var niOverNt, cosine float64

	// check if entering or leaving a surface
	if input.Direction.DotProduct(record.Normal) > 0 {
		outwardNormal = record.Normal.MultiplyScalar(-1)
		niOverNt = d.RefractiveIndex
		cosine = input.Direction.Normalise().DotProduct(record.Normal) * d.RefractiveIndex
	} else {
		outwardNormal = record.Normal
		niOverNt = 1.0 / d.RefractiveIndex
		cosine = -1 * input.Direction.Normalise().DotProduct(record.Normal) * d.RefractiveIndex
	}

	// check for success
	success, refract := input.Direction.Refract(outwardNormal, niOverNt)
	var reflectProb float64

	if success {
		// find out if we have reflection from success
		reflectProb = d.schlick(cosine)

	} else {
		// must be pure reflection
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		reflected := input.Direction.Reflect(record.Normal)
		return true, Ray{record.Point, reflected}
	}

	return true, Ray{record.Point, refract}

}

func (d Dielectric) schlick(cosine float64) float64 {
	r0 := (1 - d.RefractiveIndex) / (1 + d.RefractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0) * math.Pow(1-cosine, 5)
}



