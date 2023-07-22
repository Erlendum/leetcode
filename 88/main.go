package main

import "fmt"

func rShift(n []int, pos int, step int) {

	for i := len(n) - 1; i > pos; i-- {
		n[i] = n[i-step]
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	pn1 := 0
	pn2 := 0
	pres := 0

	cShifts := 0
	if n == 0 {
		return
	}
	if m == 0 {
		for i, el := range nums2 {
			nums1[i] = el
		}
		return
	}

	for pn1 < len(nums1) || pn2 < len(nums2) {
		if pn1 < m+cShifts && (pn2 >= len(nums2) || nums1[pn1] <= nums2[pn2]) {
			if pn1 != pres {
				rShift(nums1, pres, 1)
				cShifts++
				pn1++
			}
			nums1[pres] = nums1[pn1]
			pres++
			pn1++
		} else if pn2 < len(nums2) && (pn1 >= m+cShifts || nums2[pn2] <= nums1[pn1]) {
			rShift(nums1, pres, 1)
			cShifts++
			pn1++
			nums1[pres] = nums2[pn2]
			pres++
			pn2++
		}
	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
