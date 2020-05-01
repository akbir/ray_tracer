package primitives

import "testing"

func TestRay_Point(t *testing.T) {
	ray := Ray{Origin: Vector{1,2,3}, Direction: Vector{1,2,3}}
	point := ray.Point(5)
	exp := Vector{6,12,18}
	if point != exp{
		t.Errorf("Ray path error got: %v, expected: %v", point, exp)
	}
}
