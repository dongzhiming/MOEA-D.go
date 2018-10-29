package ZDT

import (
	"math"
	"problem"
	"solution"
	"util"
)

type ZDT3 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *ZDT3) NewZDT3(nd, no, nc int) {
	this.name = "ZDT3"
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

func (this *ZDT3) GetName() string {
	return this.name
}

func (this *ZDT3) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *ZDT3) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *ZDT3) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *ZDT3) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}

	return solu
}

func (this *ZDT3) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)

	g := 0.0
	for i := 1; i < len(x); i++ {
		g += x[i]
	}
	g = (9.0/(float64(this.numberOfVariables-1)))*g + 1.0

	h := 1.0 - math.Sqrt(x[0]/g) - (x[0]/g)*math.Sin(10.0*math.Pi*x[0])

	solution.SetObjective(0, x[0])
	solution.SetObjective(1, g*h)
}
