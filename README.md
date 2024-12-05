# Advent of Code - Go Project Template

A template structure for writing your [Advent of Code](https://adventofcode.com/) puzzle solvers in Go and printing the solution to stdout.

```bash
> go run . 1 1  # prints solution for day 1 part 1
```

Clone the `main` branch of this repo only (so you don't get my attempts):

```bash
git clone --single-branch --branch main https://github.com/al-ce/aoc-go-template.git
```

## Project structure

```bash
❯ tree
.
├── go.mod
├── inputs
│   └── 1
├── internal
│   └── helpers
│       └── helpers.go
├── main.go
├── README.md
└── solutions
    └── day1
        └── day1.go
```

Download the input for day `N` in `inputs/`. (Please note that `inputs/` is in `.gitignore` so you don't commit your input data.)

An example with [aocgofetch](https://github.com/al-ce/aocgofetch):

```bash
> aocgofetch 2024 1 > ./inputs/1
```

Then create the file for the day in `solutions/` in the format `dayN/dayN.go` Day 1 is already created.

Finally, create a new git branch for the year you are solving so you'll have a clean template for next year!

## Writing your solution

Each day's file should have a `Part1()` and `Part2()` function that returns that part's solution as a string. See [solutions/day1/day1.go] for reference.

Helpers can go in [internal/helpers/helpers.go]. `InputLines` will return a list of every line in the input file for the given day.

```go
	lines, err := h.InputLines(1, 1)
```
