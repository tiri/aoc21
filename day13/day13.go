package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	axis int
	fold int
}

type Point struct {
	x int
	y int
}

func HowManyDotsAfterFirstIteration(data map[Point]bool, instructions []Instruction) int {
	runInstruction(data, instructions[0])

	return len(data)
}

func RunAllInstructionsAndDisplay(data map[Point]bool, instructions []Instruction) int {

	for _, instruction := range instructions {
		runInstruction(data, instruction)
	}

	// print
	xMax, yMax := maxXY(data)
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if _, ok := data[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return len(data)
}

func runInstruction(data map[Point]bool, instruction Instruction) {
	for point, _ := range data {
		// x
		if instruction.axis == 0 && point.x > instruction.fold {
			diff := point.x - instruction.fold

			data[Point{instruction.fold - diff, point.y}] = true

			delete(data, point)
			continue
		}

		// y
		if instruction.axis == 1 && point.y > instruction.fold {
			diff := point.y - instruction.fold

			data[Point{point.x, instruction.fold - diff}] = true

			delete(data, point)
			continue
		}
	}
}

func HowManyDotsAfterFirstIterationFile(file *os.File) int {
	return HowManyDotsAfterFirstIteration(parseFile(file))
}

func RunAllInstructionsAndDisplayFile(file *os.File) int {
	return RunAllInstructionsAndDisplay(parseFile(file))
}

func parseFile(file *os.File) (map[Point]bool, []Instruction) {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	data := make(map[Point]bool)
	var instructions []Instruction

	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), ",")

		if len(splitString) < 2 {
			break
		}

		x, _ := strconv.ParseInt(splitString[0], 10, 32)
		y, _ := strconv.ParseInt(splitString[1], 10, 32)

		data[Point{int(x), int(y)}] = true
	}

	for scanner.Scan() {
		var axis int
		var folds int

		_, err := fmt.Sscanf(scanner.Text(), "fold along %c=%d", &axis, &folds)
		if err != nil {
			panic(0)
		}

		instructions = append(instructions, Instruction{axis - 120, folds})
	}

	return data, instructions
}

func maxXY(data map[Point]bool) (int, int) {
	var x int
	var y int

	for datum, _ := range data {
		if x < datum.x {
			x = datum.x
		}
		if y < datum.y {
			y = datum.y
		}
	}

	return x, y
}
