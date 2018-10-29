package problem

import (
	"solution"
)

type Problem interface {
	GetName() string
	GetNumberOfObjectives() int
	GetNumberOfVariables() int
	GetNumberOfConstraints() int
	CreateSolution() *solution.Solution
	Evaluate(solu *solution.Solution)
}
