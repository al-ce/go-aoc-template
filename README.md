# Advent of Code - Go Project Template

A template structure for writing your [Advent of Code](https://adventofcode.com/) puzzle solvers in Go and printing the solution to stdout.

```bash
> go run . 1 1  # prints solution for day 1 part 1
```

Clone the `main` branch only (so you don't get my own branches):

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

Download the input for day `N` in `inputs/` (please note that `inputs/` is in `.gitignore` so you don't commit the input data).

An example with [aocgofetch](https://github.com/al-ce/aocgofetch):

```bash
> aocgofetch 2024 1 > inputs/1
```

Then create the file for the day in `solutions/` in the format `dayN/dayN.go` Day 1 is already created. Update the [solutions map](https://github.com/al-ce/go-aoc-template/blob/79976fc327b31d0cbd02c81e87def446bd1a17a6/main.go#L19-L23) in `main.go` as you add more solution files.

Finally, create a new git branch for the year you are solving so you'll have a clean template for next year!

## Writing your solution

Each day's file should have a `Part1()` and `Part2()` function that returns that part's solution as a string. See [solutions/day1/day1.go](solutions/day1/day1.go) for reference.

Helpers can go in [internal/helpers/helpers.go](internal/helpers/helpers.go). The provided `InputLines` will return a list of every line in the input file for the given day.

```go
	lines, err := h.InputLines(1, 1)
```
