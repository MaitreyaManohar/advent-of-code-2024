package solutions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func TakeD3Input() string {
	// Readfile d3.txt
	data, err := os.ReadFile("inputs/d3.txt")
	if err != nil {
		fmt.Printf("error reading file %+v", err)
		return ""
	}
	return string(data)
}

func D3Part1(inputString string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(inputString, -1)
	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		sum += num1 * num2
	}
	return sum
}

func EndsWithDont(stringSlice string, checkArr string) bool {
	if len(stringSlice) < len(checkArr) {
		return false
	}
	for i := range checkArr  {
		if (checkArr[i] != stringSlice[len(stringSlice) - len(checkArr) + i]) {
			return false
		}
	}
	return true
}

func D3Part2(inputString string) int {
	l := 0
	r := 0
	sum := 0
	checkDont := true
	for l <= r && r < len(inputString) {
		for r < len(inputString) && ((checkDont && !EndsWithDont(inputString[:r], "don't()")) || (!checkDont && !EndsWithDont(inputString[:r], "do()"))) {
			r++
		}
		if checkDont {
			x := r
			if (r == len(inputString)) {
				x--
			}
			sum += D3Part1(inputString[l:x])
		}
		checkDont = !checkDont
		l = r
	}

	return sum
}
