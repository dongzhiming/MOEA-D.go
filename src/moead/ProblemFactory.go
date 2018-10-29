package moead

import (
	"fmt"
	"problem"
	"problem/DTLZ"
	"problem/LZ09"
	"problem/UF"
	"problem/ZDT"
)

func CreateProblem(name string) problem.Problem {
	var pro problem.Problem

	if name == "UF1" {
		uf := new(UF.UF1)
		uf.NewUF1(30, 2, 0)
		pro = uf
	} else if name == "UF2" {
		uf := new(UF.UF2)
		uf.NewUF2(30, 2, 0)
		pro = uf
	} else if name == "UF3" {
		uf := new(UF.UF3)
		uf.NewUF3(30, 2, 0)
		pro = uf
	} else if name == "UF4" {
		uf := new(UF.UF4)
		uf.NewUF4(30, 2, 0)
		pro = uf
	} else if name == "UF5" {
		uf := new(UF.UF5)
		uf.NewUF5(30, 2, 0)
		pro = uf
	} else if name == "UF6" {
		uf := new(UF.UF6)
		uf.NewUF6(30, 2, 0)
		pro = uf
	} else if name == "UF7" {
		uf := new(UF.UF7)
		uf.NewUF7(30, 2, 0)
		pro = uf
	} else if name == "UF8" {
		uf := new(UF.UF8)
		uf.NewUF8(30, 3, 0)
		pro = uf
	} else if name == "UF9" {
		uf := new(UF.UF9)
		uf.NewUF9(30, 3, 0)
		pro = uf
	} else if name == "UF10" {
		uf := new(UF.UF10)
		uf.NewUF10(30, 3, 0)
		pro = uf
	} else if name == "LZ09_F1" {
		lz09 := new(LZ09.LZ09_F1)
		lz09.NewLZ09_F1(30, 2, 0)
		pro = lz09
	} else if name == "LZ09_F2" {
		lz09 := new(LZ09.LZ09_F2)
		lz09.NewLZ09_F2(30, 2, 0)
		pro = lz09
	} else if name == "LZ09_F3" {
		lz09 := new(LZ09.LZ09_F3)
		lz09.NewLZ09_F3(30, 2, 0)
		pro = lz09
	} else if name == "LZ09_F4" {
		lz09 := new(LZ09.LZ09_F4)
		lz09.NewLZ09_F4(30, 2, 0)
		pro = lz09
	} else if name == "LZ09_F5" {
		lz09 := new(LZ09.LZ09_F5)
		lz09.NewLZ09_F5(30, 2, 0)
		pro = lz09
	} else if name == "LZ09_F6" {
		lz09 := new(LZ09.LZ09_F6)
		lz09.NewLZ09_F6(10, 3, 0)
		pro = lz09
	} else if name == "LZ09_F7" {
		lz09 := new(LZ09.LZ09_F7)
		lz09.NewLZ09_F7(10, 2, 0)
		pro = lz09
	} else if name == "LZ09_F8" {
		lz09 := new(LZ09.LZ09_F8)
		lz09.NewLZ09_F8(10, 2, 0)
		pro = lz09
	} else if name == "LZ09_F9" {
		lz09 := new(LZ09.LZ09_F9)
		lz09.NewLZ09_F9(30, 2, 0)
		pro = lz09
	} else if name == "ZDT1" {
		zdt := new(ZDT.ZDT1)
		zdt.NewZDT1(30, 2, 0)
		pro = zdt
	} else if name == "ZDT2" {
		zdt := new(ZDT.ZDT2)
		zdt.NewZDT2(30, 2, 0)
		pro = zdt
	} else if name == "ZDT3" {
		zdt := new(ZDT.ZDT3)
		zdt.NewZDT3(30, 2, 0)
		pro = zdt
	} else if name == "ZDT4" {
		zdt := new(ZDT.ZDT4)
		zdt.NewZDT4(10, 2, 0)
		pro = zdt
	} else if name == "ZDT6" {
		zdt := new(ZDT.ZDT6)
		zdt.NewZDT6(10, 2, 0)
		pro = zdt
	} else if name == "DTLZ1" {
		dtlz := new(DTLZ.DTLZ1)
		dtlz.NewDTLZ1(7, 3, 0)
		pro = dtlz
	} else if name == "DTLZ2" {
		dtlz := new(DTLZ.DTLZ2)
		dtlz.NewDTLZ2(12, 3, 0)
		pro = dtlz
	} else if name == "DTLZ3" {
		dtlz := new(DTLZ.DTLZ3)
		dtlz.NewDTLZ3(12, 3, 0)
		pro = dtlz
	} else if name == "DTLZ4" {
		dtlz := new(DTLZ.DTLZ4)
		dtlz.NewDTLZ4(12, 3, 0)
		pro = dtlz
	} else if name == "DTLZ5" {
		dtlz := new(DTLZ.DTLZ5)
		dtlz.NewDTLZ5(12, 3, 0)
		pro = dtlz
	} else if name == "DTLZ6" {
		dtlz := new(DTLZ.DTLZ6)
		dtlz.NewDTLZ6(12, 3, 0)
		pro = dtlz
	} else if name == "DTLZ7" {
		dtlz := new(DTLZ.DTLZ7)
		dtlz.NewDTLZ7(22, 3, 0)
		pro = dtlz
	} else {
		fmt.Println("No ", name, ",evoluate default problem UF1...")
		uf := new(UF.UF1)
		uf.NewUF1(30, 2, 0)
		pro = uf
	}

	return pro
}

func GetTestInstances(id int) []string {
	var testInstances []string

	if id == 0 {
		testInstances = []string{"UF1", "UF2", "UF3", "UF4", "UF5", "UF6", "UF7", "UF8", "UF9", "UF10"}
	} else if id == 1 {
		testInstances = []string{"LZ09_F1", "LZ09_F2", "LZ09_F3", "LZ09_F4", "LZ09_F5", "LZ09_F6", "LZ09_F7", "LZ09_F8", "LZ09_F9"}
	} else if id == 2 {
		testInstances = []string{"ZDT1", "ZDT2", "ZDT3", "ZDT4", "ZDT6"}
	} else if id == 3 {
		testInstances = []string{"DTLZ1", "DTLZ2", "DTLZ3", "DTLZ4", "DTLZ5", "DTLZ6", "DTLZ7"}
	} else if id == 4 {
		testInstances = []string{"UF1", "UF2", "UF3", "UF4", "UF5", "UF6", "UF7", "UF8", "UF9", "UF10", "LZ09_F1", "LZ09_F2", "LZ09_F3", "LZ09_F4", "LZ09_F5", "LZ09_F6", "LZ09_F7", "LZ09_F8", "LZ09_F9", "ZDT1", "ZDT2", "ZDT3", "ZDT4", "ZDT6", "DTLZ1", "DTLZ2", "DTLZ3", "DTLZ4", "DTLZ5", "DTLZ6", "DTLZ7"}
	} else {
		testInstances = []string{"UF1"}
	}

	return testInstances
}
