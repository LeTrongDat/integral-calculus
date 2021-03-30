package integral 

type Integral interface {
	F(x float64) (float64)
	GetDiscreatePts() ([]float64)
	String() string
}