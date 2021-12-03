package day03

import (
	"bufio"
	"os"
	"strconv"
)

func CalculateGammaEpsilon(data []uint, length int) uint {
	var gamma uint = 0
	var epsilon uint = 0

	if len(data) < 1 {
		return 0
	}

	for pos := 0; pos < length; pos++ {
		ones, zeros := 0, 0
		for _, datum := range data {
			if (datum>>pos)&1 == 1 {
				ones++
			}
		}

		zeros = len(data) - ones

		if ones > zeros {
			gamma |= 1 << pos
		} else {
			epsilon |= 1 << pos
		}
	}

	return gamma * epsilon
}

func CalculateGammaEpsilonFile(file *os.File) uint {
	return CalculateGammaEpsilon(parseFile(file))
}

func CalculateOxygenCO2(data []uint, length int) uint {
	mostCommon := data
	leastCommon := data

	if len(data) < 1 {
		return 0
	}

	for pos := length - 1; pos >= 0; pos-- {
		mostCommon = getMostOrLeastCommon(mostCommon, pos, true)
		leastCommon = getMostOrLeastCommon(leastCommon, pos, false)
	}

	return mostCommon[0] * leastCommon[0]
}

func getMostOrLeastCommon(data []uint, pos int, mostCommon bool) []uint {
	if len(data) < 2 {
		return data
	}

	var filteredData []uint
	ones, zeros := 0, 0
	for _, datum := range data {
		if (datum>>pos)&1 == 1 {
			ones++
		}
	}

	zeros = len(data) - ones

	var filterCharAtPos uint

	if mostCommon {
		if ones >= zeros {
			filterCharAtPos = 1
		} else {
			filterCharAtPos = 0
		}
	} else {
		if ones < zeros {
			filterCharAtPos = 1
		} else {
			filterCharAtPos = 0
		}
	}

	for _, datum := range data {
		if (datum>>uint(pos))&1 == filterCharAtPos {
			filteredData = append(filteredData, datum)
		}
	}
	return filteredData
}

func CalculateOxygenCO2File(file *os.File) uint {
	return CalculateOxygenCO2(parseFile(file))
}

func parseFile(file *os.File) ([]uint, int) {
	length := 0
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []uint

	for scanner.Scan() {
		num, _ := strconv.ParseUint(scanner.Text(), 2, 32)
		length = len(scanner.Text())
		data = append(data, uint(num))
	}
	return data, length
}
