package day10

import (
	"bufio"
	"os"
	"sort"
)

var openingBrackets = map[int32]int32{
	'{': '}',
	'(': ')',
	'[': ']',
	'<': '>',
}

var closingBrackets = map[int32]int32{
	'}': '{',
	')': '(',
	']': '[',
	'>': '<',
}

func CheckSyntax(data []string) uint {
	sum := uint(0)

	for _, row := range data {
		sum += checkSyntaxRow(row)
	}

	return sum
}

func checkSyntaxRow(row string) uint {
	var stack []int32
	for _, char := range row {
		switch char {
		case '{', '(', '[', '<':
			stack = append(stack, char)
		case '}', ')', ']', '>':
			if stack == nil || len(stack) <= 0 {
				return 0
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if openingBrackets[last] == char {
				continue
			}

			switch char {
			case ')':
				return 3
			case ']':
				return 57
			case '}':
				return 1197
			case '>':
				return 25137
			}

		}
	}
	return 0
}

func CheckSyntaxFile(file *os.File) uint {
	return CheckSyntax(parseFile(file))
}

func CompleteSyntax(data []string) uint {
	var incomplete []int

	for _, row := range data {

		var stack []int32
		for _, char := range row {
			if _, ok := openingBrackets[char]; ok {
				stack = append(stack, char)
			} else if val, ok := closingBrackets[char]; ok {
				if stack == nil || len(stack) <= 0 {
					stack = nil
					break
				}
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if val != last {
					stack = nil
					break
				}
			}
		}

		if len(stack) > 0 {
			sum := 0
			for i := len(stack) - 1; i >= 0; i-- {
				sum *= 5

				switch openingBrackets[stack[i]] {
				case ')':
					sum += 1
				case ']':
					sum += 2
				case '}':
					sum += 3
				case '>':
					sum += 4
				}
			}

			incomplete = append(incomplete, sum)
		}
	}

	sort.Ints(incomplete)

	if incomplete != nil && len(incomplete) > 0 {
		return uint(incomplete[len(incomplete)/2])
	}

	return 0
}

func CompleteSyntaxFile(file *os.File) uint {
	return CompleteSyntax(parseFile(file))
}

func parseFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
