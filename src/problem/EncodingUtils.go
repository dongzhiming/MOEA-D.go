package problem

import (
	"solution"
)

func GetReal(solution *solution.Solution) []float64 {
	var x []float64 = make([]float64, solution.GetNumberOfVariables(), solution.GetNumberOfVariables())

	for i := 0; i < solution.GetNumberOfVariables(); i++ {
		x[i] = solution.GetVariableValue(i)
	}

	return x
}
