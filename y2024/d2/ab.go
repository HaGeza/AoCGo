package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	safeAscending := sort.SliceIsSorted(levels, func(j, k int) bool {
		diff := levels[j] - levels[k]
		return diff < 1 || diff > 3
	})
	if !safeAscending {
		safeDescending := sort.SliceIsSorted(levels, func(j, k int) bool {
			diff := levels[k] - levels[j]
			return diff < 1 || diff > 3
		})
		if safeDescending {
			return true
		}
	} else {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("data/d2/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	countSafe := 0
	countAlmostSafe := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		levels := make([]int, len(parts))
		for i, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to number:", part)
				return
			}
			levels[i] = level
		}

		if isSafe(levels) {
			countSafe++
			countAlmostSafe++
		} else {
			// Could be done in linear time, by finding the number of non-safe neighbors, then:
			// - If there are more than 2, it is not almost safe
			// - If there are two but not next to each other, it is not almost safe
			// - If there are two next to each other, remove the middle index and check again
			// - If there is only one, check with removing the first as well as removing the second
			// But this is so inelegant that I just refuse to do it. Quadratic time complexity it is.
			for i := range levels {
				skipped := slices.Concat(levels[:i], levels[i+1:])
				if isSafe(skipped) {
					countAlmostSafe++
					break
				}
			}
		}

	}
	fmt.Println("Safe count:", countSafe)
	fmt.Println("Almost safe count:", countAlmostSafe)
}
