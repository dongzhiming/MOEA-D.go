package DTLZ

import (
	"math"
	"problem"
	"solution"
	"util"
)

type DTLZ1 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *DTLZ1) NewDTLZ1(nd, no, nc int) {
	this.name = "DTLZ1"
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

func (this *DTLZ1) GetName() string {
	return this.name
}

func (this *DTLZ1) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *DTLZ1) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *DTLZ1) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *DTLZ1) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}
	//	fmt.Println(solu)
	return solu
}

func (this *DTLZ1) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	var f []float64 = make([]float64, this.numberOfObjectives, this.numberOfObjectives)

	k := this.numberOfVariables - this.numberOfObjectives + 1

	g := 0.0
	for i := this.numberOfVariables - k; i < this.numberOfVariables; i++ {
		g += math.Pow(x[i]-0.5, 2.0) - math.Cos(20.0*math.Pi*(x[i]-0.5))
	}
	g = 100.0 * (float64(k) + g)

	for i := 0; i < this.numberOfObjectives; i++ {
		f[i] = 0.5 * (1.0 + g)

		for j := 0; j < this.numberOfObjectives-i-1; j++ {
			f[i] *= x[j]
		}

		if i != 0 {
			f[i] *= 1 - x[this.numberOfObjectives-i-1]
		}
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		solution.SetObjective(i, f[i])
	}
}
