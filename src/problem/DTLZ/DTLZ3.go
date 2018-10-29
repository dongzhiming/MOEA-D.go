package DTLZ

import (
	"math"
	"problem"
	"solution"
	"util"
)

type DTLZ3 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *DTLZ3) NewDTLZ3(nd, no, nc int) {
	this.name = "DTLZ3"
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

func (this *DTLZ3) GetName() string {
	return this.name
}

func (this *DTLZ3) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *DTLZ3) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *DTLZ3) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *DTLZ3) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}
	//	fmt.Println(solu)
	return solu
}

func (this *DTLZ3) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	var f []float64 = make([]float64, this.numberOfObjectives, this.numberOfObjectives)

	k := this.numberOfVariables - this.numberOfObjectives + 1

	g := 0.0
	for i := this.numberOfVariables - k; i < this.numberOfVariables; i++ {
		g += (x[i]-0.5)*(x[i]-0.5) - math.Cos(20.0*math.Pi*(x[i]-0.5))
	}

	g = 100.0 * (float64(k) + g)
	for i := 0; i < this.numberOfObjectives; i++ {
		f[i] = 1.0 + g
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		for j := 0; j < this.numberOfObjectives-(i+1); j++ {
			f[i] *= math.Cos(x[j] * 0.5 * math.Pi)
		}
		if i != 0 {
			aux := this.numberOfObjectives - (i + 1)
			f[i] *= math.Sin(x[aux] * 0.5 * math.Pi)
		}
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		solution.SetObjective(i, f[i])
	}
}
