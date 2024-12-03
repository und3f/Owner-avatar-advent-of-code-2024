package fwk

import "golang.org/x/exp/constraints"

func FindValueInSortedSlice[V constraints.Ordered](arr []V, value V) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == value {
			return mid
		}

		if arr[mid] >= arr[low] {
			if value >= arr[low] && value < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if value > arr[mid] && value <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
