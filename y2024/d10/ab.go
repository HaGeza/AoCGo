package main

import (
	"bufio"
	"fmt"
	"os"
)

type Starts struct {
	numStarts int
	numTrails int
	startKeys map[int]bool
}

func main() {
	file, err := os.Open("data/d10/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var matrix [][]byte
	var startsMat [][]Starts
	var stack []([2]int)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		M := len(line)

		matrix = append(matrix, make([]byte, M))
		startsMat = append(startsMat, make([]Starts, M))

		for j, c := range line {
			matrix[i][j] = byte(c) - '0'
			startsMat[i][j] = Starts{numStarts: 0, numTrails: 0, startKeys: make(map[int]bool)}

			if c == '0' {
				startsMat[i][j].numStarts++
				startsMat[i][j].numTrails++
				startsMat[i][j].startKeys[i*M+j] = true
				stack = append(stack, [2]int{i, j})
			}
		}
	}
	N, M := len(matrix), len(matrix[0])

	// Part 1
	sumStarts, sumTrails := 0, 0
	var pos [2]int
	dirsI := [4]int{-1, 0, 1, 0}
	dirsJ := [4]int{0, 1, 0, -1}

	for si := 0; si < len(stack); si++ {
		pos = stack[si]
		i, j := pos[0], pos[1]
		if matrix[i][j] == 9 {
			sumStarts += startsMat[i][j].numStarts
			sumTrails += startsMat[i][j].numTrails
			continue
		}

		for d := 0; d < 4; d++ {
			ii, jj := i+dirsI[d], j+dirsJ[d]
			if ii >= 0 && ii < N && jj >= 0 && jj < M && matrix[ii][jj] == matrix[i][j]+1 {
				if startsMat[ii][jj].numStarts == 0 {
					stack = append(stack, [2]int{ii, jj})
				}
				for key := range startsMat[i][j].startKeys {
					if _, seen := startsMat[ii][jj].startKeys[key]; !seen {
						startsMat[ii][jj].startKeys[key] = true
						startsMat[ii][jj].numStarts++
					}
				}
				startsMat[ii][jj].numTrails += startsMat[i][j].numTrails
			}
		}
	}
	fmt.Println("Sum of num starts:", sumStarts)
	fmt.Println("Sum of num trails:", sumTrails)
}
