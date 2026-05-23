package main

import "fmt"

func main() {
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted slice:", unsorted)

	bubbleSort(unsorted)

	fmt.Println("Sorted slice:", unsorted)

	alreadySorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	bubbleSort(unsorted)

	fmt.Println("Sorted:", alreadySorted)

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	largestSubarray := kadane(nums)

	fmt.Println("kadane: ", largestSubarray)
}
