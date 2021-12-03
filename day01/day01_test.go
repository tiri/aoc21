package day01

import (
	"fmt"
	"os"
	"testing"
)

func TestLargerThenPrevious(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	got := LargerThenPrevious(data)
	if got != 7 {
		t.Errorf("LargerThenPrevious = %d, want 7\n", got)
	}
}

func TestLargerThenPreviousChallengeInput(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := LargerThenPreviousFile(file)

	fmt.Printf("challenge input data got %d larger numbers then previous\n", got)
}

func TestLargerThenPreviousSlidingWindow(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	got := LargerThenPreviousSlidingWindow(data)
	if got != 5 {
		t.Errorf("LargerThenPreviousSlidingWindow = %d, want 5\n", got)
	}
}

func TestLargerThenPreviousSlidingWindowChallengeInput(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := LargerThenPreviousSlidingWindowFile(file)

	fmt.Printf("challenge input data got %d larger numbers then previous\n", got)
}

func BenchmarkLargerThenPreviousSlidingWindow(b *testing.B) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	for i := 0; i < b.N; i++ {
		LargerThenPreviousSlidingWindow(data)
	}
}
