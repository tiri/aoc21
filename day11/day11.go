package day11

import (
	"bufio"
	"os"
	"strconv"
)

func CalcEnergyLevel(data [][]int, iterations int) uint {
	litUp := uint(0)

	for i := 0; i < iterations; i++ {
		litUp += nextIterationEnergyLevel(data)
	}

	return litUp
}

func CalcEnergyLevelFile(file *os.File, iterations int) uint {
	return CalcEnergyLevel(parseFile(file), iterations)
}

func nextIterationEnergyLevel(data [][]int) uint {
	flashes := uint(0)

	// increase adjacent octopuses
	for y, row := range data {
		for x, _ := range row {
			flashes += recursiveIncrease(data, y, x)
		}
	}

	// set flashed elements to zero
	for y, row := range data {
		for x, _ := range row {
			if data[y][x] > 9 {
				data[y][x] = 0
			}
		}
	}

	return flashes
}

func recursiveIncrease(data [][]int, y int, x int) uint {
	flashes := uint(0)

	data[y][x]++

	if data[y][x] == 10 {
		flashes++

		for yRunner := y - 1; yRunner <= y+1; yRunner++ {
			for xRunner := x - 1; xRunner <= x+1; xRunner++ {
				// check out of bounds && same as origin pos
				if yRunner < 0 || xRunner < 0 ||
					y == yRunner && x == xRunner ||
					yRunner > len(data)-1 || xRunner > len(data[0])-1 {
					continue
				}

				flashes += recursiveIncrease(data, yRunner, xRunner)
			}
		}
	}
	return flashes
}

func HowMuchStepsToFullFlash(data [][]int) uint {
	allFlash := uint(len(data[0]) * len(data))

	iterations := uint(0)
	for {
		iterations++

		litUp := nextIterationEnergyLevel(data)

		if allFlash == litUp {
			return iterations
		}
	}
}

func HowMuchStepsToFullFlashFile(file *os.File) uint {
	return HowMuchStepsToFullFlash(parseFile(file))
}

func parseFile(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var data [][]int

	for scanner.Scan() {
		var row []int
		for _, item := range scanner.Text() {
			num, _ := strconv.ParseInt(string(item), 10, 32)
			row = append(row, int(num))
		}
		data = append(data, row)
	}
	return data
}
