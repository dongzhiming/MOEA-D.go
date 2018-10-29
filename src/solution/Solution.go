package solution

type Solution struct {
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	objectives          []float64
	variables           []float64
	upper               []float64
	lower               []float64
}

func (this *Solution) NewSolution(nd, no, nc int, lower, upper []float64) {
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc

	this.lower = make([]float64, nd, nd)
	this.upper = make([]float64, nd, nd)
	this.objectives = make([]float64, no, no)
	this.variables = make([]float64, nd, nd)

	for i := 0; i < nd; i++ {
		this.lower[i] = lower[i]
		this.upper[i] = upper[i]
	}

	//	fmt.Println("This is NewSolution method ...")
}

func (this *Solution) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *Solution) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *Solution) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *Solution) GetLowerBound(id int) float64 {
	return this.lower[id]
}

func (this *Solution) GetUpperBound(id int) float64 {
	return this.upper[id]
}

func (this *Solution) Copy() *Solution {
	copySolution := new(Solution)

	copySolution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)

	for i := 0; i < this.numberOfObjectives; i++ {
		copySolution.SetObjective(i, this.GetObjective(i))
	}
	for i := 0; i < this.numberOfVariables; i++ {
		copySolution.SetVariableValue(i, this.GetVariableValue(i))
	}

	return copySolution
}

func (this *Solution) Replace(child *Solution) {
	for i := 0; i < this.numberOfObjectives; i++ {
		this.SetObjective(i, child.GetObjective(i))
	}

	for i := 0; i < this.numberOfVariables; i++ {
		this.SetVariableValue(i, child.GetVariableValue(i))
	}
}

func (this *Solution) GetVariableValue(id int) float64 {
	return this.variables[id]
}

func (this *Solution) GetObjective(id int) float64 {
	return this.objectives[id]
}

func (this *Solution) SetObjective(id int, value float64) {
	this.objectives[id] = value
}

func (this *Solution) SetVariableValue(id int, value float64) {
	this.variables[id] = value
}
