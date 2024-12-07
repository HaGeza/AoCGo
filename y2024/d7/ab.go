package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPossibleRhs(resultSoFar int, remaining []int) []int {
	if len(remaining) == 0 {
		return []int{resultSoFar}
	}

	newResult := resultSoFar + remaining[0]
	possibleRhs := getPossibleRhs(newResult, remaining[1:])

	newResult = resultSoFar * remaining[0]
	return append(possibleRhs, getPossibleRhs(newResult, remaining[1:])...)
}

func getPossibleRhsWConcat(resultSoFar int, remaining []int) []int {
	if len(remaining) == 0 {
		return []int{resultSoFar}
	}

	newResult := resultSoFar + remaining[0]
	possibleRhs := getPossibleRhsWConcat(newResult, remaining[1:])

	multiplier := 10
	for remaining[0] >= multiplier {
		multiplier *= 10
	}
	newResult = resultSoFar*multiplier + remaining[0]
	possibleRhs = append(possibleRhs, getPossibleRhsWConcat(newResult, remaining[1:])...)

	newResult = resultSoFar * remaining[0]
	return append(possibleRhs, getPossibleRhsWConcat(newResult, remaining[1:])...)
}

func main() {
	file, err := os.Open("data/d7/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	noConcatSum := 0
	concatSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		lhsStr := parts[0][:len(parts[0])-1]
		lhs, err := strconv.Atoi(lhsStr)
		if err != nil {
			fmt.Println("Error converting:", lhsStr)
			return
		}
		var nums []int
		for _, part := range parts[1:] {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting:", part)
				return
			}
			nums = append(nums, num)
		}

		possibleRhs := getPossibleRhs(nums[0], nums[1:])
		for _, possibleRhs := range possibleRhs {
			if possibleRhs == lhs {
				noConcatSum += lhs
				break
			}
		}
		possibleRhs = getPossibleRhsWConcat(nums[0], nums[1:])
		for _, possibleRhs := range possibleRhs {
			if possibleRhs == lhs {
				concatSum += lhs
				break
			}
		}
	}
	fmt.Println("Correct LHS sum without concatenation:", noConcatSum)
	fmt.Println("Correct LHS sum with concatenation:", concatSum)
}
