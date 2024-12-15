package main

import (
	"bufio"
	"fmt"
	"os"
)

func getBoxJPositions(boxIPos int, boxC byte) [2]int {
	if boxC == '[' {
		return [2]int{boxIPos, boxIPos + 1}
	} else {
		return [2]int{boxIPos - 1, boxIPos}
	}
}

func pushDoubleBoxesUpDown(grid [][]byte, pushI int, dir [2]int, pushJs map[int]bool) bool {
	newPushJs := make(map[int]bool)
	for pushJ := range pushJs {
		switch grid[pushI][pushJ] {
		case '#':
			return false
		case '[', ']':
			jPositions := getBoxJPositions(pushJ, grid[pushI][pushJ])
			for j := jPositions[0]; j <= jPositions[1]; j++ {
				newPushJs[j] = true
			}
		}
	}

	if len(newPushJs) > 0 && !pushDoubleBoxesUpDown(grid, pushI+dir[0], dir, newPushJs) {
		return false
	}
	for pushJ := range newPushJs {
		grid[pushI+dir[0]][pushJ] = grid[pushI][pushJ]
		grid[pushI][pushJ] = '.'
	}
	return true
}

func moveRobot(grid [][]byte, pos, dir [2]int) [2]int {
	newPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
	switch grid[newPos[0]][newPos[1]] {
	case '#':
		return pos
	case '.':
		grid[pos[0]][pos[1]] = '.'
		grid[newPos[0]][newPos[1]] = '@'
	case 'O':
		pushPos := [2]int{newPos[0] + dir[0], newPos[1] + dir[1]}
		for grid[pushPos[0]][pushPos[1]] == 'O' {
			pushPos[0] += dir[0]
			pushPos[1] += dir[1]
		}
		if grid[pushPos[0]][pushPos[1]] == '#' {
			return pos
		} else {
			grid[pushPos[0]][pushPos[1]] = 'O'
			grid[newPos[0]][newPos[1]] = '@'
			grid[pos[0]][pos[1]] = '.'
		}
	case '[', ']':
		if dir[0] == 0 {
			pushJ := newPos[1]
			for grid[pos[0]][pushJ] == '[' || grid[pos[0]][pushJ] == ']' {
				pushJ += dir[1]
			}
			if grid[pos[0]][pushJ] == '#' {
				return pos
			} else {
				c, other := '[', ']'
				if dir[1] == 1 {
					c, other = other, c
				}
				for j := pushJ; j != newPos[1]; j -= dir[1] {
					grid[pos[0]][j] = byte(c)
					c, other = other, c
				}
				grid[newPos[0]][newPos[1]] = '@'
				grid[pos[0]][pos[1]] = '.'
			}
		} else if !pushDoubleBoxesUpDown(grid, newPos[0], dir, map[int]bool{newPos[1]: true}) {
			return pos
		}
	}
	return newPos
}

func main() {
	file, err := os.Open("data/d15/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid, doubleGrid [][]byte
	var pos, doublePos [2]int

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		grid = append(grid, []byte(line))
		doubleGrid = append(doubleGrid, make([]byte, 2*len(line)))
		for j := 0; j < len(line); j++ {
			if line[j] == '@' {
				pos = [2]int{i, j}
				doublePos = [2]int{i, 2 * j}
				doubleGrid[i][2*j] = '@'
				doubleGrid[i][2*j+1] = '.'
			} else if line[j] == 'O' {
				doubleGrid[i][2*j] = '['
				doubleGrid[i][2*j+1] = ']'
			} else {
				doubleGrid[i][2*j] = line[j]
				doubleGrid[i][2*j+1] = line[j]
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range line {
			switch c {
			case '<':
				pos = moveRobot(grid, pos, [2]int{0, -1})
				doublePos = moveRobot(doubleGrid, doublePos, [2]int{0, -1})
			case '>':
				pos = moveRobot(grid, pos, [2]int{0, 1})
				doublePos = moveRobot(doubleGrid, doublePos, [2]int{0, 1})
			case '^':
				pos = moveRobot(grid, pos, [2]int{-1, 0})
				doublePos = moveRobot(doubleGrid, doublePos, [2]int{-1, 0})
			case 'v':
				pos = moveRobot(grid, pos, [2]int{1, 0})
				doublePos = moveRobot(doubleGrid, doublePos, [2]int{1, 0})
			}
		}
	}

	sum, doubleSum := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'O' {
				sum += 100*i + j
			}
			for k := 2 * j; k < 2*j+2; k++ {
				if doubleGrid[i][k] == '[' {
					doubleSum += 100*i + k
				}
			}
		}
	}

	fmt.Println("Simple sum:", sum)
	fmt.Println("Double sum:", doubleSum)
}
