package day05

import (
	"fmt"
	"os"
	"testing"
)

func TestFindLargeOpaqueClouds(t *testing.T) {
	segments := []Segment{
		{Point{0, 9}, Point{5, 9}},
		{Point{8, 0}, Point{0, 8}},
		{Point{9, 4}, Point{3, 4}},
		{Point{2, 2}, Point{2, 1}},
		{Point{7, 0}, Point{7, 4}},
		{Point{6, 4}, Point{2, 0}},
		{Point{0, 9}, Point{2, 9}},
		{Point{3, 4}, Point{1, 4}},
		{Point{0, 0}, Point{8, 8}},
		{Point{5, 5}, Point{8, 2}},
	}

	got := FindLargeOpaqueClouds(segments)
	if got != 5 {
		t.Errorf("Play = %d, want 5\n", got)
	}
}

func TestFindLargeOpaqueCloudsFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := FindLargeOpaqueCloudsFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestFindLargeOpaqueCloudsFullFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := FindLargeOpaqueCloudsFullFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
