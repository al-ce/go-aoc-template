package day09

import "fmt"

func PartOne(lines []string) string {
	diskMap := lines[0]
	checkSum := 0
	blockPosition := 0
	// Cursors tracking positions of left cursor's block (file or free space) and
	// right cursor's file
	left, right := 0, len(diskMap)-1
	// Size of rightmost file
	rightSize := int(diskMap[right] - '0')

	for left < right {
		leftSize := int(diskMap[left] - '0') // Size of block under left cursor
		if left%2 == 0 {                     // Even left cursor indicates file
			id := left / 2
			for leftSize > 0 {
				checkSum += id * blockPosition
				blockPosition++
				leftSize--
			}
		} else { // Odd left cursor indicates free space
			for leftSize > 0 && left <= right {
				id := right / 2
				checkSum += id * blockPosition
				blockPosition++
				leftSize--
				rightSize--
				if rightSize == 0 {
					right -= 2
					rightSize = int(diskMap[right] - '0')
				}
			}
		}
		left++
	}

	// Add any remaining rightmost file blocks we couldn't move
	id := right / 2
	for rightSize > 0 {
		checkSum += id * blockPosition
		rightSize--
	}

	return fmt.Sprintf("%d", checkSum)
}

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

func getFileSize(blockMap []int, idx int) int {
	size := 0
	id := blockMap[idx]
	for idx >= 0 && blockMap[idx] == id {
		size++
		idx--
	}
	return size
}

func getFreeSpace(blockMap []int, idx int) int {
	size := 0
	for blockMap[idx] == -1 {
		size++
		idx++
	}
	return size
}

func findNextFile(blockMap []int, idx int) int {
	for idx >= 0 && blockMap[idx] == -1 {
		idx--
	}
	return idx
}

// Move file blocks into leftmost free space without fragmentation
func moveFileBlocks(blockMap []int) {
	fileStart := findNextFile(blockMap, len(blockMap)-1)

	for fileStart > 0 { // Check files from right to left
		// Look for sufficient free space from start of disk up to file start
		freeStart := 0
		freeSpace := 0
		fileSize := getFileSize(blockMap, fileStart)
		for freeStart < fileStart && freeSpace < fileSize {
			freeStart++
			freeSpace = getFreeSpace(blockMap, freeStart)
		}

		if freeStart != fileStart { // Found sufficient free space
			for i := range fileSize {
				blockMap[freeStart+i], blockMap[fileStart-i] = blockMap[fileStart-i], blockMap[freeStart+i]
			}
		}

		// Move file cursor to the next file
		fileStart = findNextFile(blockMap, fileStart-fileSize)
	}
}

func PartTwo(lines []string) string {
	blockMap := createBlockMap(lines[0])
	moveFileBlocks(blockMap)

	sum := 0
	for i, id := range blockMap {
		if id != -1 {
			sum += id * i
		}
	}

	return fmt.Sprintf("%d", sum)
}
