package day12

import (
	"fmt"
	"os"
	"testing"
)

func TestHowManyPathInSystemTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyPathInSystemFile(file, false)
	if got != 226 {
		t.Errorf("Play = %d, want 226\n", got)
	}
}

func TestHowManyPathInSystemChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyPathInSystemFile(file, false)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestHowManyPathInSystemEnhancedTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyPathInSystemFile(file, true)
	if got != 3509 {
		t.Errorf("Play = %d, want 3509\n", got)
	}
}

func TestHowManyPathInSystemEnhancedChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyPathInSystemFile(file, true)

	fmt.Printf("challenge input data got %d\n", got)
}
