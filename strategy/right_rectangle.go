package strategy

import (
	"math"
)

type RightRectangleStrategy struct {

}

func (stg *RightRectangleStrategy) Execute(a float64, b float64, h float64, itg Integral) (float64) {
	mul := float64(1)
	if a > b {
		a, b = b, a
		mul = float64(-1)
	}
	n := math.Ceil((b - a) / h)
	var ans float64 
	ans = 0
	for i := 1; float64(i) <= n; i++ {
		ans += h * itg.F(a + float64(i) * h)
	}
	return ans * mul
}

func (stg *RightRectangleStrategy) String() string {
	return "Right rectangle method"
}

func NewRightRectangleStrategy() (*RightRectangleStrategy) {
	return &RightRectangleStrategy{}
}