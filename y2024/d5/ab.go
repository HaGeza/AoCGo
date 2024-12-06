package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data/d5/a.txt")
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

		seenInds := make(map[int]int)
		for _, num := range nums {
			seenInds[num] = -1
		}
		seenInds[nums[0]] = 0

		correct := true

	outerLoop:
		for i := 1; i < len(nums); {
			outNeighbors, exists := graph[nums[i]]
			if exists {
				for neighbor := range outNeighbors {
					seenInd, inNums := seenInds[neighbor]
					if inNums && seenInd > -1 && seenInd < i {
						seenInds[neighbor] = i
						seenInds[nums[i]] = seenInd
						nums[i], nums[seenInd] = nums[seenInd], nums[i]

						correct = false
						i = seenInd
						continue outerLoop
					}
				}
			}
			seenInds[nums[i]] = i
			i++
		}

		if correct {
			correctSum += nums[len(nums)/2]
		} else {
			incorrectSum += nums[len(nums)/2]
		}
	}
	fmt.Println("Correctly ordered update middle-part sum:", correctSum)
	fmt.Println("Incorrectly ordered update middle-part sum:", incorrectSum)
}
