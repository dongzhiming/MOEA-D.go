package ZDT

import (
	"math"
	"problem"
	"solution"
	"util"
)

type ZDT6 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *ZDT6) NewZDT6(nd, no, nc int) {
	this.name = "ZDT6"
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

func (this *ZDT6) GetName() string {
	return this.name
}

func (this *ZDT6) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *ZDT6) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *ZDT6) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *ZDT6) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}

	return solu
}

func (this *ZDT6) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)

	f := 1.0 - math.Exp(-4.0*x[0])*math.Pow(math.Sin(6.0*math.Pi*x[0]), 6.0)

	g := 0.0
	for i := 1; i < len(x); i++ {
		g += x[i]
	}
	g = 1.0 + 9.0*math.Pow(g/float64(this.numberOfVariables-1), 0.25)

	h := 1.0 - math.Pow(f/g, 2.0)

	solution.SetObjective(0, f)
	solution.SetObjective(1, g*h)
}
