package leetcode239

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	if !reflect.DeepEqual(maxes, []int{3, 3, 5, 5, 6, 7}) {
		t.Errorf("got %v, want %v", maxes, []int{3, 3, 5, 5, 6, 7})
	}
}

func TestCase2(t *testing.T) {
	maxes := maxSlidingWindow([]int{1}, 1)
	if !reflect.DeepEqual(maxes, []int{1}) {
		t.Errorf("got %v, want %v", maxes, []int{1})
	}
}

func TestCase3(t *testing.T) {
	maxes := maxSlidingWindow([]int{-4, 2, -5, 3, 6}, 3)
	if !reflect.DeepEqual(maxes, []int{2, 3, 6}) {
		t.Errorf("got %v, want %v", maxes, []int{2, 3, 6})
	}
}
func TestCase4(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	if !reflect.DeepEqual(maxes, []int{3, 3, 5, 5, 6, 7}) {
		t.Errorf("got %v, want %v", maxes, []int{3, 3, 5, 5, 6, 7})
	}
}
func TestCase5(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)
	if !reflect.DeepEqual(maxes, []int{4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("got %v, want %v", maxes, []int{4, 5, 6, 7, 8, 9, 10})
	}
}

func TestCase6(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, 2, 3, 4, 5, 6}, 6)
	if !reflect.DeepEqual(maxes, []int{6}) {
		t.Errorf("got %v, want %v", maxes, []int{6})
	}
}

func TestCase7(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, -1}, 1)
	if !reflect.DeepEqual(maxes, []int{1, -1}) {
		t.Errorf("got %v, want %v", maxes, []int{1, -1})
	}
}

func TestCase8(t *testing.T) {
	maxes := maxSlidingWindow([]int{7, 2, 4}, 2)
	if !reflect.DeepEqual(maxes, []int{7, 4}) {
		t.Errorf("got %v, want %v", maxes, []int{7, 4})
	}
}
func TestCase9(t *testing.T) {
	maxes := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	if !reflect.DeepEqual(maxes, []int{3, 3, 2, 5}) {
		t.Errorf("got %v, want %v", maxes, []int{3, 3, 2, 5})
	}
}

func TestCase10(t *testing.T) {
	maxes := maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5)
	if !reflect.DeepEqual(maxes, []int{10, 10, 9, 2}) {
		t.Errorf("got %v, want %v", maxes, []int{10, 10, 9, 2})
	}
}
