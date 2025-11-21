package solutions

import (
	// "fmt"
	"os"
	// "strings"
)

func TakeD4Input() (string, error) {
	input, err := os.ReadFile("inputs/d4.txt")
	if err != nil {
		return "", err
	}
	return string(input), nil
}

func D4Part1(inputString string) int {
	var sum int
	charMatrix := CreateCharMatrix(inputString)

	for i := range charMatrix {
		for j := range charMatrix[i] {
			if charMatrix[i][j] == 'X' {
				sum += CountMAS(&charMatrix, i, j)
			}
		}
	}
	return sum
}

func CreateCharMatrix(inputString string) [][]byte {
	charMatrix := make([][]byte, 0)
	row := make([]byte, 0)
	for i := range inputString {
		if inputString[i] == '\n' {
			charMatrix = append(charMatrix, row)
			row = make([]byte, 0)
		} else {
			row = append(row, inputString[i])
		}
	}
	return charMatrix
}

func CountMAS(charMatrix *[][]byte, i int, j int) int {
	dx := [8][3]int{{1, 2, 3}, {1, 2, 3}, {0, 0, 0}, {-1, -2, -3}, {-1, -2, -3}, {-1, -2, -3}, {0, 0, 0}, {1, 2, 3}}
	dy := [8][3]int{{0, 0, 0}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {0, 0, 0}, {-1, -2, -3}, {-1, -2, -3}, {-1, -2, -3}}
	count := 0
	for p := range 8 {
		check := true
		chars := [3]byte{'M', 'A', 'S'}
		for l := range 3 {
			x := i + dx[p][l]
			y := j + dy[p][l]
			if checkWithinBounds(charMatrix, x, y) && ((*charMatrix)[x][y] == chars[l]) {

			} else {
				check = false
				break
			}
		}

		if check {
			count += 1
		}
	}

	return count
}

func checkWithinBounds(charMatrix *[][]byte, i int, j int) bool {
	return (i >= 0) && (i < len(*charMatrix)) && (j >= 0) && (j < len((*charMatrix)[0]))
}

func CountXMAS(charMatrix *[][]byte, i int, j int) int {
	// Check left diagonal
	leftDiagonal := false
	checkBounds := checkWithinBounds(charMatrix, i-1, j-1) && checkWithinBounds(charMatrix, i+1, j+1)
	if checkBounds && (((*charMatrix)[i-1][j-1] == 'M' && (*charMatrix)[i+1][j+1] == 'S') || ((*charMatrix)[i-1][j-1] == 'S' && (*charMatrix)[i+1][j+1] == 'M')) {
		 leftDiagonal = true
	}
	// Check right diagonal
	rightDiagonal := false
	checkBounds = checkWithinBounds(charMatrix, i+1, j-1) && checkWithinBounds(charMatrix, i-1, j+1)
	if checkBounds && (((*charMatrix)[i+1][j-1] == 'M' && (*charMatrix)[i-1][j+1] == 'S') || ((*charMatrix)[i+1][j-1] == 'S' && (*charMatrix)[i-1][j+1] == 'M')) {
		rightDiagonal = true
	}
	if leftDiagonal && rightDiagonal {
		return 1
	}
	return 0
}

func D4Part2(inputString string) int {
	charMatrix := CreateCharMatrix(inputString)
	sum := 0
	for i := range charMatrix {
		for j := range charMatrix[i] {
			if charMatrix[i][j] == 'A' {
				sum += CountXMAS(&charMatrix, i, j)
			}
		}
	}
	return sum
}
