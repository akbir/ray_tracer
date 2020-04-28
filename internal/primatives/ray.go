package primatives

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Point(t float64) Vector {
	// r_0 + u * t
	travel := r.Direction.MultiplyScalar(t)
	return r.Origin.Add(travel)
}
