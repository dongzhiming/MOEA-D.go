package main

import (
	"fmt"
	"global"
	"moead"
	"time"
	"util"
)

func main() {
	testInstances := moead.GetTestInstances(4)

	for t := 0; t < len(testInstances); t++ {
		for i := 0; i < 30; i++ {
			global.Running = i

			tBefore := time.Now()
			algor := new(moead.MOEADCCUF)
			problem := moead.CreateProblem(testInstances[t])
			algor.NewMOEADCCUF(problem)
			algor.Run()
			tAfter := time.Now()
			fmt.Println("Total execution time:", (tAfter.Minute()-tBefore.Minute())*60+tAfter.Second()-tBefore.Second())
			util.SolutionListOutput(algor.GetResult(), problem)
		}
	}

}
