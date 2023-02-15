package leetcode209

import "testing"

func TestCase1(t *testing.T) {
	result := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	if result != 2 {
		t.Errorf("got %d, want %d", result, 2)
	}
}

func TestCase2(t *testing.T) {
	result := minSubArrayLen(4, []int{1, 4, 4})
	if result != 1 {
		t.Errorf("got %d, want %d", result, 1)
	}
}

func TestCase3(t *testing.T) {
	result := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase4(t *testing.T) {
	result := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	if result != 2 {
		t.Errorf("got %d, want %d", result, 2)
	}
}

func TestCase5(t *testing.T) {
	result := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 3})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase6(t *testing.T) {
	result := minSubArrayLen(9, []int{1, 2, 7, 1, 8})
	if result != 2 {
		t.Errorf("got %d, want %d", result, 2)
	}
}

func TestCase7(t *testing.T) {
	result := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	if result != 3 {
		t.Errorf("got %d, want %d", result, 3)
	}
}

func TestCase8(t *testing.T) {
	result := minSubArrayLen(6, []int{7, 2, 4, 6, 5, 8})
	if result != 1 {
		t.Errorf("got %d, want %d", result, 1)
	}
}
