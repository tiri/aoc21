package day03

import (
	"fmt"
	"os"
	"testing"
)

func TestCalculateGammaEpsilon(t *testing.T) {
	data := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}
	got := CalculateGammaEpsilon(data, 5)
	if got != 198 {
		t.Errorf("CalculateGammaEpsilon = %d, want 198\n", got)
	}
}

func TestCalculateGammaEpsilonFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalculateGammaEpsilonFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestCalculateOxygenCO2(t *testing.T) {
	data := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}
	got := CalculateOxygenCO2(data, 5)
	if got != 230 {
		t.Errorf("CalculateOxygenCO2 = %d, want 230\n", got)
	}
}

func TestCalculateOxygenCO2File(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CalculateOxygenCO2File(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func BenchmarkNavigateSubFile(b *testing.B) {
	data := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	for i := 0; i < b.N; i++ {
		CalculateOxygenCO2(data, 5)
	}
}
