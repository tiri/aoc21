package day13

import (
	"fmt"
	"os"
	"testing"
)

func TestHowManyDotsAfterFirstIterationTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyDotsAfterFirstIterationFile(file)
	if got != 17 {
		t.Errorf("Play = %d, want 17\n", got)
	}
}

func TestHowManyDotsAfterFirstIterationChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyDotsAfterFirstIterationFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestRunAllInstructionsAndDisplayChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := RunAllInstructionsAndDisplayFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
