package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

func toInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	fileName := "a"
	file, err := os.Open(fmt.Sprintf("data/d14/%s.txt", fileName))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var N, M int
	if fileName == "test" {
		N, M = 7, 11
	} else {
		N, M = 103, 101
	}

	var positions, velocities [][2]int
	var finalPositions [][2]int
	steps := 100

	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-{0,1}\d+),(-{0,1}\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		x, y := toInt(matches[1]), toInt(matches[2])
		vx, vy := toInt(matches[3]), toInt(matches[4])
		positions = append(positions, [2]int{x, y})
		velocities = append(velocities, [2]int{vx, vy})
		finalPositions = append(finalPositions, [2]int{
			((x+steps*vx)%M + M) % M,
			((y+steps*vy)%N + N) % N,
		})
	}

	sums := [2][2]int{{0, 0}, {0, 0}}
	for _, pos := range finalPositions {
		x, y := pos[0], pos[1]
		if x != M/2 && y != N/2 {
			sums[2*x/M][2*y/N]++
		}
	}
	prod := 1
	for _, sum := range sums {
		for _, s := range sum {
			prod *= s
		}
	}
	fmt.Println("After", steps, "seconds:", prod)

	// Part 2
	grid := make([][]bool, N)
	for i := range grid {
		grid[i] = make([]bool, M)
	}
	for _, pos := range positions {
		grid[pos[1]][pos[0]] = true
	}

	for s := 0; s < 10000; s++ {
		hasBlock := false
		for i := 0; i < len(grid)-2; i++ {
			for j := 0; j < len(grid[i])-2; j++ {
				if grid[i][j] {
					block := true
					for ii := i; ii < i+3; ii++ {
						for jj := j; jj < j+3; jj++ {
							block = block && grid[ii][jj]
						}
					}
					hasBlock = hasBlock || block
				}
			}
		}
		if hasBlock {
			fmt.Println("Found block at", s)
		}

		// Create a new binary image
		img := image.NewGray(image.Rect(0, 0, M, N))

		// Set the pixels based on the grid
		for y, row := range grid {
			for x, cell := range row {
				if cell {
					img.SetGray(x, y, color.Gray{Y: 255})
				} else {
					img.SetGray(x, y, color.Gray{Y: 0})
				}
			}
		}

		// Save the image to a file
		fileName := fmt.Sprintf("data/d14/p2/output_%05d.png", s)
		outFile, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer outFile.Close()
		png.Encode(outFile, img)

		for i, row := range grid {
			for j := range row {
				grid[i][j] = false
			}
		}

		for i := 0; i < len(positions); i++ {
			positions[i][0] = (positions[i][0] + velocities[i][0] + M) % M
			positions[i][1] = (positions[i][1] + velocities[i][1] + N) % N
			grid[positions[i][1]][positions[i][0]] = true
		}
	}
}
