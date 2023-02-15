package leetcode121

import "testing"

func TestCase1(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if result != 5 {
		t.Errorf("got %d, want %d", result, 5)
	}
}

func TestCase2(t *testing.T) {
	result := maxProfit([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase3(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if result != 5 {
		t.Errorf("got %d, want %d", result, 5)
	}
}

func TestCase4(t *testing.T) {
	result := maxProfit([]int{10, 8, 6, 4, 2})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase5(t *testing.T) {
	result := maxProfit([]int{10, 4, 11, 1, 5})
	if result != 7 {
		t.Errorf("got %d, want %d", result, 7)
	}
}

func TestCase6(t *testing.T) {
	result := maxProfit([]int{7, 7, 6, 6, 6})
	if result != 0 {
		t.Errorf("got %d, want %d", result, 0)
	}
}

func TestCase7(t *testing.T) {
	result := maxProfit([]int{4, 10, 5, 1, 6, 7})
	if result != 6 {
		t.Errorf("got %d, want %d", result, 6)
	}
}

func TestCase8(t *testing.T) {
	result := maxProfit([]int{4, 4, 4, 3, 3, 4})
	if result != 1 {
		t.Errorf("got %d, want %d", result, 1)
	}
}
