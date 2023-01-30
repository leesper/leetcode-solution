package leetcode4

import "testing"

func TestCase1(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 3}, []int{2})
	if median != 2.0 {
		t.Errorf("got %.5f, want %.5f\n", median, 2.0)
	}
}

func TestCase2(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if median != 2.5 {
		t.Errorf("got %.5f, want %.5f\n", median, 2.5)
	}
}
