package integral 

type IntegralImpl struct {
	Fx func(float64)(float64) 
	Exp string
	DiscreatePts []float64
}

func (itg *IntegralImpl) F(x float64) (float64) {
	return itg.Fx(x)
}

func (itg *IntegralImpl) String() string {
	return itg.Exp
}

func (itg *IntegralImpl) GetDiscreatePts() ([]float64) {
	return itg.DiscreatePts
}

func NewIntegral(expression string, fx func(float64)(float64), dicrete_pts ...float64) (*IntegralImpl) {
	itg := &IntegralImpl{fx, expression, dicrete_pts}
	return itg
}