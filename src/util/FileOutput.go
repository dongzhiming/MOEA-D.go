package util

import (
	"fmt"
	"global"
	"os"
	"problem"
	"solution"
	"strconv"
)

func SolutionListOutput(population []*solution.Solution, problem problem.Problem) {
	varPath := "src/result/VAR/"
	varFileOutput(population, varPath+problem.GetName()+"_"+strconv.Itoa(global.Running)+".VAR")
	funPath := "src/result/FUN/"
	funFileOutput(population, funPath+problem.GetName()+"_"+strconv.Itoa(global.Running)+".FUN")
}

func varFileOutput(population []*solution.Solution, userFile string) {
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	for i := 0; i < len(population); i++ {
		for j := 0; j < population[i].GetNumberOfVariables(); j++ {
			fout.WriteString(fmt.Sprintf("%16.12f", population[i].GetVariableValue(j)))

			fout.WriteString("  ")
		}
		fout.WriteString("\n")
	}

	fout.Close()
}

func funFileOutput(population []*solution.Solution, userFile string) {
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	for i := 0; i < len(population); i++ {
		for j := 0; j < population[i].GetNumberOfObjectives(); j++ {
			fout.WriteString(fmt.Sprintf("%16.12E", population[i].GetObjective(j)))

			fout.WriteString("  ")
		}
		fout.WriteString("\n")
	}

	fout.Close()
}

func fileOutput(userFile string) {
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
}
