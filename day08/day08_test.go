package day08

import (
	"fmt"
	"os"
	"testing"
)

func TestHowManyTimesAppearDigit1478TestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyTimesAppearDigit1478File(file)
	if got != 26 {
		t.Errorf("Play = %d, want 26\n", got)
	}
}

func TestHowManyTimesAppearDigit1478ChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowManyTimesAppearDigit1478File(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestDecodeSevenSegmentTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := DecodeSevenSegmentFile(file)
	if got != 61229 {
		t.Errorf("Play = %d, want 61229\n", got)
	}
}

func TestDecodeSevenSegmentChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := DecodeSevenSegmentFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
