package main

import (
	"bufio"
	"fmt"
	"os"
)

func simulatePatrol(si, sj int, matrix [][]byte) (int, bool) {
	sum := 0
	movements := map[byte]([2]int){
		'u': [2]int{-1, 0},
		'r': [2]int{0, 1},
		'd': [2]int{1, 0},
		'l': [2]int{0, -1},
	}
	directions := [4]byte{'u', 'r', 'd', 'l'}
	directionInd := 0
	direction := directions[directionInd]

	N, M := len(matrix), len(matrix[0])
	pathMat := make([][]byte, N)
	for i := range matrix {
		pathMat[i] = make([]byte, M)
		copy(pathMat[i], matrix[i])
	}

	loop := false
	for true {
		if pathMat[si][sj] == '.' || pathMat[si][sj] == '^' {
			sum++
			pathMat[si][sj] = direction
		}
		movement := movements[direction]
		ni, nj := si+movement[0], sj+movement[1]
		if ni < 0 || nj < 0 || ni >= N || nj >= M {
			break
		}
		if pathMat[ni][nj] == direction {
			loop = true
			break
		}

		if pathMat[ni][nj] == '#' {
			directionInd = (directionInd + 1) % len(directions)
			direction = directions[directionInd]
		} else {
			si, sj = ni, nj
		}
	}
	return sum, loop
}

func main() {
	file, err := os.Open("data/d6/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]byte
	var si, sj int

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j, c := range line {
			if c == '^' {
				si, sj = i, j
			}
		}
		matrix = append(matrix, []byte(line))
	}
	N, M := len(matrix), len(matrix[0])

	// Part one
	sum, _ := simulatePatrol(si, sj, matrix)
	fmt.Println("Tiles covered:", sum)

	// Part two
	sum = 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if matrix[i][j] == '^' || matrix[i][j] == '#' {
				continue
			}
			matrix[i][j] = '#'
			_, loop := simulatePatrol(si, sj, matrix)
			if loop {
				sum++
			}
			matrix[i][j] = '.'
		}
	}
	fmt.Println("Number of possible looping obstacle positions:", sum)
}
