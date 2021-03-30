package invoker 

import (
	"fmt"
	"math"
	"errors"
	"log"
	"os"

	"itmo/integral"
	"itmo/solver"
	"itmo/strategy"
	"itmo/estimator"
)

type Integral = integral.Integral
type Strategy = strategy.Strategy
type Estimator = estimator.Estimator
type Solver = solver.Solver


type InvokerImpl struct {
	integrals []Integral
	strategies []Strategy 
	estimators []Estimator
	chosen_integral int 
	chosen_strategy int 
	chosen_estimator int
}

func (ivk *InvokerImpl) Invoke() {
	fmt.Println("> Welcome to the integral world!")

	ivk.initIntegrals()
	ivk.initStrategies()
	ivk.initEstimators()

	for ;; {
		err := ivk.chooseIntegral()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		err = ivk.chooseStrategy()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		err = ivk.chooseEstimator()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		solver := solver.NewSolver(ivk.integrals[ivk.chosen_integral], ivk.strategies[ivk.chosen_strategy], ivk.estimators[ivk.chosen_estimator])
		err = solver.Solve()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("> Do you want to continue? (Y/N): ")
		var exit string 
		fmt.Scanf("%s", &exit)
		if exit == "N" {
			os.Exit(0)
		}
	}
}

func (ivk *InvokerImpl) initIntegrals() {
	ivk.integrals = []Integral{
		integral.NewIntegral("3x^2 + 2", func (x float64) (float64) { return 3 * math.Pow(x, 2) + 2 }),
		integral.NewIntegral("sin(x) + 4x^3 - x", func (x float64) (float64) { return math.Sin(x) + 4 * math.Pow(x, 3) - x }),
		integral.NewIntegral("e^x + cos(x) - x^3", func (x float64) (float64) { return math.Exp(x) + math.Cos(x) - math.Pow(x, 3)}),
		integral.NewIntegral("ln(x) - e^x", func (x float64) (float64) { return math.Log(x) - math.Exp(x) }, 0),
		integral.NewIntegral("1/x", func (x float64) (float64) { return 1./x }, 0),
	}
}
func (ivk *InvokerImpl) chooseIntegral() (error) {
	fmt.Printf("-------------------- INTEGRAL --------------------\n")
	fmt.Println("> Which integral do you want to choose?")
	for i, option := range ivk.integrals {
		fmt.Printf("> %d. %s\n", i, option.String())
	}
	fmt.Printf("> Enter your variant: ")
	fmt.Scanf("%d", &ivk.chosen_integral)
	if ivk.chosen_integral >= len(ivk.integrals) || ivk.chosen_integral < 0 {
		return errors.New("Invalid option")
	}
	return nil
}

func (ivk *InvokerImpl) initStrategies() {
	ivk.strategies = []Strategy {
		strategy.NewLeftRectangleStrategy(),
		strategy.NewRightRectangleStrategy(),
		strategy.NewMiddleRectangleStrategy(),
	}
}
func (ivk *InvokerImpl) chooseStrategy() (error) {
	fmt.Printf("-------------------- STRATEGY --------------------\n")
	fmt.Println("> Which strategy do you want to choose?")
	for i, option := range ivk.strategies {
		fmt.Printf("> %d. %s\n", i, option.String())
	}
	fmt.Printf("> Enter your variant: ")
	fmt.Scanf("%d", &ivk.chosen_strategy)
	if ivk.chosen_strategy >= len(ivk.strategies) || ivk.chosen_strategy < 0 {
		return errors.New("Invalid option")
	}
	return nil
}

func (ivk *InvokerImpl) initEstimators() {
	ivk.estimators = []Estimator {
		estimator.NewRungeEstimator(),
	}
}
func (ivk *InvokerImpl) chooseEstimator() (error) {
	fmt.Printf("-------------------- ESTIMATOR --------------------\n")
	fmt.Println("> Which estimator do you want to choose?")
	for i, option := range ivk.estimators {
		fmt.Printf("> %d. %s\n", i, option.String())
	}
	fmt.Printf("> Enter your variant: ")
	fmt.Scanf("%d", &ivk.chosen_estimator)
	if ivk.chosen_estimator >= len(ivk.estimators) || ivk.chosen_estimator < 0 {
		return errors.New("Invalid option")
	}
	return nil
}