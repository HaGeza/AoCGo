package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data/d1/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var listA, listB []int
	N := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		a, errA := strconv.Atoi(parts[0])
		b, errB := strconv.Atoi(parts[1])
		if errA != nil || errB != nil {
			fmt.Println("Error converting to numbers:", a, b)
			continue
		}

		listA = append(listA, a)
		listB = append(listB, b)
		N++
	}

	// First part
	sort.Ints(listA)
	sort.Ints(listB)

	sum := 0
	for i := 0; i < N; i++ {
		a, b := listA[i], listB[i]
		if a > b {
			sum += a - b
		} else {
			sum += b - a
		}
	}
	fmt.Println("Answer A:", sum)

	// Second part
	counts := make(map[int]int)

	for _, a := range listA {
		counts[a] = 0
	}
	for _, b := range listB {
		count, ok := counts[b]
		if ok {
			counts[b] = count + 1
		}
	}
	similarityScore := 0
	for _, a := range listA {
		count, _ := counts[a]
		similarityScore += a * count
	}
	fmt.Println("Answer B:", similarityScore)
}
