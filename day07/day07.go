package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CalcLowestFuelBurned(crabs []uint, moreExpensiveWithMoreSteps bool) uint {

	minCrab, maxCrab := getMinMax(crabs)

	lowestFuelUsage := uint(0)

	horizontalPos := minCrab
	for {
		currentFuelUsage := uint(0)
		for _, crab := range crabs {
			steps := uint(0)
			if horizontalPos > crab {
				steps += horizontalPos - crab
			} else {
				steps += crab - horizontalPos
			}

			if !moreExpensiveWithMoreSteps {
				currentFuelUsage += steps
			} else {
				currentFuelUsage += steps * (steps + 1) / 2
			}

		}

		if currentFuelUsage < lowestFuelUsage || lowestFuelUsage == 0 {
			lowestFuelUsage = currentFuelUsage
		}

		if horizontalPos >= maxCrab {
			break
		}

		horizontalPos++
	}

	return lowestFuelUsage
}

func CalcLowestFuelBurnedFile(file *os.File, moreExpensiveWithMoreSteps bool) uint {
	return CalcLowestFuelBurned(parseFile(file), moreExpensiveWithMoreSteps)
}

func parseFile(file *os.File) []uint {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var crabs []uint

	if scanner.Scan() {
		abc := strings.Split(scanner.Text(), ",")

		for _, s := range abc {
			num, _ := strconv.ParseUint(s, 10, 32)
			crabs = append(crabs, uint(num))
		}
	}
	return crabs
}

func getMinMax(slice []uint) (uint, uint) {
	var min, max uint
	for _, item := range slice {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}
	return min, max
}
