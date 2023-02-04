package leetcode76

import (
	"testing"
)

func TestCase1(t *testing.T) {
	minCovered := minWindow("ADOBECODEBANC", "ABC")
	if minCovered != "BANC" {
		t.Errorf("got %s, want %s", minCovered, "BANC")
	}
}

func TestCase2(t *testing.T) {
	minCovered := minWindow("a", "a")
	if minCovered != "a" {
		t.Errorf("got %s, want %s", minCovered, "a")
	}
}

func TestCase3(t *testing.T) {
	minCovered := minWindow("a", "aa")
	if minCovered != "" {
		t.Errorf("got %s, want %s", minCovered, "")
	}
}

func TestCase4(t *testing.T) {
	minCovered := minWindow("acbbaca", "aba")
	if minCovered != "baca" {
		t.Errorf("got %s, want %s", minCovered, "baca")
	}
}

func TestCase5(t *testing.T) {
	minCovered := minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd")
	if minCovered != "abbbbbcdd" {
		t.Errorf("got %s, want %s", minCovered, "abbbbbcdd")
	}
}

func TestCase6(t *testing.T) {
	minCovered := minWindow("a", "b")
	if minCovered != "" {
		t.Errorf("got %s, want %s", minCovered, "")
	}
}

func TestCase7(t *testing.T) {
	minCovered := minWindow("ab", "b")
	if minCovered != "b" {
		t.Errorf("got %s, want %s", minCovered, "b")
	}
}
