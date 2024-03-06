package main

import "fmt"

// remove all occurrences of val in nums in place
// return the number of elements in nums which are not equal to val
func removeElement(nums []int, val int) int {
	k := 0

	// loop through
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums1 := [4]int{3, 2, 2, 3}
	val := 3

	test1 := removeElement(nums1[:], val)
	fmt.Println(test1)
	// Output: 2

	nums2 := [5]int{0, 1, 2, 2, 3}
	val = 2

	test2 := removeElement(nums2[:], val)
	fmt.Println(test2)
	// Output: 2
}

// iterate through nums keeping a separate pointer to our insert location. If the value in nums isn't the val we're removing, then insert it at the insert location and increment the insert location. Since the insert location is always the next available spot, it is also keeping the length of the returned array.
