package LZ09

import (
	"solution"
	"util"
)

type LZ09_F4 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *LZ09_F4) NewLZ09_F4(nd, no, nc int) {
	this.name = "LZ09_F4"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd, nd)
	this.lower = make([]float64, nd, nd)

	for i := 0; i < nd; i++ {
		this.lower[i] = 0
		this.upper[i] = 1
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *LZ09_F4) GetName() string {
	return this.name
}

func (this *LZ09_F4) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *LZ09_F4) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *LZ09_F4) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *LZ09_F4) CreateSolution() *solution.Solution {
	solution := new(solution.Solution)
	solution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solution.SetVariableValue(i, value)
	}

	return solution
}

func (this *LZ09_F4) Evaluate(solution *solution.Solution) {
	lz09(30, 2, 21, 24, 1)
	evaluate(solution)
}
