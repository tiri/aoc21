package day04

import (
	"fmt"
	"os"
	"testing"
)

func TestCalculateGammaEpsilon(t *testing.T) {
	drawnNumbers := []uint{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	boards := []Board{
		{
			{{22, false}, {13, false}, {17, false}, {11, false}, {0, false}},
			{{8, false}, {2, false}, {23, false}, {4, false}, {24, false}},
			{{21, false}, {9, false}, {14, false}, {16, false}, {7, false}},
			{{6, false}, {10, false}, {3, false}, {18, false}, {5, false}},
			{{1, false}, {12, false}, {20, false}, {15, false}, {19, false}},
		},
		{
			{{3, false}, {15, false}, {0, false}, {2, false}, {22, false}},
			{{9, false}, {18, false}, {13, false}, {17, false}, {5, false}},
			{{19, false}, {8, false}, {7, false}, {25, false}, {23, false}},
			{{20, false}, {11, false}, {10, false}, {24, false}, {4, false}},
			{{14, false}, {21, false}, {16, false}, {12, false}, {6, false}},
		},
		{
			{{14, false}, {21, false}, {17, false}, {24, false}, {4, false}},
			{{10, false}, {16, false}, {15, false}, {9, false}, {19, false}},
			{{18, false}, {8, false}, {23, false}, {26, false}, {20, false}},
			{{22, false}, {11, false}, {13, false}, {6, false}, {5, false}},
			{{2, false}, {0, false}, {12, false}, {3, false}, {7, false}},
		},
	}
	got := Play(drawnNumbers, boards)
	if got != 4512 {
		t.Errorf("Play = %d, want 4512\n", got)
	}
}

func TestPlayFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := PlayFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}

func TestPlayLastFile(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	got := PlayLastFile(file)

	fmt.Printf("challenge input data got %d\n", got)
}
