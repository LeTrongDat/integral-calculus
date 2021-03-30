package strategy 

import (
	"itmo/integral"
)

type Integral = integral.Integral


type Strategy interface {
	Execute(a float64, b float64, h float64, itg Integral) (float64)
	String() string
}