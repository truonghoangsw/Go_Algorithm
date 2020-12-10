package main

import (
	"fmt"
)

func merge(arr []int, pointFirst, pointMid, pointLast int) []int {
	arrLengthMergeLeft, arrLengthMergeRight := 0, 0
	arrLengthMergeLeft = pointMid - pointFirst + 1
	arrLengthMergeRight = pointLast - pointMid

	arrLeft := make([]int, arrLengthMergeLeft)
	arrRight := make([]int, arrLengthMergeRight)

	starMerge := pointFirst

	for i := 0; i < arrLengthMergeLeft; i++ {
		arrLeft[i] = arr[pointFirst+i]
	}

	for i := 0; i < arrLengthMergeRight; i++ {
		arrRight[i] = arr[i+1+pointMid]
	}
	indexLeft, indexRight := 0, 0
	for indexLeft < arrLengthMergeLeft && indexRight < arrLengthMergeRight {
		if arrLeft[indexLeft] >= arrRight[indexRight] {
			arr[starMerge] = arrRight[indexRight]
			indexRight++
		} else {
			arr[starMerge] = arrLeft[indexLeft]
			indexLeft++
		}
		starMerge++
	}
	for indexLeft < arrLengthMergeLeft {
		arr[starMerge] = arrLeft[indexLeft]
		starMerge++
		indexLeft++
	}

	for indexRight < arrLengthMergeRight {
		arr[starMerge] = arrRight[indexRight]
		indexRight++
		starMerge++
	}
	return arr
}

// Hello returns a greeting for the named person.
func mergesort(arr []int, first, last int) {
	if first >= last {
		return
	}
	mid := (first + last - 1) / 2
	mergesort(arr, first, mid)
	mergesort(arr, mid+1, last)
	arr = merge(arr, first, mid, last)
}

func main() {
	arr := []int{3, 1, 6, 34}
	mergesort(arr, 0, len(arr)-1)
	for _, value := range arr {
		fmt.Printf("Received %d ", value)
	}
}
