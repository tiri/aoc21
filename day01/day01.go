package day01

import (
	"bufio"
	"os"
	"strconv"
)

func LargerThenPrevious(list []int) int {
	largerThenPrevious := 0

	if len(list) < 1 {
		return 0
	}

	for i := 1; i < len(list); i++ {
		if list[i-1] < list[i] {
			largerThenPrevious++
		}
	}

	return largerThenPrevious
}

func LargerThenPreviousSlidingWindow(list []int) int {
	const window = 3

	largerThenPrevious := 0

	if len(list) < 4 {
		return 0
	}

	previous := sumArray(list[0:window])

	for i := window; i < len(list); i++ {
		current := sumArray(list[i-window+1 : i+1])
		if previous < current {
			largerThenPrevious++
		}
		previous = current
	}

	return largerThenPrevious
}

func LargerThenPreviousFile(file *os.File) int {
	return LargerThenPrevious(parseFile(file))
}

func LargerThenPreviousSlidingWindowFile(file *os.File) int {
	return LargerThenPreviousSlidingWindow(parseFile(file))
}

func parseFile(file *os.File) []int {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}
	return data
}

func sumArray(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
