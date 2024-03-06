package main

import "fmt"

// 88. Merge Sorted Array

// two integer arrays nums1 and nums2, sorted non-decreasing order,
// two integers m, n. num elements in nums1, nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// Constraints:
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -109 <= nums1[i], nums2[j] <= 109

// O(n + m)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1 2 2 3 5 6]

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1]

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1]
}

// ---

// Solution Explanation

// The merge_sorted_arr function implements an efficient algorithm to merge two sorted arrays, nums1 and nums2, directly into nums1. Here's how it works:

// Three Pointers:

// i: Points to the last element in the initialized part of nums1.
// j: Points to the last element in nums2.
// k: Points to the position in nums1 where the next merged element should be placed (starts at the end).
// Merging Logic:

// The loop continues while j >= 0 (we still have elements in nums2).
// Comparison: It compares the current elements pointed by i and j.
// Larger Element Placement:
// If nums1[i] is larger, it is placed at nums1[k] and i is decremented (we move towards the beginning of nums1).
// Otherwise, nums2[j] is placed at nums1[k] and j is decremented.
// k is decremented with each placement.
// In-Place Modification:  The key idea is that nums1 has enough space to accommodate the merged result.

// Notes

// Efficiency: The solution has a time complexity of O(n + m), making it very efficient for merging sorted arrays.
// Edge Cases: The code correctly handles edge cases where either array might be empty, as demonstrated in your main function examples.
// Clarity: The code is well-structured and readable.
// Overall, this is an excellent solution to the problem of merging sorted arrays!

// Possible Extensions

// Error Handling: Consider adding checks to ensure m and n are within valid bounds based on the lengths of the arrays, to enhance robustness.
