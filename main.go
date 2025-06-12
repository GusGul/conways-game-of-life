package main

import (
	"fmt"
	"os"
	"time"

	"conways-game-of-life/game"
)

func main() {
	grid, err := game.ReadLife106(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	minX, maxX := int64(0), int64(5)
	minY, maxY := int64(0), int64(5)

	for i := 0; i < 10; i++ {
		fmt.Printf("Generation %d:\n", i)
		fmt.Print(grid.Visualize(minX, maxX, minY, maxY))
		fmt.Println()
		time.Sleep(1 * time.Second)
		grid = grid.NextGeneration()
	}

	if err := game.WriteLife106(os.Stdout, grid); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
