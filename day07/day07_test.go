package day06

import (
	"fmt"
	"os"
	"testing"
)

func TestCalcLowestFuelBurned(t *testing.T) {
	crabs := []uint{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	got := CalcLowestFuelBurned(crabs)
	if got != 37 {
		t.Errorf("Play = %d, want 37\n", got)
	}
}

func TestCalcLowestFuelBurnedFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcLowestFuelBurnedFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestCalcLowestFuelBurnedExpensive(t *testing.T) {
	crabs := []uint{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	got := CalcLowestFuelBurnedExpensive(crabs)
	if got != 168 {
		t.Errorf("Play = %d, want 168\n", got)
	}
}

func TestCalcLowestFuelBurnedExpensiveFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcLowestFuelBurnedExpensiveFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
