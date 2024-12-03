package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getSimpleSum(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`[^m]*mul\((\d+),(\d+)\).*`)

	for _, line := range lines {
		match := re.FindStringSubmatchIndex(line)

		for len(match) >= 6 {
			strA := line[match[2]:match[3]]
			strB := line[match[4]:match[5]]
			a, errA := strconv.Atoi(strA)
			b, errB := strconv.Atoi(strB)

			if errA != nil || errB != nil {
				fmt.Println("Error during conversion:", strA, strB)
				return -1
			}

			sum += a * b
			line = line[match[5]+1:]
			match = re.FindStringSubmatchIndex(line)
		}
	}
	return sum
}

func getComplexSum(lines []string) int {
	sum := 0
	enabled := true
	re := regexp.MustCompile(`[^md]*(do\(\)|don\'t\(\)|mul\((\d+),(\d+)\)).*`)

	for _, line := range lines {
		match := re.FindStringSubmatchIndex(line)

		for len(match) >= 4 {
			strMatch := line[match[2]:match[3]]

			if strMatch == "do()" {
				enabled = true
			} else if strMatch == "don't()" {
				enabled = false
			} else if enabled && len(match) >= 8 {
				strA := line[match[4]:match[5]]
				strB := line[match[6]:match[7]]
				a, errA := strconv.Atoi(strA)
				b, errB := strconv.Atoi(strB)

				if errA != nil || errB != nil {
					fmt.Println("Error during conversion:", strA, strB)
					return -1
				}

				sum += a * b
			}
			line = line[match[3]:]
			match = re.FindStringSubmatchIndex(line)
		}
	}
	return sum
}

func main() {
	file, err := os.Open("data/d3/a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	simpleSum := getSimpleSum(lines)
	fmt.Println("Final simple sum:", simpleSum)

	complexSum := getComplexSum(lines)
	fmt.Println("Final complex sum:", complexSum)
}
