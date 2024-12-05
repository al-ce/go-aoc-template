package imports

import (
	"go-aoc-template/solutions/day01"
	_ "go-aoc-template/solutions/day01"
	"go-aoc-template/solutions/day02"
	_ "go-aoc-template/solutions/day02"
	"go-aoc-template/solutions/day03"
	_ "go-aoc-template/solutions/day03"
	"go-aoc-template/solutions/day04"
	_ "go-aoc-template/solutions/day04"
	"go-aoc-template/solutions/day05"
	_ "go-aoc-template/solutions/day05"
	"go-aoc-template/solutions/day06"
	_ "go-aoc-template/solutions/day06"
	"go-aoc-template/solutions/day07"
	_ "go-aoc-template/solutions/day07"
	"go-aoc-template/solutions/day08"
	_ "go-aoc-template/solutions/day08"
	"go-aoc-template/solutions/day09"
	_ "go-aoc-template/solutions/day09"
	"go-aoc-template/solutions/day10"
	_ "go-aoc-template/solutions/day10"
	"go-aoc-template/solutions/day11"
	_ "go-aoc-template/solutions/day11"
	"go-aoc-template/solutions/day12"
	_ "go-aoc-template/solutions/day12"
	"go-aoc-template/solutions/day13"
	_ "go-aoc-template/solutions/day13"
	"go-aoc-template/solutions/day14"
	_ "go-aoc-template/solutions/day14"
	"go-aoc-template/solutions/day15"
	_ "go-aoc-template/solutions/day15"
	"go-aoc-template/solutions/day16"
	_ "go-aoc-template/solutions/day16"
	"go-aoc-template/solutions/day17"
	_ "go-aoc-template/solutions/day17"
	"go-aoc-template/solutions/day18"
	_ "go-aoc-template/solutions/day18"
	"go-aoc-template/solutions/day19"
	_ "go-aoc-template/solutions/day19"
	"go-aoc-template/solutions/day20"
	_ "go-aoc-template/solutions/day20"
	"go-aoc-template/solutions/day21"
	_ "go-aoc-template/solutions/day21"
	"go-aoc-template/solutions/day22"
	_ "go-aoc-template/solutions/day22"
	"go-aoc-template/solutions/day23"
	_ "go-aoc-template/solutions/day23"
	"go-aoc-template/solutions/day24"
	_ "go-aoc-template/solutions/day24"
	"go-aoc-template/solutions/day25"
	_ "go-aoc-template/solutions/day25"
)

type InputCoords struct {
	Day  int
	Part int
}
type SolutionFunc func([]string) string

var Solutions = map[InputCoords]SolutionFunc{
	{1, 1}:  day01.PartOne,
	{1, 2}:  day01.PartTwo,
	{2, 1}:  day02.PartOne,
	{2, 2}:  day02.PartTwo,
	{3, 1}:  day03.PartOne,
	{3, 2}:  day03.PartTwo,
	{4, 1}:  day04.PartOne,
	{4, 2}:  day04.PartTwo,
	{5, 1}:  day05.PartOne,
	{5, 2}:  day05.PartTwo,
	{6, 1}:  day06.PartOne,
	{6, 2}:  day06.PartTwo,
	{7, 1}:  day07.PartOne,
	{7, 2}:  day07.PartTwo,
	{8, 1}:  day08.PartOne,
	{8, 2}:  day08.PartTwo,
	{9, 1}:  day09.PartOne,
	{9, 2}:  day09.PartTwo,
	{10, 1}: day10.PartOne,
	{10, 2}: day10.PartTwo,
	{11, 1}: day11.PartOne,
	{11, 2}: day11.PartTwo,
	{12, 1}: day12.PartOne,
	{12, 2}: day12.PartTwo,
	{13, 1}: day13.PartOne,
	{13, 2}: day13.PartTwo,
	{14, 1}: day14.PartOne,
	{14, 2}: day14.PartTwo,
	{15, 1}: day15.PartOne,
	{15, 2}: day15.PartTwo,
	{16, 1}: day16.PartOne,
	{16, 2}: day16.PartTwo,
	{17, 1}: day17.PartOne,
	{17, 2}: day17.PartTwo,
	{18, 1}: day18.PartOne,
	{18, 2}: day18.PartTwo,
	{19, 1}: day19.PartOne,
	{19, 2}: day19.PartTwo,
	{20, 1}: day20.PartOne,
	{20, 2}: day20.PartTwo,
	{21, 1}: day21.PartOne,
	{21, 2}: day21.PartTwo,
	{22, 1}: day22.PartOne,
	{22, 2}: day22.PartTwo,
	{23, 1}: day23.PartOne,
	{23, 2}: day23.PartTwo,
	{24, 1}: day24.PartOne,
	{24, 2}: day24.PartTwo,
	{25, 1}: day25.PartOne,
	{25, 2}: day25.PartTwo,
}
