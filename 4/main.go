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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	nums1 = append(nums1, nums2...)
	merge(nums1, l1, nums2, len(nums2))
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}

	return (float64(nums1[len(nums1)/2-1]) + float64(nums1[len(nums1)/2])) / 2.0
}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
