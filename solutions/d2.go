package solutions

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func TakeD2Input(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	var result [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		var row []int
		for _, strNum := range strings.Fields(line) {
			var num int
			_, err := fmt.Sscanf(strNum, "%d", &num)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result, nil

}

func D2() int {
	input, err := TakeD2Input(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1
	}
	sum := 0
	for _, row := range input {
		sign := 1
		safe := true
		for i := 1; i < len(row); i++ {
			if (math.Abs(float64(row[i])-float64(row[i-1])) >= 1 && math.Abs(float64(row[i])-float64(row[i-1])) <= 3 && (i == 1 || ((row[i]-row[i-1])*sign > 0))) {
				sign = (row[i] - row[i-1])
			} else {
				safe = false
				break
			}
		}
		if safe {
			sum++
		}
	}
	return sum

}
