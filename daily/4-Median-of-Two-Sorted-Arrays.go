package daily

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	if len(A) > len(B) {
		A, B = B, A
	}

	m, n := len(A), len(B)
	half := (m + n) / 2
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := half - i

		Aleft := math.MinInt
		if i > 0 {
			Aleft = A[i-1]
		}
		Aright := math.MaxInt
		if i < m {
			Aright = A[i]
		}
		Bleft := math.MinInt
		if j > 0 {
			Bleft = B[j-1]
		}
		Bright := math.MaxInt
		if j < n {
			Bright = B[j]
		}

		if Aleft <= Bright && Bleft <= Aright {
			if (m+n)%2 != 0 {
				return float64(min(Aright, Bright))
			}
			return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2.0
		} else if Aleft > Bright {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return 0.0
}
