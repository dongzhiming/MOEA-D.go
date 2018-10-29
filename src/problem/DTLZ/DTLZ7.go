package DTLZ

import (
	"math"
	"problem"
	"solution"
	"util"
)

type DTLZ7 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *DTLZ7) NewDTLZ7(nd, no, nc int) {
	this.name = "DTLZ7"
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

func (this *DTLZ7) GetName() string {
	return this.name
}

func (this *DTLZ7) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *DTLZ7) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *DTLZ7) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *DTLZ7) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}
	//	fmt.Println(solu)
	return solu
}

func (this *DTLZ7) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	var f []float64 = make([]float64, this.numberOfObjectives, this.numberOfObjectives)

	k := this.numberOfVariables - this.numberOfObjectives + 1

	g := 0.0

	for i := this.numberOfVariables - k; i < this.numberOfVariables; i++ {
		g += x[i]
	}

	g = 1.0 + (9.0*g)/float64(k)

	h := float64(this.numberOfObjectives)
	for i := 0; i < this.numberOfObjectives-1; i++ {
		h -= x[i] / (1.0 + g) * (1.0 + math.Sin(3.0*math.Pi*x[i]))
	}

	for i := 0; i < this.numberOfObjectives-1; i++ {
		f[i] = x[i]
	}
	f[this.numberOfObjectives-1] = (1.0 + g) * h

	for i := 0; i < this.numberOfObjectives; i++ {
		solution.SetObjective(i, f[i])
	}
}
