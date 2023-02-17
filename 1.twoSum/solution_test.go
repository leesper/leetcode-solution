package leetcode1

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want %v", result, []int{0, 1})
	}
}

func TestCase2(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("got %v, want %v", result, []int{1, 2})
	}
}

func TestCase3(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want %v", result, []int{0, 1})
	}
}

func TestCase4(t *testing.T) {
	result := twoSum([]int{-1, -2, -3, -4, -5}, -8)
	if !reflect.DeepEqual(result, []int{2, 4}) {
		t.Errorf("got %v, want %v", result, []int{2, 4})
	}
}

func TestCase5(t *testing.T) {
	result := twoSum([]int{3, 2, 3}, 6)
	if !reflect.DeepEqual(result, []int{0, 2}) {
		t.Errorf("got %v, want %v", result, []int{0, 2})
	}
}
