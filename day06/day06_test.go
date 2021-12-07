package day06

import (
	"fmt"
	"os"
	"testing"
)

func TestCalcAmountOfLanternfishAfter18Days(t *testing.T) {
	initialDay := []uint{
		3, 4, 3, 1, 2,
	}

	got := CalcAmountOfLanternfishAfterXDays(initialDay, 18)
	if got != 26 {
		t.Errorf("Play = %d, want 26\n", got)
	}
}

func TestCalcAmountOfLanternfishAfter80DaysFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcAmountOfLanternfishAfterXDaysFile(file, 80)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestCalcAmountOfLanternfishAfter256DaysFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalcAmountOfLanternfishAfterXDaysFile(file, 256)

	fmt.Printf("challenge input data got %d\n", got)
}
