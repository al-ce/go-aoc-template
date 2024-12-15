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
├── inputs
│   ├── 1
├── internal
│   ├── helpers
│   │   └── helpers.go
│   └── imports
│       └── imports.go
└── solutions
    ├── day01
    │   ├── day01.go
    │   └── test01
    │       └── test01.go
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

Fill out `PartOne()` and `PartTwo()` in the day's template with each part's solution. See [solutions/day1/day1.go](solutions/day01/day01.go) for reference.

Helpers can go in [internal/helpers/helpers.go](internal/helpers/helpers.go). In `main.go`, the provided `InputLines()` helper generates the lines from the input file and passes them as a string slice to `PartOne()` and `PartTwo()`.


## justfile

This is the `justfile` I use to help my workflow.

```just
set quiet
releasepath := './bin/aocgosolutions'
_dayf := "'$(printf %02d)'"

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

solve day:
    mkdir -p bin && go build -o {{releasepath}} go-aoc-template && {{releasepath}} {{day}}

test day:
    go run $(printf "./solutions/day%02d/test%02d/test%02d.go" {{day}} {{day}} {{day}})

setup day:  # Create template for a specific day (!overwrites current file!)
    ./setup.sh {{day}}

everyday:  # Create template for every day (!overwrites all files!)
    bash -c 'for i in {1..25}; do just setup "$i"; done'
```
