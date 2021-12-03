package day02

import (
	"fmt"
	"os"
	"testing"
)

func TestNavigateSub(t *testing.T) {
	data := []Operation{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	got := NavigateSub(data)
	if got != 150 {

		t.Errorf("NavigateSub = %d, want 150\n", got)
	}
}

func TestNavigateSubFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := NavigateSubFile(file)

	fmt.Printf("challenge input data got %d larger numbers then previous\n", got)
}

func TestNavigateSubAim(t *testing.T) {
	data := []Operation{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	got := NavigateSubAim(data)
	if got != 900 {
		t.Errorf("NavigateSubAim = %d, want 900\n", got)
	}
}

func TestNavigateSubAimFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := NavigateSubAimFile(file)

	fmt.Printf("challenge input data got %d larger numbers then previous\n", got)
}

func BenchmarkNavigateSubFile(b *testing.B) {
	data := []Operation{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}

	for i := 0; i < b.N; i++ {
		NavigateSub(data)
	}
}
