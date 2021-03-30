package solver 

import (
	"math"
	"fmt"
	"errors"
	"itmo/integral"
	"itmo/strategy"
	"itmo/estimator"
)

type Integral = integral.Integral
type Strategy = strategy.Strategy
type Estimator = estimator.Estimator

type SolverImpl struct {
	Itg Integral
	Stg Strategy
	Est Estimator	
}

func (solver *SolverImpl) Solve() (error) {
	var a, b, eps float64
	fmt.Printf("> Enter your parameters (a, b, epsilon): ")
	fmt.Scanf("%f %f %f", &a, &b, &eps)
	if solver.Itg.GetDiscreatePts() != nil {
		for _, pt := range solver.Itg.GetDiscreatePts() {
			if (a <= pt && pt <= b) || (a >= pt && pt >= b) {
				return errors.New("The interval is not continuos")
			}
		}
	}
	h := math.Pow(eps, 1./4.)
	cnt_split := 0
	for ; solver.Est.Estimate(a, b, h, solver.Itg, solver.Stg) >= eps; {
		cnt_split++
		h /= 2
	}
	fmt.Printf("-------------------- Result --------------------\n")
	fmt.Printf("> Result: %f\n", solver.Stg.Execute(a, b, h, solver.Itg))
	fmt.Printf("> Number of iterations: %d\n", int64(math.Ceil(math.Abs(b-a)/h)))
	fmt.Printf("> Number of splits: %d\n", cnt_split)
	fmt.Printf("> Error: %f\n", solver.Est.Estimate(a, b, h, solver.Itg, solver.Stg))
	return nil
}

func NewSolver(itg Integral, stg Strategy, est Estimator) (*SolverImpl) {
	solver := &SolverImpl{itg, stg, est}
	return solver
}