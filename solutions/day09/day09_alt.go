// Alternate solution for day9
// up to 10x slower, but easier to visualize

package day09

import (
	"fmt"
)

// Create representation of file blocks and free space e.g. "0..111....22222"
// (where '.' is a -1)
func createBlockMap(diskMap string) []int {
	blockMap := []int{}
	var block int
	id := 0

	for idx, digit := range diskMap {
		size := int(digit - '0')
		if idx%2 == 0 {
			id = idx / 2
			block = id
		} else {
			block = -1
		}

		for i := 0; i < size; i++ {
			blockMap = append(blockMap, block)
		}
	}
	return blockMap
}

// Move file blocks into leftmost free space
func fragmentFileBlocks(blockMap []int) {
	i, j := 0, len(blockMap)-1
	for i < j {
		if blockMap[i] == -1 {
			blockMap[i], blockMap[j] = blockMap[j], blockMap[i]
			for blockMap[j] == -1 {
				j--
			}
		}
		i++
	}
}

func PartOneAlt(lines []string) string {
	blockMap := createBlockMap(lines[0])
	fragmentFileBlocks(blockMap)

	sum := 0
	i := 0
	for blockMap[i] != -1 {
		sum += blockMap[i] * i
		i++
	}

	return fmt.Sprintf("%d", sum)
}
func PartTwoAlt(lines []string) string {
	return "TODO"
}

