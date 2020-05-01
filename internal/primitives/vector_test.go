package primitives

import (
	"testing"
)

func TestVector_Add(t *testing.T){
	v1 := Vector{1,2,3}
	v2 := Vector{4,5,6}
	sum1 := v1.Add(v2)
	sum2 := v2.Add(v1)
	if sum1 != sum2 {
		t.Errorf("Vector addition is not cummative, got %v and %v", sum1, sum2)
	}
	ans := Vector{5,7,9}
	if sum1 != ans{
		t.Errorf("Vector addition error got %v and %v", sum1, ans)
	}
}

func TestVector_Subtract(t *testing.T){
	v1 := Vector{1,2,3}
	v2 := Vector{4,5,6}
	sum1 := v1.Subtract(v2)
	sum2 := v2.Subtract(v1)
	if sum1 == sum2 {
		t.Errorf("Vector subtract is cummative, got %v and %v", sum1, sum2)
	}
	ans := Vector{3,3,3}
	if sum2 != ans{
		t.Errorf("Vector addition error got %v and %v", sum1, ans)

	}
}

func TestVector_DivideScalar(t *testing.T) {
	v1 := Vector{6,9,12}
	d := float64(3)
	exp := Vector{2, 3, 4}
	ans := v1.DivideScalar(d)
	if ans != exp {
		t.Errorf("Vector scalar division error got %v and %v", ans, exp)
	}

}