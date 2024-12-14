package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func getMinCost(ax, ay, bx, by, px, py int) int {
	am, bm := float64(ay)/float64(ax), float64(by)/float64(bx)
	bn := float64(py) - bm*float64(px)
	x := bn / (am - bm)

	xInt := int(math.Round(x))
	if xInt%ax == 0 && (px-xInt)%bx == 0 {
		return xInt/ax*3 + (px-xInt)/bx
	}
	return 0
}

func main() {
	file, err := os.Open("data/d13/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sum := 0
	sumOffset, offset := 0, 10000000000000
	var ax, ay, bx, by, px, py int
	reButton := regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		switch i % 4 {
		case 0, 1:
			matches := reButton.FindStringSubmatch(line)
			x, xErr := strconv.Atoi(matches[1])
			y, yErr := strconv.Atoi(matches[2])
			if xErr != nil || yErr != nil {
				fmt.Println("Error converting:", line)
				return
			}
			if i%4 == 0 {
				ax, ay = x, y
			} else {
				bx, by = x, y
			}
		case 2:
			matches := rePrize.FindStringSubmatch(line)
			var pxErr, pyErr error
			px, pxErr = strconv.Atoi(matches[1])
			py, pyErr = strconv.Atoi(matches[2])
			if pxErr != nil || pyErr != nil {
				fmt.Println("Error converting", line)
				return
			}
		case 3:
			sum += getMinCost(ax, ay, bx, by, px, py)
			sumOffset += getMinCost(ax, ay, bx, by, px+offset, py+offset)
		}
	}

	sum += getMinCost(ax, ay, bx, by, px, py)
	sumOffset += getMinCost(ax, ay, bx, by, px+offset, py+offset)

	fmt.Println("Minimum cost:", sum)
	fmt.Println("Minimum cost with offset:", sumOffset)
}
