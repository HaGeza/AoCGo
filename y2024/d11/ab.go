package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumSplitsAfter(num int, splits map[int](map[int]int), after int) int {
	numSplits, found := splits[num]
	if found {
		if numSplitsAfter, foundAfter := numSplits[after]; foundAfter {
			return numSplitsAfter
		}
	} else {
		splits[num] = make(map[int]int)
		numSplits = splits[num]
	}

	numSplitsAfter := 1
	if after > 0 {
		if num == 0 {
			numSplitsAfter = getNumSplitsAfter(1, splits, after-1)
		} else {
			numDigits, po10 := 1, 10
			for num >= po10 {
				po10 *= 10
				numDigits++
			}
			if numDigits%2 == 0 {
				po10 = int(math.Pow10(numDigits / 2))
				leftHalf, rightHalf := num/po10, num%po10
				numSplitsAfter = getNumSplitsAfter(leftHalf, splits, after-1) +
					getNumSplitsAfter(rightHalf, splits, after-1)
			} else {
				numSplitsAfter = getNumSplitsAfter(num*2024, splits, after-1)
			}
		}
	}

	numSplits[after] = numSplitsAfter
	return numSplitsAfter
}

func main() {
	file, err := os.Open("data/d11/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sum25, sum75 := 0, 0
	splits := make(map[int](map[int]int))

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Empty file!")
		return
	}

	line := scanner.Text()
	parts := strings.Split(line, " ")

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting:", part)
			return
		}
		sum25 += getNumSplitsAfter(num, splits, 25)
		sum75 += getNumSplitsAfter(num, splits, 75)
	}
	fmt.Println("After 25:", sum25)
	fmt.Println("After 75:", sum75)
}
