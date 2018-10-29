package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

type UF3 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF3) NewUF3(nd, no, nc int) {
	this.name = "UF3"
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

func (this *UF3) GetName() string {
	return this.name
}

func (this *UF3) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF3) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF3) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF3) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}

	return solu
}

func (this *UF3) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	f := []float64{0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	sum1 := 0.0
	sum2 := 0.0
	prod1 := 1.0
	prod2 := 1.0
	var yj, pj float64

	for j := 2; j <= nx; j++ {
		yj = x[j-1] - math.Pow(x[0], 0.5*(1.0+3.0*(float64(j)-2.0)/(float64(nx)-2.0)))
		pj = math.Cos(20.0 * yj * PI / math.Sqrt(float64(j)))

		if j%2 == 0 {
			sum2 += yj * yj
			prod2 *= pj
			count2++
		} else {
			sum1 += yj * yj
			prod1 *= pj
			count1++
		}
	}

	f[0] = x[0] + 2.0*(4.0*sum1-2.0*prod1+2.0)/float64(count1)
	f[1] = 1.0 - math.Sqrt(x[0]) + 2.0*(4.0*sum2-2.0*prod2+2.0)/float64(count2)

	solution.SetObjective(0, f[0])
	solution.SetObjective(1, f[1])
}
