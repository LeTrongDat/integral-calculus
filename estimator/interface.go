package estimator 

import (
	"itmo/strategy"
	"itmo/integral"
)

type Integral = integral.Integral
type Strategy = strategy.Strategy

type Estimator interface {
	Estimate(a float64, b float64, h float64, itg Integral, stg Strategy) (float64)
	String() string
}