package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDir(dInd int) [2]int {
	return [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}[dInd]
}

func fill(i, j int, seen [][]bool, matrix [][]byte) (int, int) {
	seen[i][j] = true
	N, M := len(matrix), len(matrix[0])
	area, perimeter := 1, 0

	for dInd := 0; dInd < 4; dInd++ {
		d := getDir(dInd)
		ii, jj := i+d[0], j+d[1]
		if ii < 0 || ii >= N || jj < 0 || jj >= M || matrix[ii][jj] != matrix[i][j] {
			perimeter++
		} else if !seen[ii][jj] {
			otherArea, otherPerimeter := fill(ii, jj, seen, matrix)
			area += otherArea
			perimeter += otherPerimeter
		}
	}
	return area, perimeter
}

func markEdge(i, j, dInd int, seenEdges [][][4]bool, matrix [][]byte, matValue byte) {
	revDir := getDir((dInd + 2) % 4)
	ri, rj := i-1+revDir[0], j-1+revDir[1]
	N, M := len(matrix), len(matrix[0])
	if ri < 0 || ri >= N || rj < 0 || rj >= M || matrix[ri][rj] != matValue || seenEdges[i][j][dInd] {
		return
	}
	if i > 0 && i <= N && j > 0 && j <= M && matrix[i-1][j-1] == matValue {
		return
	}

	seenEdges[i][j][dInd] = true
	d1, d2 := getDir((dInd+1)%4), getDir((dInd+3)%4)
	markEdge(i+d1[0], j+d1[1], dInd, seenEdges, matrix, matValue)
	markEdge(i+d2[0], j+d2[1], dInd, seenEdges, matrix, matValue)
}

func fillEdges(i, j int, seen [][]bool, seenEdges [][][4]bool, matrix [][]byte) (int, int) {
	seen[i][j] = true

	N, M := len(matrix), len(matrix[0])
	area, edges := 1, 0

	for dInd := 0; dInd < 4; dInd++ {
		d := getDir(dInd)
		ii, jj := i+d[0], j+d[1]
		if ii < 0 || ii >= N || jj < 0 || jj >= M || matrix[ii][jj] != matrix[i][j] {
			if !seenEdges[ii+1][jj+1][dInd] {
				edges++
				markEdge(ii+1, jj+1, dInd, seenEdges, matrix, matrix[i][j])
			}
		} else if !seen[ii][jj] {
			otherArea, otherEdges := fillEdges(ii, jj, seen, seenEdges, matrix)
			area += otherArea
			edges += otherEdges
		}
	}
	return area, edges
}

func main() {
	file, err := os.Open("data/d12/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var matrix [][]byte
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		M := len(line)
		matrix = append(matrix, make([]byte, M))
		for j, c := range line {
			matrix[i][j] = byte(c)
		}
	}

	N, M := len(matrix), len(matrix[0])

	// Part 1
	seen := make([][]bool, N)
	for i := 0; i < N; i++ {
		seen[i] = make([]bool, M)
	}

	sumCost := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !seen[i][j] {
				area, perimeter := fill(i, j, seen, matrix)
				sumCost += area * perimeter
			}
		}
	}
	fmt.Println("Cost of fence:", sumCost)

	// Part 2
	seenEdges := make([][][4]bool, N+2)
	for i := 0; i < N+2; i++ {
		seenEdges[i] = make([][4]bool, M+2)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			seen[i][j] = false
		}
	}
	sumCost = 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !seen[i][j] {
				area, edges := fillEdges(i, j, seen, seenEdges, matrix)
				sumCost += area * edges
			}
		}
	}
	fmt.Println("Cost fence with edges:", sumCost)
}
