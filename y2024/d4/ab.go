package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("data/d4/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]byte
	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, []byte(line))
	}

	reForward := regexp.MustCompile("XMAS")
	reBackward := regexp.MustCompile("SAMX")
	numXMAS := 0

	N := len(matrix)
	M := len(matrix[0])

	for _, row := range matrix {
		numXMAS += len(reForward.FindAllIndex(row, -1))
		numXMAS += len(reBackward.FindAllIndex(row, -1))
	}
	for j := 0; j < M; j++ {
		var column []byte
		for i := 0; i < N; i++ {
			column = append(column, matrix[i][j])
		}
		numXMAS += len(reForward.FindAllIndex(column, -1))
		numXMAS += len(reBackward.FindAllIndex(column, -1))
	}
	for i := 0; i < N; i++ {
		var mainDiag []byte
		ii := i
		for j := 0; ii < N && j < M; ii, j = ii+1, j+1 {
			mainDiag = append(mainDiag, matrix[ii][j])
		}
		numXMAS += len(reForward.FindAllIndex(mainDiag, -1))
		numXMAS += len(reBackward.FindAllIndex(mainDiag, -1))

		var secondaryDiag []byte
		ii = i
		for j := 0; ii >= 0 && j < M; ii, j = ii-1, j+1 {
			secondaryDiag = append(secondaryDiag, matrix[ii][j])
		}
		numXMAS += len(reForward.FindAllIndex(secondaryDiag, -1))
		numXMAS += len(reBackward.FindAllIndex(secondaryDiag, -1))
	}
	for j := 1; j < M; j++ {
		var mainDiag []byte
		jj := j
		for i := 0; i < N && jj < M; i, jj = i+1, jj+1 {
			mainDiag = append(mainDiag, matrix[i][jj])
		}
		numXMAS += len(reForward.FindAllIndex(mainDiag, -1))
		numXMAS += len(reBackward.FindAllIndex(mainDiag, -1))

		var secondaryDiag []byte
		jj = j
		for i := N - 1; i >= 0 && jj < M; i, jj = i-1, jj+1 {
			secondaryDiag = append(secondaryDiag, matrix[i][jj])
		}
		numXMAS += len(reForward.FindAllIndex(secondaryDiag, -1))
		numXMAS += len(reBackward.FindAllIndex(secondaryDiag, -1))
	}
	fmt.Println("Number of XMAS:", numXMAS)

	// Part two
	numCrossMAS := 0
	reCrossMAS := regexp.MustCompile("(M.S.A.M.S|M.M.A.S.S|S.M.A.S.M|S.S.A.M.M)")
	for i := 0; i < N-2; i++ {
		for j := 0; j < M-2; j++ {
			var window [9]byte
			for ii := 0; ii < 3; ii++ {
				for jj := 0; jj < 3; jj++ {
					window[ii*3+jj] = matrix[i+ii][j+jj]
				}
			}
			if reCrossMAS.Match(window[:]) {
				numCrossMAS++
			}
		}
	}
	fmt.Println("Number of X-MAS:", numCrossMAS)
}
