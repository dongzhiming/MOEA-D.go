package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

type UF9 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF9) NewUF9(nd, no, nc int) {
	this.name = "UF9"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd, nd)
	this.lower = make([]float64, nd, nd)

	this.lower[0] = 0
	this.upper[0] = 1
	this.lower[1] = 0
	this.upper[1] = 1
	for i := 2; i < nd; i++ {
		this.lower[i] = -2
		this.upper[i] = 2
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *UF9) GetName() string {
	return this.name
}

func (this *UF9) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF9) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF9) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF9) CreateSolution() *solution.Solution {
	solution := new(solution.Solution)
	solution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solution.SetVariableValue(i, value)
	}

	return solution
}

func (this *UF9) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	f := []float64{0, 0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	count3 := 0
	sum1 := 0.0
	sum2 := 0.0
	sum3 := 0.0
	var yj float64
	E := 0.1

	for j := 3; j <= nx; j++ {
		yj = x[j-1] - 2.0*x[1]*math.Sin(2.0*PI*x[0]+float64(j)*PI/float64(nx))

		if j%3 == 1 {
			sum1 += yj * yj
			count1++
		} else if j%3 == 2 {
			sum2 += yj * yj
			count2++
		} else {
			sum3 += yj * yj
			count3++
		}
	}

	yj = (1.0 + E) * (1.0 - 4.0*(2.0*x[0]-1.0)*(2.0*x[0]-1.0))

	if yj < 0.0 {
		yj = 0.0
	}

	f[0] = 0.5*(yj+2*x[0])*x[1] + 2.0*sum1/float64(count1)
	f[1] = 0.5*(yj-2*x[0]+2.0)*x[1] + 2.0*sum2/float64(count2)
	f[2] = 1.0 - x[1] + 2.0*sum3/float64(count3)

	solution.SetObjective(0, f[0])
	solution.SetObjective(1, f[1])
	solution.SetObjective(2, f[2])
}
