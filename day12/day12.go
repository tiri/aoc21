package day12

import (
	"bufio"
	"os"
	"strings"
)

func HowManyPathInSystem(data map[string][]string, enhanced bool) uint {
	var paths [][]string

	for _, nextCave := range data["start"] {
		paths = append(paths, followPath(data, []string{"start", nextCave}, enhanced)...)
	}

	return uint(len(paths))
}

func followPath(data map[string][]string, currentPath []string, enhanced bool) [][]string {
	var paths [][]string

	lastElement := currentPath[len(currentPath)-1]

	// check if end
	if lastElement == "end" {
		return [][]string{currentPath}
	}

	for _, nextCave := range data[lastElement] {
		if !enhanced {
			// check if small char already inside
			if nextCave == strings.ToLower(nextCave) {
				if isStringInSlice(nextCave, currentPath) {
					continue
				}
			}
		} else {
			// there should be at most twice visits of small caves
			if nextCave == strings.ToLower(nextCave) {
				if isStringInSlice(nextCave, currentPath) {
					if nextCave == "start" || nextCave == "end" {
						continue
					}
					if areThereSmallStringTwiceInSlice(currentPath) {
						continue
					}
				}
			}
		}

		paths = append(paths, followPath(data, append(currentPath, nextCave), enhanced)...)
	}

	return paths
}

func HowManyPathInSystemFile(file *os.File, enhanced bool) uint {
	return HowManyPathInSystem(parseFile(file), enhanced)
}

func parseFile(file *os.File) map[string][]string {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	data := make(map[string][]string)

	for scanner.Scan() {
		var p1, p2 string

		splitString := strings.Split(scanner.Text(), "-")
		p1 = splitString[0]
		p2 = splitString[1]

		if _, ok := data[p1]; ok {
			data[p1] = append(data[p1], p2)
		} else {
			data[p1] = []string{p2}
		}

		if _, ok := data[p2]; ok {
			data[p2] = append(data[p2], p1)
		} else {
			data[p2] = []string{p1}
		}
	}

	return data
}

func isStringInSlice(string string, slice []string) bool {
	for _, item := range slice {
		if item == string {
			return true
		}
	}
	return false
}

func areThereSmallStringTwiceInSlice(slice []string) bool {
	for _, item := range slice {
		if item != strings.ToLower(item) {
			continue
		}

		alreadyFound := false
		for _, compareItem := range slice {
			if item == compareItem {
				if alreadyFound {
					return true
				}
				alreadyFound = true
			}

		}
	}
	return false
}
