package game

import (
	"strings"
	"testing"
)

func TestReadLife106(t *testing.T) {
	input := `#Life 1.06
0 1
1 2
2 0
2 1
2 2`

	grid, err := ReadLife106(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to read Life 1.06 format: %v", err)
	}

	expectedPoints := []Point{
		{0, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	for _, p := range expectedPoints {
		if !grid.IsAlive(p.X, p.Y) {
			t.Errorf("Expected point (%d, %d) to be alive", p.X, p.Y)
		}
	}
}

func TestNextGeneration(t *testing.T) {
	grid := NewGrid()
	grid.Set(0, 1)
	grid.Set(1, 2)
	grid.Set(2, 0)
	grid.Set(2, 1)
	grid.Set(2, 2)

	nextGrid := grid.NextGeneration()

	expectedAlive := map[Point]bool{
		{1, 0}: true,
		{2, 1}: true,
		{2, 2}: true,
		{1, 2}: true,
		{3, 1}: true,
	}

	for p := range nextGrid.Cells {
		if !expectedAlive[p] {
			t.Errorf("Unexpected live cell at (%d, %d)", p.X, p.Y)
		}
	}

	for p := range expectedAlive {
		if !nextGrid.IsAlive(p.X, p.Y) {
			t.Errorf("Expected cell at (%d, %d) to be alive", p.X, p.Y)
		}
	}
}

func TestLargeCoordinates(t *testing.T) {
	input := `#Life 1.06
-2000000000000 -2000000000000
-2000000000001 -2000000000001
-2000000000000 -2000000000001`

	grid, err := ReadLife106(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to read Life 1.06 format with large coordinates: %v", err)
	}

	expectedPoints := []Point{
		{-2000000000000, -2000000000000},
		{-2000000000001, -2000000000001},
		{-2000000000000, -2000000000001},
	}

	for _, p := range expectedPoints {
		if !grid.IsAlive(p.X, p.Y) {
			t.Errorf("Expected point (%d, %d) to be alive", p.X, p.Y)
		}
	}

	nextGen := grid.NextGeneration()

	expectedNextGen := map[Point]bool{
		{-2000000000000, -2000000000000}: true,
		{-2000000000001, -2000000000001}: true,
		{-2000000000000, -2000000000001}: true,
		{-2000000000001, -2000000000000}: true,
	}

	for p := range nextGen.Cells {
		if !expectedNextGen[p] {
			t.Errorf("Unexpected live cell at (%d, %d)", p.X, p.Y)
		}
	}

	for p := range expectedNextGen {
		if !nextGen.IsAlive(p.X, p.Y) {
			t.Errorf("Expected cell at (%d, %d) to be alive", p.X, p.Y)
		}
	}
}
