package day09

import (
	"fmt"
	"os"
	"testing"
)

func TestCalcRiskLevelsTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcRiskLevelsFile(file)
	if got != 15 {
		t.Errorf("Play = %d, want 15\n", got)
	}
}

func TestCalcRiskLevelsChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcRiskLevelsFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestCalcBasinTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcBasinFile(file)
	if got != 1134 {
		t.Errorf("Play = %d, want 1134\n", got)
	}
}

func TestCalcBasinChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcBasinFile(file)
	fmt.Printf("challenge input data got %d\n", got)
}
