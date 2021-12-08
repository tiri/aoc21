package day08

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type WireData struct {
	signalPattern []string
	outputValue   []string
}

/*
unique length

0 = 6
1 = 2 <-
2 = 5
3 = 5
4 = 4 <-
5 = 5
6 = 6
7 = 3 <-
8 = 7 <-
9 = 6
*/

func HowManyTimesAppearDigit1478(data []WireData) uint {
	sum := uint(0)

	for _, valueRow := range data {
		for _, value := range valueRow.outputValue {
			switch len(value) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

func HowManyTimesAppearDigit1478File(file *os.File) uint {
	return HowManyTimesAppearDigit1478(parseFile(file))
}

func DecodeSevenSegment(data []WireData) uint {
	sum := uint(0)

	for _, datum := range data {
		pattern := analysePattern(datum.signalPattern)
		number := decodeNumber(datum.outputValue, pattern)
		sum += number
	}

	return sum
}

func analysePattern(signalPattern []string) [10]string {
	var pattern [10]string

	// = 1,4,7,8
	var filteredSignalPatternList []string
	for _, s := range signalPattern {
		switch len(s) {
		case 2:
			pattern[1] = s
		case 4:
			pattern[4] = s
		case 3:
			pattern[7] = s
		case 7:
			pattern[8] = s
		default:
			filteredSignalPatternList = append(filteredSignalPatternList, s)
		}
	}
	signalPattern = filteredSignalPatternList

	// 0 -> 6, 2 -> 5, 3 -> 5, 5 -> 5, 6 -> 6, 9 -> 6

	// = 9
	fourPlusSeven := addPattern(pattern[4], pattern[7])
	for _, s := range signalPattern {
		if len(diffChars(s, fourPlusSeven)) == 1 {
			pattern[9] = s
			break
		}
	}

	// 0 -> 6, 2 -> 5, 3 -> 5, 5 -> 5, 6 -> 6

	// = 0,6
	filteredSignalPatternList = nil
	for _, signal := range signalPattern {
		if len(signal) == 6 {
			signalSubOne := subPattern(signal, pattern[1])
			switch len(signalSubOne) {
			case 5:
				pattern[6] = signal
			case 4:
				pattern[0] = signal
			}
		} else {
			filteredSignalPatternList = append(filteredSignalPatternList, signal)
		}
	}
	signalPattern = filteredSignalPatternList

	// 2 -> 5, 3 -> 5, 5 -> 5

	// = 3
	filteredSignalPatternList = nil
	for _, signal := range signalPattern {
		signalSubOne := subPattern(signal, pattern[1])
		if len(signalSubOne) == 3 {
			pattern[3] = signal
		} else {
			filteredSignalPatternList = append(filteredSignalPatternList, signal)
		}
	}
	signalPattern = filteredSignalPatternList

	// 2 -> 5, 5 -> 5

	// = 2, 5
	for _, signal := range signalPattern {
		signalSubSix := subPattern(signal, pattern[6])
		if len(signalSubSix) == 0 {
			pattern[5] = signal
		} else {
			pattern[2] = signal
		}
	}
	return pattern
}

func decodeNumber(outputValue []string, pattern [10]string) uint {
	number := uint(0)

	for i, value := range outputValue {
		for x, patternI := range pattern {
			if len(diffChars(value, patternI)) == 0 {
				number += uint(x * int(math.Pow10(len(outputValue)-1-i)))
			}
		}
	}

	return number
}

func DecodeSevenSegmentFile(file *os.File) uint {
	return DecodeSevenSegment(parseFile(file))
}

func parseFile(file *os.File) []WireData {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var data []WireData

	for scanner.Scan() {
		splitRow := strings.Split(scanner.Text(), " | ")

		var signalPattern []string
		var outputValue []string

		for _, element := range strings.Fields(splitRow[0]) {
			signalPattern = append(signalPattern, element)
		}

		for _, element := range strings.Fields(splitRow[1]) {
			outputValue = append(outputValue, element)
		}

		data = append(data, WireData{
			signalPattern: signalPattern,
			outputValue:   outputValue,
		})
	}
	return data
}

func addPattern(patternA string, patternB string) string {
	for _, char := range patternB {
		if !strings.ContainsRune(patternA, char) {
			patternA += string(char)
		}
	}

	return patternA
}

func subPattern(patternA string, patternB string) string {
	var pattern string

	for _, char := range patternA {
		if !strings.ContainsRune(patternB, char) {
			pattern += string(char)
		}
	}

	return pattern
}

func diffChars(stringA string, stringB string) string {
	var diff string
	diff += subPattern(stringA, stringB)
	diff += subPattern(stringB, stringA)
	return diff
}
