package game

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int64
}

type Grid struct {
	Cells map[Point]bool
}

func NewGrid() *Grid {
	return &Grid{
		Cells: make(map[Point]bool),
	}
}

func (g *Grid) Set(x, y int64) {
	g.Cells[Point{X: x, Y: y}] = true
}

func (g *Grid) IsAlive(x, y int64) bool {
	return g.Cells[Point{X: x, Y: y}]
}

func (g *Grid) CountLiveNeighbors(x, y int64) int {
	count := 0
	for dx := int64(-1); dx <= 1; dx++ {
		for dy := int64(-1); dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) NextGeneration() *Grid {
	cellsToCheck := make(map[Point]bool)

	for p := range g.Cells {
		cellsToCheck[p] = true
		for dx := int64(-1); dx <= 1; dx++ {
			for dy := int64(-1); dy <= 1; dy++ {
				neighbor := Point{X: p.X + dx, Y: p.Y + dy}
				cellsToCheck[neighbor] = true
			}
		}
	}

	nextGen := NewGrid()

	for p := range cellsToCheck {
		neighbors := g.CountLiveNeighbors(p.X, p.Y)
		isAlive := g.IsAlive(p.X, p.Y)

		if isAlive {
			if neighbors == 2 || neighbors == 3 {
				nextGen.Set(p.X, p.Y)
			}
		} else {
			if neighbors == 3 {
				nextGen.Set(p.X, p.Y)
			}
		}
	}

	return nextGen
}

func ReadLife106(r io.Reader) (*Grid, error) {
	grid := NewGrid()
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return nil, fmt.Errorf("empty input")
	}
	if header := scanner.Text(); header != "#Life 1.06" {
		return nil, fmt.Errorf("invalid format: expected '#Life 1.06', got %s", header)
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		coords := strings.Fields(line)
		if len(coords) != 2 {
			return nil, fmt.Errorf("invalid coordinate line: %s", line)
		}

		x, err := strconv.ParseInt(coords[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid x coordinate: %s", coords[0])
		}

		y, err := strconv.ParseInt(coords[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid y coordinate: %s", coords[1])
		}

		grid.Set(x, y)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func WriteLife106(w io.Writer, grid *Grid) error {
	if _, err := fmt.Fprintln(w, "#Life 1.06"); err != nil {
		return err
	}

	points := make([]Point, 0, len(grid.Cells))
	for p := range grid.Cells {
		points = append(points, p)
	}

	for _, p := range points {
		if _, err := fmt.Fprintf(w, "%d %d\n", p.X, p.Y); err != nil {
			return err
		}
	}

	return nil
}

func (g *Grid) Visualize(minX, maxX, minY, maxY int64) string {
	var sb strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if g.IsAlive(x, y) {
				sb.WriteRune('O')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
