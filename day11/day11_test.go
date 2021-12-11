package day11

import (
	"fmt"
	"os"
	"testing"
)

func TestCalcEnergyLevelTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcEnergyLevelFile(file, 100)
	if got != 1656 {
		t.Errorf("Play = %d, want 1656\n", got)
	}
}

func TestCalcEnergyLevelChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcEnergyLevelFile(file, 100)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestHowMuchStepsToFullFlashTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowMuchStepsToFullFlashFile(file)
	if got != 195 {
		t.Errorf("Play = %d, want 195\n", got)
	}
}

func TestHowMuchStepsToFullFlashChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := HowMuchStepsToFullFlashFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
