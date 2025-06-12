package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GusGul/conways-game-of-life/game"
)

func main() {
	grid, err := game.ReadLife106(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	minX, maxX := int64(-7), int64(7)
	minY, maxY := int64(-7), int64(7)

	maxGens := 10

	for i := 0; i < maxGens; i++ {
		fmt.Printf("Generation %d:\n", i)
		fmt.Print(grid.Visualize(minX, maxX, minY, maxY))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
		grid = grid.NextGeneration()
	}

	fmt.Printf("Generation %d:\n", maxGens)
	fmt.Print(grid.Visualize(minX, maxX, minY, maxY))
	fmt.Println()

	if err := game.WriteLife106(os.Stdout, grid); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
