package day09

import (
	"bufio"
	"os"
	"sort"
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
	var basins []int

	for _, pos := range findLowPoints(data) {
		sum := analyseBasin(data, pos[0], pos[1])
		basins = append(basins, int(sum))
	}

	sort.Ints(basins)

	sum := uint(1)
	for _, value := range basins[len(basins)-3:] {
		sum *= uint(value)
	}
	return sum
}

func analyseBasin(data [][]uint, x uint, y uint) uint {
	sum := uint(1)
	data[y][x] = 9

	if x > 0 && data[y][x-1] < 9 {
		sum += analyseBasin(data, x-1, y)
	}

	if x < uint(len(data[y]))-1 && data[y][x+1] < 9 {
		sum += analyseBasin(data, x+1, y)
	}

	if y > 0 && data[y-1][x] < 9 {
		sum += analyseBasin(data, x, y-1)
	}

	if y < uint(len(data))-1 && data[y+1][x] < 9 {
		sum += analyseBasin(data, x, y+1)
	}

	return sum
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
