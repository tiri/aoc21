package day09

import (
	"bufio"
	"os"
	"strconv"
)

func CalcRiskLevels(data [][]uint) uint {
	sum := uint(0)

	for _, pos := range findLowPoints(data) {
		sum += data[pos[1]][pos[0]] + 1
	}

	return sum
}

func CalcRiskLevelsFile(file *os.File) uint {
	return CalcRiskLevels(parseFile(file))
}

func CalcBasin(data [][]uint) uint {
	var basins []uint

	for _, pos := range findLowPoints(data) {
		sum, _ := analyseBasin(data, pos[0], pos[1], uint(0))
		basins = append(basins, sum)
	}

	firstMaxElement, firstMaxPos := findMax(basins)
	secondMaxElement, secondMaxPos := findMax(append(basins[:firstMaxPos], basins[firstMaxPos+1:]...))
	thirdMaxElement, _ := findMax(append(basins[:secondMaxPos], basins[secondMaxPos+1:]...))

	return firstMaxElement * secondMaxElement * thirdMaxElement
}

func analyseBasin(data [][]uint, x uint, y uint, sum uint) (uint, [][]uint) {
	sum += 1
	data[y][x] = 9

	if x > 0 && data[y][x-1] < 9 {
		sum, data = analyseBasin(data, x-1, y, sum)
	}

	if x < uint(len(data[y]))-1 && data[y][x+1] < 9 {
		sum, data = analyseBasin(data, x+1, y, sum)
	}

	if y > 0 && data[y-1][x] < 9 {
		sum, data = analyseBasin(data, x, y-1, sum)
	}

	if y < uint(len(data))-1 && data[y+1][x] < 9 {
		sum, data = analyseBasin(data, x, y+1, sum)
	}

	return sum, data
}

func CalcBasinFile(file *os.File) uint {
	return CalcBasin(parseFile(file))
}

func parseFile(file *os.File) [][]uint {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var data [][]uint

	for scanner.Scan() {
		var row []uint
		for _, item := range scanner.Text() {
			number, _ := strconv.ParseUint(string(item), 10, 32)
			row = append(row, uint(number))
		}

		data = append(data, row)
	}
	return data
}

func findLowPoints(data [][]uint) [][2]uint {
	var lowPoints [][2]uint

	for y, row := range data {
		for x, value := range row {
			top := x > 0 && data[y][x-1] <= value
			left := y > 0 && data[y-1][x] <= value
			right := x < len(row)-1 && data[y][x+1] <= value
			bottom := y < len(data)-1 && data[y+1][x] <= value
			if !top && !left && !right && !bottom {
				lowPoints = append(lowPoints, [2]uint{uint(x), uint(y)})
			}
		}
	}

	return lowPoints
}

func findMax(array []uint) (uint, uint) {
	max := array[0]
	pos := uint(0)

	for index, element := range array {
		if element > max {
			pos = uint(index)
			max = element
		}
	}

	return max, pos
}
