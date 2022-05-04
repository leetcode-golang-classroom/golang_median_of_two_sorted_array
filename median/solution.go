package median

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedian(&nums2, &nums1)
	} else {
		return FindMedian(&nums1, &nums2)
	}
}

func FindMedian(small *[]int, large *[]int) float64 {
	sLen := len(*small)
	bLen := len(*large)
	total := sLen + bLen
	half := total / 2
	if bLen == 0 {
		return 0
	}
	if sLen == 0 {
		return float64(((*small)[(bLen-1)/2] + (*small)[bLen/2]) / 2)
	}
	left, right := 0, sLen-1
	small_idx, large_idx := 0, 0
	small_left, small_right := math.MinInt32, math.MaxInt32
	large_left, large_right := math.MinInt32, math.MaxInt32
	for left <= right+1 {
		if left+right >= 0 { // not shift left most
			small_idx = (left + right) / 2
		} else {
			small_idx = -1
		}
		large_idx = half - small_idx - 2

		if small_idx >= 0 {
			small_left = (*small)[small_idx]
		} else {
			small_left = math.MinInt32
		}
		if small_idx+1 < sLen {
			small_right = (*small)[small_idx+1]
		} else {
			small_right = math.MaxInt32
		}
		if large_idx >= 0 {
			large_left = (*large)[large_idx]
		} else {
			large_left = math.MinInt32
		}
		if large_idx+1 < bLen {
			large_right = (*large)[large_idx+1]
		} else {
			large_right = math.MaxInt32
		}
		if small_right >= large_left && large_right >= small_left {
			if total%2 == 1 {
				return float64(Min(large_right, small_right))
			} else {
				return float64(Min(large_right, small_right)+Max(large_left, small_left)) / 2
			}
		}
		if small_left > large_right {
			right = small_idx - 1
		} else {
			left = small_idx + 1
		}
	}
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
