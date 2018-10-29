package DTLZ

import (
	"math"
	"problem"
	"solution"
	"util"
)

type DTLZ6 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *DTLZ6) NewDTLZ6(nd, no, nc int) {
	this.name = "DTLZ6"
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

func (this *DTLZ6) GetName() string {
	return this.name
}

func (this *DTLZ6) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *DTLZ6) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *DTLZ6) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *DTLZ6) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}
	//	fmt.Println(solu)
	return solu
}

func (this *DTLZ6) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	var f []float64 = make([]float64, this.numberOfObjectives, this.numberOfObjectives)
	var theta []float64 = make([]float64, this.numberOfObjectives-1, this.numberOfObjectives-1)

	k := this.numberOfVariables - this.numberOfObjectives + 1

	g := 0.0

	for i := this.numberOfVariables - k; i < this.numberOfVariables; i++ {
		g += math.Pow(x[i], 0.1)
	}

	t := math.Pi / (4.0 * (1.0 + g))

	theta[0] = x[0] * math.Pi / 2.0
	for i := 1; i < (this.numberOfObjectives - 1); i++ {
		theta[i] = t * (1.0 + 2.0*g*x[i])
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		f[i] = 1.0 + g
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		for j := 0; j < this.numberOfObjectives-(i+1); j++ {
			f[i] *= math.Cos(theta[j])
		}
		if i != 0 {
			aux := this.numberOfObjectives - (i + 1)
			f[i] *= math.Sin(theta[aux])
		}
	}

	for i := 0; i < this.numberOfObjectives; i++ {
		solution.SetObjective(i, f[i])
	}
}
