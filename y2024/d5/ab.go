package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isOrdered(node int, graph map[int](map[int]bool), predecessors map[int]bool) bool {
	outNeighbors, exists := graph[node]
	if exists {
		for neighbor := range outNeighbors {
			isPredecessor, isConsidered := predecessors[neighbor]
			if isConsidered && isPredecessor {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("data/d5/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(map[int](map[int]bool))

	for true {
		scanned := scanner.Scan()
		line := scanner.Text()

		if scanned && len(line) > 0 {
			nodeStrs := strings.Split(line, "|")
			fromNode, fromErr := strconv.Atoi(nodeStrs[0])
			toNode, toErr := strconv.Atoi(nodeStrs[1])

			if fromErr != nil || toErr != nil {
				fmt.Println("Error converting to int:", fromNode, toNode)
				return
			}
			outNeighbors, fromNodeExists := graph[fromNode]
			if !fromNodeExists {
				outNeighbors = make(map[int]bool)
				graph[fromNode] = outNeighbors
			}
			outNeighbors[toNode] = true
		} else {
			break
		}
	}

	correctSum := 0
	incorrectSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		nums := make([]int, len(parts))
		for i, part := range parts {
			nums[i], err = strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				return
			}
		}

		seenInNums := make(map[int]bool)
		for _, num := range nums {
			seenInNums[num] = false
		}
		seenInNums[nums[0]] = true

		sorted := true
		for _, num := range nums[1:] {
			sorted = sorted && isOrdered(num, graph, seenInNums)
			if !sorted {
				break
			}
			seenInNums[num] = true
		}
		if sorted {
			correctSum += nums[len(nums)/2]
		}
	}
	fmt.Println("Correctly ordered update middle-part sum:", correctSum)
	fmt.Println("Incorrectly ordered update middle-part sum:", incorrectSum)
}
