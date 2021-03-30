package strategy 

import (
	"math"
)


type MiddleRectangleStrategy struct {

}

func (stg *MiddleRectangleStrategy) Execute(a float64, b float64, h float64, itg Integral) (float64) {
	mul := float64(1)
	if a > b {
		a, b = b, a 
		mul = float64(-1)
	}
	n := math.Ceil((b - a) / h)
	var ans float64 
	ans = 0
	for i := 1; float64(i) <= n; i++ {
		ans += h * itg.F((a + float64(i - 1) * h + a + float64(i) * h) / 2)
	}
	return ans * mul
}

func (stg *MiddleRectangleStrategy) String() string {
	return "Middle rectangle method"
}

func NewMiddleRectangleStrategy() (*MiddleRectangleStrategy) {
	return &MiddleRectangleStrategy{}
}