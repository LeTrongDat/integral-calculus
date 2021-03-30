package strategy

import (
	"math"
)


type LeftRectangleStrategy struct {

}

func (stg *LeftRectangleStrategy) Execute(a float64, b float64, h float64, itg Integral) (float64) {
	mul := float64(1) 
	if a > b {
		a, b = b, a
		mul = float64(-1) 
	}
	var ans float64
	ans = 0
	n := math.Ceil((b - a) / h)
	for i := 0; float64(i) < n; i++ {
		ans += h * itg.F(a + float64(i) * h)
	}
	return ans * mul
}

func (stg *LeftRectangleStrategy) String() string {
	return "Left rectangle method"
}

func NewLeftRectangleStrategy() (*LeftRectangleStrategy) {
	return &LeftRectangleStrategy{}
}