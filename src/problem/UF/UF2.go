package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

type UF2 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF2) NewUF2(nd, no, nc int) {
	this.name = "UF2"
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

func (this *UF2) GetName() string {
	return this.name
}

func (this *UF2) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF2) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF2) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF2) CreateSolution() *solution.Solution {
	solu := new(solution.Solution)
	solu.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solu.SetVariableValue(i, value)
	}
	//	fmt.Println(solu)
	return solu
}

func (this *UF2) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	f := []float64{0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	sum1 := 0.0
	sum2 := 0.0
	var yj float64

	for j := 2; j <= nx; j++ {
		if j%2 == 0 {
			yj = x[j-1] - 0.3*x[0]*(x[0]*math.Cos(24.0*PI*x[0]+4.0*float64(j)*PI/float64(nx))+2.0)*math.Sin(6.0*PI*x[0]+float64(j)*PI/float64(nx))
			sum2 += yj * yj
			count2++
		} else {
			yj = x[j-1] - 0.3*x[0]*(x[0]*math.Cos(24.0*PI*x[0]+4.0*float64(j)*PI/float64(nx))+2.0)*math.Cos(6.0*PI*x[0]+float64(j)*PI/float64(nx))
			sum1 += yj * yj
			count1++
		}
	}

	f[0] = x[0] + 2.0*sum1/float64(count1)
	f[1] = 1.0 - math.Sqrt(x[0]) + 2.0*sum2/float64(count2)

	solution.SetObjective(0, f[0])
	solution.SetObjective(1, f[1])
}
