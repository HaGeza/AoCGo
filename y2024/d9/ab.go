package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data/d9/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Empty file")
		return
	}
	line := scanner.Text()

	var data, free []byte
	for i, c := range line {
		if i%2 == 0 {
			data = append(data, byte(c)-'0')
		} else {
			free = append(free, byte(c)-'0')
		}
	}
	dataOriginal := make([]byte, len(data))
	copy(dataOriginal, data)
	freeOriginal := make([]byte, len(free))
	copy(freeOriginal, free)

	sum, pos := 0, 0
outer:
	for i, j := 0, len(data)-1; i < len(data); i++ {
		d := int(data[i])
		sum += i * (d*pos + d*(d-1)/2)
		pos += d

		for free[i] > 0 {
			if j <= i {
				break outer
			}

			moved := min(free[i], data[j])
			free[i] -= moved
			data[j] -= moved

			m := int(moved)
			sum += j * (m*pos + m*(m-1)/2)
			pos += m

			if data[j] == 0 {
				j--
			}
		}
	}
	fmt.Println("Partial move sum:", sum)

	// Part 2
	sum, pos = 0, 0

	copy(data, dataOriginal)
	copy(free, freeOriginal)

	for i := 0; i < len(data); i++ {
		d := int(data[i])
		sum += i * (d*pos + d*(d-1)/2)
		pos += int(dataOriginal[i])

		if i >= len(free) {
			continue
		}

		for j := len(data) - 1; free[i] > 0 && j >= 0; j-- {
			if data[j] > 0 && data[j] <= free[i] {
				m := int(data[j])

				free[i] -= data[j]
				data[j] = 0

				sum += j * (m*pos + m*(m-1)/2)
				pos += m
			}
		}
		pos += int(free[i])
	}
	fmt.Println("Whole move sum:", sum)
}
