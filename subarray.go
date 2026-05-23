package main

func maxSubarrayBrute(arr []int) int {
	best := arr[0] // largest sum we've found so far, anywhere

	// i is the STARTING index of the subarray we are considering
	for i := 0; i < len(arr); i++ {
		running := 0 // sum of the current subarray we're considering

		// j is the ENDINg index; it sweeps from iforward to the end
		for j := i; j < len(arr); j++ {
			running += arr[j]
			if running > best {
				best = running
			}
		}
	}
	return best
}

func kadane(arr []int) int {
	// current = best subarray sum ENDING at the current index.
	// best = best subarray sum we have seen anywhere so far
	// both start as arr[0]: the only subarray ending at (and seen by)
	// index 0 is the single element arr[0] itself.
	current := arr[0]
	best := arr[0]

	for j := 1; j < len(arr); j++ {
		// the recurrence: either start fresh at arr[j]
		// or extend the best run that ended at j-1.
		if current+arr[j] > arr[j] {
			current = current + arr[j]
		} else {
			current = arr[j]
		}

		// Has the best-ending-here run beaten our global record?
		if current > best {
			best = current
		}
	}
	return best
}

// Divide and Conquer Version
// Recursion:
// Takes advantage of the following idea:
// If we split the array into two halfs, the max subarray
// either lies completely in the left half or completely in the right half
// OR it spans the two.  The first two options are just a smaller version of the larger
// problem.
func maxSubarrayDC(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return recmax(arr, 0, len(arr)-1)
}

// recmax returns the maximum subarray sum within arr[1..u] inclusive.
// l is the lower index bound, u is the upper index bound.
func recmax(arr []int, l, u int) int {
	// base case 1: empty range (1 past u). Contributes nothing
	if l > u {
		return 0
	}

	// base case 2: a single element. The best subarray of a one-element
	// range is either that element alone, or -- if the problem allows an
	// empty subarray -- zero. The pseudocode uses max(0, A[l]) here.
	if l == u {
		return max(0, arr[l])
	}

	// Divide: find the midpoint
	m := (l + u) / 2

	// Find the best stretch reaching left from the midpoint m ---
	// we walk backward from m down to 1, accumulating into sum,
	// and remember the largest running total in lmax.
	lmax, sum := 0, 0
	for i := m; i >= 1; i-- {
		sum += arr[i]
		if sum > lmax {
			lmax = sum // best place to cut the left boundary so far
		}
	}

	// Find the best stretch reaching RIGHT from m+1 ---
	// Mirror image: walk forward from m+1 up to u .
	rmax, sum2 := 0, 0
	for i := m + 1; i <= u; i++ {
		sum2 += arr[i]
		if sum2 > rmax {
			rmax = sum2 // best place to cut the right boundary
		}
	}

	// Conquer + combine: the answer is the largest of three candidates:
	//		1. best subarray entirely in the left half: recmax(arr, 1, m)
	//		2. best subarray entirely in the right half: recmax(arr, m+1, u)
	//		3. best subarray crossing the middle:  lmax + rmax
	leftBest := recmax(arr, 1, m)
	rightBest := recmax(arr, m+1, u)
	crossBest := lmax + rmax

	return max(max(leftBest, rightBest), crossBest)

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	// pass 1: output[i] = product of everything to the left of i.
	// We record the running product before *= nums[i]
	// so nums[i] is never included in its own calculation.
	leftProduct := 1
	for i := 0; i < n; i++ {
		output[i] = leftProduct
		leftProduct *= nums[i]
	}

	// pass 2: multiply in the product of everything to the right of i.
	// sweeping right to left
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}
