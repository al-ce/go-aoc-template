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

func PartTwo(lines []string) string {
	return "TODO"
}
