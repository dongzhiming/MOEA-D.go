package util

import (
	"math"
	"problem"
	"solution"
	"strconv"
)

func GetInvertedGenerationalDistance(population []*solution.Solution, problem problem.Problem) float64 {
	var front, referenceFront [][]float64

	lenPopulation := len(population)

	front = make([][]float64, lenPopulation, lenPopulation)
	for i := 0; i < lenPopulation; i++ {
		front[i] = make([]float64, problem.GetNumberOfObjectives(), problem.GetNumberOfObjectives())
		for j := 0; j < population[i].GetNumberOfObjectives(); j++ {
			front[i][j] = population[i].GetObjective(j)
		}
	}

	// 读取真实paretofront
	referenceFront = make([][]float64, 0)

	var path string = "src/resources/pareto_fronts/"
	var dataFileName string = problem.GetName() + ".pf"

	if problem.GetName() == "DTLZ1" || problem.GetName() == "DTLZ2" || problem.GetName() == "DTLZ3" || problem.GetName() == "DTLZ4" || problem.GetName() == "DTLZ5" || problem.GetName() == "DTLZ6" || problem.GetName() == "DTLZ7" {
		dataFileName = problem.GetName() + "." + strconv.Itoa(problem.GetNumberOfObjectives()) + "D.pf"
	}

	referenceFront = ReadDoubleDataFile(path + dataFileName)

	//
	return getInvertedGenerationalDistance(front, referenceFront)
}

func getInvertedGenerationalDistance(front, referenceFront [][]float64) float64 {
	var sum float64 = 0.0

	size := len(referenceFront)

	for i := 0; i < size; i++ {
		sum += distanceToClosestPoint(referenceFront[i], front)
	}

	return sum / float64(size)
}

func distanceToClosestPoint(point []float64, front [][]float64) float64 {
	size := len(front)
	minDistance := compute(point, front[0])

	for i := 1; i < size; i++ {
		aux := compute(point, front[i])
		if aux < minDistance {
			minDistance = aux
		}
	}

	return minDistance
}

func compute(point1, point2 []float64) float64 {
	size := len(point1)

	sum := 0.0

	for i := 0; i < size; i++ {
		sum += math.Pow(point2[i]-point1[i], 2)
	}

	return math.Sqrt(sum)
}
