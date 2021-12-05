package day05

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y uint
}

type Segment struct {
	start, end Point
}

func FindLargeOpaqueClouds(segments []Segment) uint {
	// filter diagonal
	var filteredSlice []Segment
	for _, segment := range segments {
		if segment.start.x == segment.end.x || segment.start.y == segment.end.y {
			filteredSlice = append(filteredSlice, segment)
		}
	}

	// build map
	largeOpaqueCloudMap := generateMap(filteredSlice)

	// sum dangerous points
	return sumDangerousPoints(largeOpaqueCloudMap)
}

func FindLargeOpaqueCloudsFile(file *os.File) uint {
	return FindLargeOpaqueClouds(parseFile(file))
}

func FindLargeOpaqueCloudsFull(segments []Segment) uint {
	// build map
	largeOpaqueCloudMap := generateMap(segments)

	// sum dangerous points
	return sumDangerousPoints(largeOpaqueCloudMap)
}

func FindLargeOpaqueCloudsFullFile(file *os.File) uint {
	return FindLargeOpaqueCloudsFull(parseFile(file))
}

func sumDangerousPoints(largeOpaqueCloudMap map[Point]uint) uint {
	dangerousPoints := uint(0)
	for _, u := range largeOpaqueCloudMap {
		if u > 1 {
			dangerousPoints++
		}
	}
	return dangerousPoints
}

func generateMap(filteredSlice []Segment) map[Point]uint {
	largeOpaqueCloudMap := make(map[Point]uint)
	for _, segment := range filteredSlice {
		firstPointLargerThenSecondX := segment.start.x > segment.end.x
		firstPointLargerThenSecondY := segment.start.y > segment.end.y

		runnerX := segment.start.x
		runnerY := segment.start.y

		for {
			if val, ok := largeOpaqueCloudMap[Point{runnerX, runnerY}]; ok {
				largeOpaqueCloudMap[Point{runnerX, runnerY}] = val + 1
			} else {
				largeOpaqueCloudMap[Point{runnerX, runnerY}] = 1
			}

			if runnerX == segment.end.x && runnerY == segment.end.y {
				break
			}

			if runnerX != segment.end.x {
				if !firstPointLargerThenSecondX {
					runnerX++
				} else {
					runnerX--
				}
			}

			if runnerY != segment.end.y {
				if !firstPointLargerThenSecondY {
					runnerY++
				} else {
					runnerY--
				}
			}
		}
	}
	return largeOpaqueCloudMap
}

func parseFile(file *os.File) []Segment {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var segments []Segment

	for scanner.Scan() {
		var x1, y1, x2, y2 uint
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(0)
		}

		segments = append(segments, Segment{Point{x1, y1}, Point{x2, y2}})

	}
	return segments
}
