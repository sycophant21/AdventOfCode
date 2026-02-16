# AdventOfCode

Solutions to [Advent of Code 2024](https://adventofcode.com/2024) programming challenges, written in Go.

## Overview

This repository contains solutions for Advent of Code 2024 puzzles. Advent of Code is an annual set of Christmas-themed programming challenges that are released daily from December 1-25. Each day has two parts (H1 and H2), where part 2 typically builds on or modifies part 1.

## Structure

```
2024/
├── Day1/
│   ├── H1/app/main.go       # Part 1 solution
│   └── H2/app/main.go       # Part 2 solution
├── Day2/
│   ├── H1/app/main.go
│   └── H2/app/main.go
...
└── Day7/
    ├── H1/app/main.go
    └── H2/app/main.go
```

Each solution reads input from a corresponding `input.txt` file in the same directory.

## Completed Days

### Day 1: Historian Hysteria
**Part 1**: Calculate total distance between sorted lists by pairing smallest numbers from each list.
**Part 2**: Calculate similarity score by counting occurrences of left list numbers in right list.

**Input**: Two columns of numbers  
**Output**: Numeric result

### Day 2: Red-Nosed Reports
**Part 1**: Count reports where numbers are strictly increasing or decreasing with differences of 1-3.
**Part 2**: Same as Part 1, but allow removing one number from invalid reports to make them valid.

**Input**: Lines of space-separated numbers  
**Output**: Count of valid reports

### Day 3: Mull It Over
**Part 1**: Parse corrupted memory for valid `mul(X,Y)` instructions and sum their products.
**Part 2**: Same as Part 1, but respect `do()` and `don't()` instructions that enable/disable multiplication.

**Input**: Single line of text with embedded instructions  
**Output**: Sum of multiplication results

### Day 4: Ceres Search
**Part 1**: Find all occurrences of "XMAS" in a character grid (horizontal, vertical, diagonal, any direction).
**Part 2**: Find X-shaped patterns where "MAS" appears on both diagonals (in any direction).

**Input**: Grid of characters  
**Output**: Count of pattern occurrences

### Day 5: Print Queue
**Part 1**: Validate page orderings against rules, sum middle page numbers of valid orderings.
**Part 2**: Fix invalid orderings to satisfy rules, sum middle page numbers of corrected orderings.

**Input**: Ordering rules (X|Y format) followed by page sequences  
**Output**: Sum of middle page numbers

### Day 6: Guard Gallivant
**Part 1**: Simulate a guard's path on a grid (turning right at obstacles), count distinct positions visited.

**Input**: Grid with starting position (^) and obstacles (#)  
**Output**: Count of visited positions

### Day 7: Bridge Repair
**Part 1**: Find equations that can be satisfied using + and * operators (left to right), sum valid results.
**Part 2**: Same as Part 1, but also allow || (concatenation) operator.

**Input**: Lines with target number and operands (e.g., "190: 10 19")  
**Output**: Sum of valid target numbers

## Running Solutions

Each day's solution is a standalone Go program:

```bash
# Run Day 1, Part 1
go run 2024/Day1/H1/app/main.go

# Run Day 3, Part 2
go run 2024/Day3/H2/app/main.go
```

**Important**: Solutions expect to be run from the repository root directory because they reference input files using relative paths (e.g., `2024/Day1/H1/app/input.txt`).

## Input Files

Each solution directory contains:
- `input.txt` - The puzzle input (typically large)
- `input_temp.txt` - Sample/test input (smaller, for development/testing)

Some solutions are configured to read from `input_temp.txt` instead of `input.txt` (e.g., Day 7 Part 1).

## Implementation Notes

- All solutions read input from text files using `bufio.Scanner`
- Most solutions process input line by line
- Solutions use Go's standard library (no external dependencies)
- Some solutions include debug output (e.g., Day 4 prints matching positions)
- Type constraints and generics are used in some helper functions

## Current State

The repository contains working solutions for Days 1-7 (Parts 1 and 2 for most days, Part 1 only for Day 6).

**Note**: This repository uses `module AdventOfCode` (from go.mod) and does not have tests. All solutions are executable programs that print results to stdout.
