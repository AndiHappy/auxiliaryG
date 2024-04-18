package sort

import (
	"fmt"
	"testing"
)

func Test_FindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1}
	nums2 = []int{2, 3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	nums2 = []int{0, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	//[1,2] ,[3,4] m=2,n=2, first=m+n/2,m+n+1/2 := 2; second=m+n+2/2,m+n+3/2 :=3
	//[1,2,3],[4,5]  m=3,m=2 first=m+n/2,m+n+1/2 := 3 ;  second=m+n+2/2,m+n+3/2 :=3
	bI := (m + n + 1) / 2
	eI := (m + n + 2) / 2

	bIv := findIndexInSortedArrays(nums1, nums2, 0, 0, bI)
	eIv := findIndexInSortedArrays(nums1, nums2, 0, 0, eI)
	return (bIv + eIv) / 2.0
}

// findIndexInSortedArrays 获取index的值
// ith 必须大于1
func findIndexInSortedArrays(nums1 []int, nums2 []int, s1, s2, ith int) float64 {
	if s1 >= len(nums1) {
		return float64(nums2[s2+ith-1])
	}

	if s2 >= len(nums2) {
		return float64(nums1[s1+ith-1])
	}

	if ith == 1 {
		if nums2[s2] > nums1[s1] {
			return float64(nums1[s1])
		} else {
			return float64(nums2[s2])
		}
	}

	//how to get middle value
	mid := ith / 2
	nums1Midv := 0
	nums2Midv := 0
	if s1+mid-1 < len(nums1) {
		nums1Midv = nums1[s1+mid-1]
	} else {
		//说明中间的长度已经大于nums1的长度了，
		nums1Midv = nums1[len(nums1)-1]
		if s2+mid-1 < len(nums2) {
			nums2Midv = nums2[s2+mid-1]
		}
		if nums1Midv <= nums2Midv {
			return findIndexInSortedArrays(nums1, nums2, s1+mid, s2, ith-len(nums1))
		}
	}
	if s2+mid-1 < len(nums2) {
		nums2Midv = nums2[s2+mid-1]
	} else {
		nums2Midv = nums2[len(nums2)-1]
		if nums2Midv <= nums1Midv {
			return findIndexInSortedArrays(nums1, nums2, s1, s2+mid, ith-len(nums2))
		}
	}

	if nums1Midv < nums2Midv {
		return findIndexInSortedArrays(nums1, nums2, s1+mid, s2, ith-ith/2)
	} else {
		return findIndexInSortedArrays(nums1, nums2, s1, s2+mid, ith-ith/2)
	}
}
