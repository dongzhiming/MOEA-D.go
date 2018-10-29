package moead

import (
	"fmt"
	"global"
	"math"
	"problem"
	"solution"
	"strconv"
	"util"
)

type MOEAD struct {
	maxEvaluations                   int
	evaluations                      int
	populationSize                   int
	algorithmName                    string
	population                       []*solution.Solution
	problem                          problem.Problem
	idealPoint                       []float64
	lambda                           [][]float64
	neighborSize                     int
	neighborhood                     [][]int
	maximumNumberOfReplacedSolutions int
	neighborhoodSelectionProbability float64
	resultPopulationSize             int
	mutationProbability              float64
	mutationDistributionIndex        float64
	rand                             *util.Random
	functionType                     FunctionType
}

func (this *MOEAD) NewMOEAD(problem problem.Problem) {
	this.problem = problem
	this.algorithmName = "MOEADDE"
	this.maxEvaluations = 300000
	this.populationSize = 300
	if problem.GetNumberOfObjectives() == 3 {
		this.populationSize = 600
	}
	this.neighborSize = 20
	this.maximumNumberOfReplacedSolutions = 2
	this.neighborhoodSelectionProbability = 0.9
	this.mutationProbability = 1.0 / float64(problem.GetNumberOfVariables())
	this.mutationDistributionIndex = 20.0

	this.rand = new(util.Random)
	this.rand.NewRand()

	this.functionType = TCHE
}

func (this *MOEAD) Run() {
	fmt.Println("MOEADDE is running", global.Running, "...")

	this.initializePopulation()
	this.initializeUniformWeight()
	this.initializeNeighborhood()
	this.initializeIdealPoint()

	this.evaluations = this.populationSize

	for !this.isStoppingConditionReached() {

		perm := this.rand.Perm(this.populationSize)

		for i := 0; i < this.populationSize; i++ {

			id := perm[i]
			//			id := i

			neighborType := this.chooseNeighborType()

			// 选择
			//			var parents []*solution.Solution = make([]*solution.Solution, 3, 3)
			parents := this.parentsSelection(id, neighborType)

			// 交叉
			child := this.differentialEvolutionCrossover(parents)
			//			fmt.Println(child)

			// 变异
			this.mutation(child)

			this.problem.Evaluate(child)
			//			fmt.Println(child.GetObjective(0), child.GetObjective(1))

			this.updateIdealPoint(child)

			this.updatePopulation(child, id, neighborType)

			this.evaluations++
		}
	}

	fmt.Println(this.problem.GetName(), "IGD value =", util.GetInvertedGenerationalDistance(this.GetResult(), this.problem))
}

func (this *MOEAD) GetResult() []*solution.Solution {
	return util.GetResultPopulation(this.problem, this.population)
}

func (this *MOEAD) updatePopulation(child *solution.Solution, id int, nType NeighborType) {
	size, time := 0, 0

	if nType == NEIGHBOR {
		size = this.neighborSize
	} else {
		size = this.populationSize
	}

	perm := this.rand.Perm(size)

	for i := 0; i < size; i++ {
		var k int
		if nType == NEIGHBOR {
			k = this.neighborhood[id][perm[i]]
		} else {
			k = perm[i]
		}

		f1 := this.fitnessFunction(this.population[k], this.lambda[k])
		f2 := this.fitnessFunction(child, this.lambda[k])

		if f2 < f1 {
			//			fmt.Println("child", child)
			this.population[k] = child.Copy()
			//			this.population[k].Replace(child)
			//			fmt.Println("population[k]", this.population[k])
			time++
		}

		if time >= this.maximumNumberOfReplacedSolutions {
			return
		}
	}
}

func (this *MOEAD) fitnessFunction(child *solution.Solution, lamb []float64) float64 {
	var fitness float64

	if this.functionType == TCHE {
		var maxFun float64 = -1.0e+30

		for n := 0; n < this.problem.GetNumberOfObjectives(); n++ {
			diff := math.Abs(child.GetObjective(n) - this.idealPoint[n])

			var feval float64
			if lamb[n] == 0 {
				feval = 0.0001 * diff
			} else {
				feval = diff * lamb[n]
			}

			if feval > maxFun {
				maxFun = feval
			}
		}

		fitness = maxFun
	} else if this.functionType == EWC {
		p := 100.0
		sum := 0.0
		for i := 0; i < this.problem.GetNumberOfObjectives(); i++ {
			sum += math.Exp(p*lamb[i]-1) * math.Exp(p*(child.GetObjective(i)-this.idealPoint[i]))
		}

		fitness = sum
	} else if this.functionType == WCP {
		p := 9.0
		sum := 0.0
		for i := 0; i < this.problem.GetNumberOfObjectives(); i++ {
			sum += math.Pow(lamb[i]*(child.GetObjective(i)-this.idealPoint[i]), p)
		}

		fitness = sum
	}

	return fitness
}

func (this *MOEAD) differentialEvolutionCrossover(parents []*solution.Solution) *solution.Solution {
	child := parents[2].Copy()
	jrand := this.rand.Intn(child.GetNumberOfVariables())

	cr := 1.0
	f := 0.5

	for j := 0; j < child.GetNumberOfVariables(); j++ {
		if this.rand.Float64() < cr || j == jrand {
			value := parents[2].GetVariableValue(j) + f*(parents[0].GetVariableValue(j)-parents[1].GetVariableValue(j))
			if value < child.GetLowerBound(j) {
				value = child.GetLowerBound(j)
			}

			if value > child.GetUpperBound(j) {
				value = child.GetLowerBound(j)
			}

			child.SetVariableValue(j, value)
		} else {
			value := parents[2].GetVariableValue(j)
			child.SetVariableValue(j, value)
		}
	}

	return child
}

func (this *MOEAD) mutation(child *solution.Solution) {
	var rnd, delta1, delta2, mutPow, deltaq float64
	var y, yl, yu, val, xy float64

	for i := 0; i < this.problem.GetNumberOfVariables(); i++ {
		if this.rand.Float64() <= this.mutationProbability {
			y = child.GetVariableValue(i)
			yl = child.GetLowerBound(i)
			yu = child.GetUpperBound(i)

			if yl == yu {
				y = yl
			} else {
				delta1 = (y - yl) / (yu - yl)
				delta2 = (yu - y) / (yu - yl)
				rnd = this.rand.Float64()
				mutPow = 1.0 / (this.mutationDistributionIndex + 1.0)
				if rnd <= 0.5 {
					xy = 1.0 - delta1
					val = 2.0*rnd + (1.0-2.0*rnd)*(math.Pow(xy, this.mutationDistributionIndex+1.0))
					deltaq = math.Pow(val, mutPow) - 1.0
				} else {
					xy = 1.0 - delta2
					val = 2.0*(1.0-rnd) + 2.0*(rnd-0.5)*(math.Pow(xy, this.mutationDistributionIndex+1.0))
					deltaq = 1.0 - math.Pow(val, mutPow)
				}
				y = y + deltaq*(yu-yl)

				if y < yl {
					y = yl
				}

				if y > yu {
					y = yu
				}
			}

			child.SetVariableValue(i, y)
		}
	}
}

func (this *MOEAD) parentsSelection(id int, neighborType NeighborType) []*solution.Solution {
	matingPool := this.matingSelection(id, 2, neighborType)

	parents := make([]*solution.Solution, 3)

	parents[0] = this.population[matingPool[0]]
	parents[1] = this.population[matingPool[1]]
	parents[2] = this.population[id]

	return parents
}

func (this *MOEAD) matingSelection(id, size int, neighborType NeighborType) []int {
	var selectedSolution int
	listOfSolutions := make([]int, 0, size)

	for len(listOfSolutions) < size {
		if neighborType == NEIGHBOR {
			random := this.rand.Intn(this.neighborSize)
			selectedSolution = this.neighborhood[id][random]
		} else {
			selectedSolution = this.rand.Intn(this.populationSize)
		}

		flag := true
		for _, item := range listOfSolutions {
			if item == selectedSolution {
				flag = false
				break
			}
		}

		if flag {
			listOfSolutions = append(listOfSolutions, selectedSolution)
		}
	}

	return listOfSolutions
}

func (this *MOEAD) chooseNeighborType() NeighborType {
	var neighborType NeighborType
	rnd := this.rand.Float64()

	if rnd < this.neighborhoodSelectionProbability {
		neighborType = NEIGHBOR
	} else {
		neighborType = POPULATION
	}

	return neighborType
}

func (this *MOEAD) isStoppingConditionReached() bool {
	return this.evaluations >= this.maxEvaluations
}

func (this *MOEAD) initializePopulation() {
	this.population = make([]*solution.Solution, this.populationSize, this.populationSize)

	for i := 0; i < this.populationSize; i++ {
		newSolution := this.problem.CreateSolution()
		this.problem.Evaluate(newSolution)
		this.population[i] = newSolution
		//		fmt.Println("newSolution", newSolution)
	}

	//	for i := 0; i < this.populationSize; i++ {
	//		fmt.Println("this.population", this.population[i])
	//	}
}

func (this *MOEAD) initializeUniformWeight() {
	if this.problem.GetNumberOfObjectives() == 2 && this.populationSize <= 300 {
		this.lambda = make([][]float64, this.populationSize)
		no := this.problem.GetNumberOfObjectives()
		for i := 0; i < this.populationSize; i++ {
			this.lambda[i] = make([]float64, no)
		}

		for n := 0; n < this.populationSize; n++ {
			var a float64 = 1.0 * float64(n) / float64(this.populationSize-1)
			this.lambda[n][0] = a
			this.lambda[n][1] = 1 - a
		}
	} else {
		var path string = "src/resources/MOEAD_Weights/"
		var dataFileName string = "W" + strconv.Itoa(this.problem.GetNumberOfObjectives()) + "D_" + strconv.Itoa(this.populationSize) + ".dat"

		//		fmt.Println(path + dataFileName)

		this.lambda = util.ReadDoubleDataFile(path + dataFileName)

		//		var unitSize int = 33

		//		for i := 0; i <= unitSize; i++ {
		//			for j := 0; j <= unitSize; j++ {
		//				if i+j <= unitSize {
		//					sub := make([]float64, 3)
		//					sub = append(sub, float64(i))
		//					sub = append(sub, float64(j))
		//					sub = append(sub, float64(unitSize-i-j))
		//					for k := 0; k < this.problem.GetNumberOfObjectives(); k++ {
		//						sub[k] = sub[k] / float64(unitSize)
		//					}
		//					this.lambda = append(this.lambda, sub)
		//				}
		//			}
		//		}

		//		this.populationSize = len(this.lambda)
	}
}

func (this *MOEAD) initializeNeighborhood() {
	this.neighborhood = make([][]int, this.populationSize, this.populationSize)
	for i := 0; i < this.populationSize; i++ {
		this.neighborhood[i] = make([]int, this.neighborSize, this.neighborSize)
	}

	var x = make([]float64, this.populationSize, this.populationSize)
	var idx = make([]int, this.populationSize, this.populationSize)

	for i := 0; i < this.populationSize; i++ {
		for j := 0; j < this.populationSize; j++ {
			x[j] = this.distVector(this.lambda[i], this.lambda[j])
			idx[j] = j
		}

		this.minFastSort(x, idx, this.populationSize, this.neighborSize)

		for j := 0; j < this.neighborSize; j++ {
			this.neighborhood[i][j] = idx[j]
		}
	}

	//	fmt.Println(this.neighborhood)
}

func (this *MOEAD) distVector(lam1, lam2 []float64) float64 {
	dim := len(lam1)
	sum := 0.0
	for n := 0; n < dim; n++ {
		sum += (lam1[n] - lam2[n]) * (lam1[n] - lam2[n])
	}

	return math.Sqrt(sum)
}

func (this *MOEAD) minFastSort(x []float64, idx []int, popSize, neibSize int) {
	for i := 0; i < neibSize; i++ {
		for j := i + 1; j < popSize; j++ {
			if x[i] > x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
				id := idx[i]
				idx[i] = idx[j]
				idx[j] = id
			}
		}
	}
}

func (this *MOEAD) initializeIdealPoint() {
	this.idealPoint = make([]float64, this.problem.GetNumberOfObjectives(), this.problem.GetNumberOfObjectives())

	for i := 0; i < this.problem.GetNumberOfObjectives(); i++ {
		this.idealPoint[i] = 1.0e+30
	}

	for i := 0; i < this.populationSize; i++ {
		this.updateIdealPoint(this.population[i])
	}

	//	fmt.Println(this.idealPoint)
}

func (this *MOEAD) updateIdealPoint(solution *solution.Solution) {
	for i := 0; i < this.problem.GetNumberOfObjectives(); i++ {
		if solution.GetObjective(i) < this.idealPoint[i] {
			this.idealPoint[i] = solution.GetObjective(i)
		}
	}
}

func (this *MOEAD) GetName() string {
	return this.algorithmName
}

func (this *MOEAD) GetMaxEvaluations() int {
	return this.maxEvaluations
}

func (this *MOEAD) GetPopulationSize() int {
	return this.populationSize
}
