package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Action string
	Amount int
}

func NavigateSub(operations []Operation) int {
	depth, horizontal := 0, 0
	for _, operation := range operations {
		switch operation.Action {
		case "forward":
			horizontal += operation.Amount
		case "down":
			depth += operation.Amount
		case "up":
			depth -= operation.Amount
		}
	}

	return depth * horizontal
}

func NavigateSubFile(file *os.File) int {
	return NavigateSub(parseFile(file))
}

func NavigateSubAim(operations []Operation) int {
	depth, aim, horizontal := 0, 0, 0
	for _, operation := range operations {
		switch operation.Action {
		case "forward":
			horizontal += operation.Amount
			depth += aim * operation.Amount
		case "down":
			aim += operation.Amount
		case "up":
			aim -= operation.Amount
		}
	}

	return depth * horizontal
}

func NavigateSubAimFile(file *os.File) int {
	return NavigateSubAim(parseFile(file))
}

func parseFile(file *os.File) []Operation {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []Operation

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if len(split) < 2 {
			break
		}
		amount, _ := strconv.Atoi(split[1])
		data = append(data, Operation{
			split[0],
			amount,
		})
	}
	return data
}
