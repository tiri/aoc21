package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CalcAmountOfLanternfishAfterXDays(initialState []uint, xDays uint) uint {

	// group
	var groupedInitialState [9]uint
	for _, lanternfish := range initialState {
		groupedInitialState[lanternfish]++
	}

	// calc
	for day := 0; day < int(xDays); day++ {
		zeroState := groupedInitialState[0]
		copy(groupedInitialState[:], groupedInitialState[1:])
		groupedInitialState[8] = zeroState
		groupedInitialState[6] += zeroState
	}

	// sum up
	sum := uint(0)
	for _, value := range groupedInitialState {
		sum += value
	}

	return sum
}

func CalcAmountOfLanternfishAfterXDaysFile(file *os.File, xDays uint) uint {
	return CalcAmountOfLanternfishAfterXDays(parseFile(file), xDays)
}

func parseFile(file *os.File) []uint {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var initialDay []uint

	if scanner.Scan() {
		abc := strings.Split(scanner.Text(), ",")

		for _, s := range abc {
			num, _ := strconv.ParseUint(s, 10, 32)
			initialDay = append(initialDay, uint(num))
		}
	}
	return initialDay
}
