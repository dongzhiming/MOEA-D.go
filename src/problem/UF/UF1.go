package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

const PI float64 = 3.1415926535897932384626433832795

type UF1 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF1) NewUF1(nd, no, nc int) {
	this.name = "UF1"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd)
	this.lower = make([]float64, nd)

	this.lower[0] = 0
	this.upper[0] = 1
	for i := 1; i < nd; i++ {
		this.lower[i] = -1
		this.upper[i] = 1
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *UF1) GetName() string {
	return this.name
}

func (this *UF1) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF1) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF1) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF1) CreateSolution() *solution.Solution {
	solution := new(solution.Solution)
	solution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solution.SetVariableValue(i, value)
	}

	return solution
}

func (this *UF1) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	y := []float64{0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	sum1 := 0.0
	sum2 := 0.0
	var yj float64

	for j := 2; j <= nx; j++ {
		yj = x[j-1] - math.Sin(6.0*math.Pi*x[0]+float64(j)*math.Pi/float64(nx))
		yj = yj * yj

		if j%2 == 0 {
			sum2 += yj
			count2++
		} else {
			sum1 += yj
			count1++
		}
	}

	y[0] = x[0] + 2.0*sum1/float64(count1)
	y[1] = 1.0 - math.Sqrt(x[0]) + 2.0*sum2/float64(count2)

	solution.SetObjective(0, y[0])
	solution.SetObjective(1, y[1])
}
