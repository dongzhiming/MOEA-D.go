package ZDT

import (
	"math"
	"problem"
	"solution"
	"util"
)

type ZDT4 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *ZDT4) NewZDT4(nd, no, nc int) {
	this.name = "ZDT4"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd, nd)
	this.lower = make([]float64, nd, nd)

	for i := 0; i < nd; i++ {
		this.upper[i] = 1
		this.lower[i] = 0
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *ZDT4) GetName() string {
	return this.name
}

func (this *ZDT4) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *ZDT4) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *ZDT4) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *ZDT4) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}

	return solu
}

func (this *ZDT4) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)

	g := 0.0
	for i := 1; i < len(x); i++ {
		g += math.Pow(x[i], 2.0) - 10.0*math.Cos(4.0*math.Pi*x[i])
	}
	g += 1.0 + 10.0*float64(this.numberOfVariables-1)

	h := 1.0 - math.Sqrt(x[0]/g)

	solution.SetObjective(0, x[0])
	solution.SetObjective(1, g*h)
}
