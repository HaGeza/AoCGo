package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data/d8/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freqPositions := make(map[byte]([][2]int))
	N, M := 0, 0

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			if c != '.' {
				bc := byte(c)
				cPositions, cExists := freqPositions[bc]
				pos := [2]int{i, j}
				if !cExists {
					freqPositions[bc] = [][2]int{pos}
				} else {
					freqPositions[bc] = append(cPositions, pos)
				}
			}
			M = max(j+1, M)
		}
		N = max(i+1, N)
	}

	antinodes := make(map[int]bool)
	repeatedAntinodes := make(map[int]bool)
	sumAntinodes := 0
	sumRepeatedAntinodes := 0

	for _, positions := range freqPositions {
		for i, iPos := range positions {
			for j, jPos := range positions {
				if i == j {
					continue
				}
				deltaY := iPos[0] - jPos[0]
				deltaX := iPos[1] - jPos[1]

			antinodeChecking:
				for d := 0; true; d++ {
					antinodeX := iPos[0] + d*deltaY
					antinodeY := iPos[1] + d*deltaX

					if antinodeX < 0 || antinodeX >= M || antinodeY < 0 || antinodeY >= N {
						break antinodeChecking
					}

					antinodeKey := M*antinodeY + antinodeX
					if _, keyExists := repeatedAntinodes[antinodeKey]; !keyExists {
						repeatedAntinodes[antinodeKey] = true
						sumRepeatedAntinodes++
					}
					if d == 1 {
						if _, keyExists := antinodes[antinodeKey]; !keyExists {
							antinodes[antinodeKey] = true
							sumAntinodes++
						}
					}
				}
			}
		}
	}
	fmt.Println("Number of unique antinodes:", sumAntinodes)
	fmt.Println("Number of unique repeated antinodes:", sumRepeatedAntinodes)
}
