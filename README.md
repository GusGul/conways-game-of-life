# Conway's Game of Life

This is an implementation of Conway's Game of Life that works with 64-bit signed integer coordinates. The program reads the initial state in Life 1.06 format from standard input, runs 10 generations of the simulation, and outputs the final state in Life 1.06 format to standard output.

## Features

- Supports 64-bit signed integer coordinates
- Uses sparse grid representation for efficient memory usage
- Implements standard Conway's Game of Life rules:
  1. Any live cell with fewer than two live neighbors dies (underpopulation)
  2. Any live cell with two or three live neighbors lives on to the next generation
  3. Any live cell with more than three live neighbors dies (overpopulation)
  4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction)
- Reads and writes in Life 1.06 format

## Building

```bash
go build
```

## Usage

The program reads from standard input and writes to standard output. Input should be in Life 1.06 format:

```bash
# Run with input from a file
./conways-game-of-life < input.life

# Or pipe input directly
echo "#Life 1.06
0 1
1 2
2 0
2 1
2 2" | ./conways-game-of-life
```

### Input Format (Life 1.06)

The input should be in Life 1.06 format:
- First line must be `#Life 1.06`
- Each subsequent line contains two space-separated integers representing x and y coordinates of a live cell
- Coordinates can be anywhere in the signed 64-bit integer range

Example input:
```
#Life 1.06
0 1
1 2
2 0
2 1
2 2
-2000000000000 -2000000000000
```

### Output Format

The output is also in Life 1.06 format, showing the state after 10 generations.

## Testing

Run the tests with:

```bash
go test ./...
```

## Implementation Details

The implementation uses a sparse grid representation with a map to store only the live cells. This makes it efficient for patterns with relatively few live cells, even when they're spread across a large coordinate space.

The `game` package provides the core functionality:
- `Grid`: Represents the game state using a sparse grid
- `NextGeneration()`: Computes the next generation according to Conway's rules
- `ReadLife106()`: Reads the initial state in Life 1.06 format
- `WriteLife106()`: Writes the final state in Life 1.06 format 