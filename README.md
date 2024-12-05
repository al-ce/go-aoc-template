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
├── setup.sh
├── go.mod
├── inputs
│   └── 1
├── internal
│   └── helpers
│       └── helpers.go
├── justfile
├── main.go
├── README.md
└── solutions
    ├── all
    │   └── all.go
    ├── day01
    │   ├── day01.go
    │   └── day01_test.go
    ├── # etc.

31 directories, 58 files
```

Download the input for day `N` (no leading zeros) in `inputs/` (please note that `inputs/` is in `.gitignore` so you don't commit the input data).

An example with [aocgofetch](https://github.com/al-ce/aocgofetch):

```bash
> aocgofetch 2024 1 > inputs/1
```

Finally, create a new git branch for the year you are solving so you'll have a clean template for next year!

## Writing a test

Take the example input from the daily instructions and replace the values for `lines`, `partOneAnswer`, and `partTwoAnswer` in the test template. For example, day 1 of 2024 would look like this:

```go
var lines = strings.Split(`3   4
4   3
2   5
1   3
3   9
3   3`, "\n")

var partOneAnswer = "11"
var partTwoAnswer = "31"
```

After writing the solution, run the test:

```bash
❯ go test go-aoc-template/solutions/day1
ok      go-aoc-template/solutions/day1  0.001s
```

This might not cover any edge cases from the actual puzzle input.

## Writing your solution

Fill out `PartOne()` and `PartTwo()` in the day's template with each part's solution. See [solutions/day1/day1.go](solutions/day1/day1.go) for reference.

Helpers can go in [internal/helpers/helpers.go](internal/helpers/helpers.go). In `main.go`, the provided `InputLines()` helper generates the lines from the input file and passes them as a string slice to `PartOne()` and `PartTwo()`.


## justfile

This is the `justfile` I use to help my workflow.

```just
set quiet

get year day:  # Write the puzzle input to the input folder.
    aocgofetch {{year}} {{day}} > inputs/{{day}}  # Use your preferred tool

setup day:  # Clean slate for a given day. Used to create templates.
    ./setup.sh {{day}}

everyday:  # Loops over setup command. Sets up blank year.
    bash -c 'for i in {1..25}; do just setup "$i"; done'

solve day part:  # Prints the solution for a given day and part.
    go run . {{day}} {{part}}

test day:  # Run the test for the given day, assuming example input and known solution.
    go test go-aoc-template/solutions/day$(printf %02d {{day}})
```
