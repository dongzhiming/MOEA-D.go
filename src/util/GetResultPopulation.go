package util

import (
	"math"
	"problem"
	"solution"
)

func GetResultPopulation(problem problem.Problem, population []*solution.Solution) []*solution.Solution {
	var resultPopulationSize int
	if problem.GetNumberOfObjectives() == 2 {
		resultPopulationSize = 100
	} else if problem.GetNumberOfObjectives() == 3 {
		resultPopulationSize = 150
	} else {
		resultPopulationSize = 100
	}

	return getResultPopulation(population, resultPopulationSize)
}

func getResultPopulation(population []*solution.Solution, resultPopulationSize int) []*solution.Solution {
	var resultPopulation []*solution.Solution

	// Step 1 Check for data number
	no := population[0].GetNumberOfObjectives()

	size := len(population)

	// Step 2 init information
	var domV []bool = make([]bool, size, size)
	var nnIndex []int = make([]int, size, size)

	for i := 0; i < size; i++ {
		domV[i] = false
		nnIndex[i] = 0
	}

	var dom int
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			dom = DominanceComparator(population[i], population[j])

			if dom > 0 {
				domV[i] = true
			} else if dom < 0 {
				domV[j] = true
			}
		}
	}

	// init the no-dominated solution size
	var nndom int = 0
	for i := 0; i < size; i++ {
		if !domV[i] {
			nnIndex[nndom] = i
			nndom++
		}
	}

	// Step 2.2: find the final solutions
	if nndom <= resultPopulationSize {
		/* set these no-dominated solution as the result population*/
		resultPopulation = make([]*solution.Solution, nndom, nndom)

		for i := 0; i < nndom; i++ {
			resultPopulation[i] = population[nnIndex[i]].Copy()
		}

	} else {
		/*select resultPopulation size of population*/
		resultPopulation = make([]*solution.Solution, resultPopulationSize, resultPopulationSize)

		var indS []int = make([]int, resultPopulationSize, resultPopulationSize)
		var dis2set []float64 = make([]float64, nndom, nndom)
		var selected []bool = make([]bool, nndom, nndom)
		for i := 0; i < nndom; i++ {
			selected[i] = false
			dis2set[i] = 1.0e100
		}

		/*compute distance between the nndom individuals*/
		var disV [][]float64 = make([][]float64, nndom, nndom)
		for i := 0; i < nndom; i++ {
			disV[i] = make([]float64, nndom, nndom)
		}

		for i := 0; i < nndom-1; i++ {
			for j := i + 1; j < nndom; j++ {
				sum := 0.0
				for k := 0; k < no; k++ {
					sum += math.Pow(population[nnIndex[i]].GetObjective(k)-population[nnIndex[j]].GetObjective(k), 2)
				}

				disV[i][j] = sum
				disV[j][i] = disV[i][j]
			}
		}

		/* select extreme points*/
		var fmin float64
		var index int
		for k := 0; k < no; k++ {
			fmin = 1.0e100
			index = 0
			for i := 0; i < nndom; i++ {
				if !selected[i] && population[nnIndex[i]].GetObjective(k) < fmin {
					fmin = population[nnIndex[i]].GetObjective(k)
					index = i
				}
			}

			indS[k] = nnIndex[index]
			selected[index] = true

			for i := 0; i < nndom; i++ {
				if !selected[i] && dis2set[i] > disV[i][index] {
					dis2set[i] = disV[i][index]
				}
			}
		}

		for k := no; k < resultPopulationSize; k++ {
			fmin := -1.0e100
			index := 0
			for i := 0; i < nndom; i++ {
				if !selected[i] && dis2set[i] > fmin {
					fmin = dis2set[i]
					index = i
				}
			}

			indS[k] = nnIndex[index]
			selected[index] = true

			for i := 0; i < nndom; i++ {
				if !selected[i] && dis2set[i] > disV[i][index] {
					dis2set[i] = disV[i][index]
				}
			}
		}

		for i := 0; i < resultPopulationSize; i++ {
			resultPopulation[i] = population[indS[i]].Copy()
		}
	}

	return resultPopulation
}
