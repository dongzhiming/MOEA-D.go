package util

import (
	"solution"
)

func DominanceComparator(solution1, solution2 *solution.Solution) int {
	bestIsOne := 0
	bestIsTwo := 0
	var result int

	for i := 0; i < solution1.GetNumberOfObjectives(); i++ {
		value1 := solution1.GetObjective(i)
		value2 := solution2.GetObjective(i)
		if value1 != value2 {
			if value1 < value2 {
				bestIsOne = 1
			}
			if value2 < value1 {
				bestIsTwo = 1
			}
		}
	}

	if bestIsOne > bestIsTwo {
		result = -1
	} else if bestIsTwo > bestIsOne {
		result = 1
	} else {
		result = 0
	}

	return result
}
