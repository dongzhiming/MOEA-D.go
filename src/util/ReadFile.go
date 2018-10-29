package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadDoubleDataFile(path string) [][]float64 {
	var ret [][]float64 = make([][]float64, 0)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to open ", path)
		return nil
	}
	defer file.Close()

	i := 0

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //每次读取一行
		if err != nil {
			break // 读完或发生错误
		}

		var in []float64 = make([]float64, 0)

		splitStr := strings.Fields(str) //将数据安装空格分开
		for j := 0; j < len(splitStr); j++ {
			var value float64
			value, _ = strconv.ParseFloat(splitStr[j], 64)
			in = append(in, value)
		}

		ret = append(ret, in)

		i++
	}

	return ret
}
