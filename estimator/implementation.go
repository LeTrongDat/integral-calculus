package estimator 

import (
	"math"
)


type RungeEstimator struct {

}


func (est *RungeEstimator) Estimate(a float64, b float64, h float64, itg Integral, stg Strategy) (float64) {
	return math.Abs(stg.Execute(a, b, h, itg) - stg.Execute(a, b, h * 2, itg)) / 3
}

func (est *RungeEstimator) String() string {
	return "Runge estimate method"
}

func NewRungeEstimator() (*RungeEstimator) {
	return &RungeEstimator{}
}