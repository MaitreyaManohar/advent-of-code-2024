package solutions

import (
	"os"
	"strconv"
	"strings"
)

func TakeD5Input() (string, error) {
	data, err := os.ReadFile("inputs/d5.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func D5Part1(inputString string) int {
	split := strings.Split(inputString, "\n\n")
	rules := split[0]
	orderings := split[1]
	orderingRules := make(map[int][]int, 0)
	for _, rule := range strings.Split(rules, "\n") {
		ruleKey, _ := strconv.Atoi(strings.Split(rule, "|")[0])
		ruleValue, _ := strconv.Atoi(strings.Split(rule, "|")[1])
		orderingRules[ruleKey] = append(orderingRules[ruleKey], ruleValue)
	}
	sum := 0
	for _, ordering := range strings.Split(orderings, "\n") {
		orderingNums := convertToNums(ordering)
		correct := true
	OuterLoop:
		for i, num := range orderingNums {
			for j := range i {
				for _, val := range orderingRules[num] {
					if orderingNums[j] == val {
						correct = false
						break OuterLoop
					}
				}
			}
		}
		if correct {
			sum += returnMiddle(&orderingNums)
		}
	}
	return sum
}

func D5Part2(inputString string) int {
	split := strings.Split(inputString, "\n\n")
	rules := split[0]
	orderings := split[1]
	orderingRules := make(map[int][]int, 0)
	for _, rule := range strings.Split(rules, "\n") {
		ruleKey, _ := strconv.Atoi(strings.Split(rule, "|")[0])
		ruleValue, _ := strconv.Atoi(strings.Split(rule, "|")[1])
		orderingRules[ruleKey] = append(orderingRules[ruleKey], ruleValue)
	}
	sum := 0
	for _, ordering := range strings.Split(orderings, "\n") {
		orderingNums := convertToNums(ordering)
		correct := true
	OuterLoop:
		for i, num := range orderingNums {
			for j := range i {
				for _, val := range orderingRules[num] {
					if orderingNums[j] == val {
						correct = false
						break OuterLoop
					}
				}
			}
		}
		if !correct {
			// Find the last element to keep in the ordering
			// That would be the element which does not have
			// any of the other elements in its ordering
			// rules.
			correctList := make([]int, 0)
			for len(correctList) != len(orderingNums) {
				for _, num := range orderingNums {
					innerCorrect := true
					if contains(correctList, num) {
						continue
					}
					for _, innerNum := range orderingNums {
						if contains(correctList, innerNum) {
							continue
						}
						if contains(orderingRules[num], innerNum) {
							innerCorrect = false
							break
						}
					}
					if innerCorrect {
						correctList = append(correctList, num)
					}
				}
			}
			sum += returnMiddle(&correctList)
		}
	}
	return sum
}

func contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func returnMiddle(nums *[]int) int {
	return (*nums)[len(*nums)/2]
}

func convertToNums(ordering string) []int {
	nums := make([]int, 0)
	for _, num := range strings.Split(ordering, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(num))
		nums = append(nums, n)
	}
	return nums
}
