package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var (
		first, second int
	)

	for {
		if second == len(nums2) {
			merged = append(merged, nums1[first:]...)
			break
		}

		if first == len(nums1) {
			merged = append(merged, nums2[second:]...)
			break
		}

		if nums1[first] > nums2[second] {
			merged = append(merged, nums2[second])
			second++

			continue
		}

		if nums1[first] < nums2[second] {
			merged = append(merged, nums1[first])
			first++

			continue
		}

		if nums1[first] == nums2[second] {
			merged = append(merged, nums1[first])
			merged = append(merged, nums2[second])

			first++
			second++
		}
	}

	if len(merged)%2 == 0 {
		return float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2
	}

	return float64(merged[len(merged)/2])
}
