package day10

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckSyntaxTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CheckSyntaxFile(file)
	if got != 26397 {
		t.Errorf("Play = %d, want 26397\n", got)
	}
}

func TestCheckSyntaxChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CheckSyntaxFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestCompleteSyntaxTestFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CompleteSyntaxFile(file)
	if got != 288957 {
		t.Errorf("Play = %d, want 288957\n", got)
	}
}

func TestCompleteSyntaxChallengeFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := CompleteSyntaxFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
