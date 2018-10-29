package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

type UF7 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF7) NewUF7(nd, no, nc int) {
	this.name = "UF7"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd, nd)
	this.lower = make([]float64, nd, nd)

	this.lower[0] = 0
	this.upper[0] = 1
	for i := 1; i < nd; i++ {
		this.lower[i] = -1
		this.upper[i] = 1
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *UF7) GetName() string {
	return this.name
}

func (this *UF7) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF7) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF7) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF7) CreateSolution() *solution.Solution {
	solution := new(solution.Solution)
	solution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solution.SetVariableValue(i, value)
	}

	return solution
}

func (this *UF7) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	f := []float64{0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	sum1 := 0.0
	sum2 := 0.0
	var yj float64

	for j := 2; j <= nx; j++ {
		yj = x[j-1] - math.Sin(6.0*PI*x[0]+float64(j)*PI/float64(nx))

		if j%2 == 0 {
			sum2 += yj * yj
			count2++
		} else {
			sum1 += yj * yj
			count1++
		}
	}

	yj = math.Pow(x[0], 0.2)
	f[0] = yj + 2.0*sum1/float64(count1)
	f[1] = 1.0 - yj + 2.0*sum2/float64(count2)

	solution.SetObjective(0, f[0])
	solution.SetObjective(1, f[1])
}
