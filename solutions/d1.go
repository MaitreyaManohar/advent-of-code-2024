package solutions

import (
	"fmt"
	"math"
	"sort"
)


// Might break if the input is not given in pairs ending with "0 0"
func TakeInputAndSort() ([]int, []int) {
	var list1 []int
	var list2 []int
	for {
		var n1 int
		var n2 int
		fmt.Scanf("%d %d", &n1, &n2)
		if n1 == 0 && n2 == 0 {
			break
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}


func Part1() int {
	list1, list2 := TakeInputAndSort()
	sum := 0
	for i := range(list1) {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return sum
}

func Part2() int {
	list1, list2 := TakeInputAndSort()
	sum := 0
	freqMap := make(map[int]int)
	for _, val := range(list2) {
		freqMap[val]++;
	}
	for _, val := range(list1) {
		sum += freqMap[val] * val;
	}
	return sum
}