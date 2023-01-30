package leetcode12

import (
	"testing"
)

func TestShouldReturn100And80And2(t *testing.T) {
	k, h, d, u := decompose(182)
	if k != 0 {
		t.Errorf("got %d, want %d", k, 0)
	}

	if h != 100 {
		t.Errorf("got %d, want %d", h, 100)
	}

	if d != 80 {
		t.Errorf("got %d, want %d", d, 80)
	}

	if u != 2 {
		t.Errorf("got %d, want %d", u, 2)
	}
}

func TestShouldReturn1000And900And90(t *testing.T) {
	k, h, d, u := decompose(1990)
	if k != 1000 {
		t.Errorf("got %d, want %d", k, 1000)
	}

	if h != 900 {
		t.Errorf("got %d, want %d", h, 900)
	}

	if d != 90 {
		t.Errorf("got %d, want %d", d, 90)
	}

	if u != 0 {
		t.Errorf("got %d, want %d", u, 0)
	}
}
func TestCase1(t *testing.T) {
	roman := Solution(182)
	if roman != "CLXXXII" {
		t.Errorf("got %s, want %s", roman, "CLXXXII")
	}
}

func TestCase2(t *testing.T) {
	roman := Solution(1990)
	if roman != "MCMXC" {
		t.Errorf("got %s, want %s", roman, "MCMXC")
	}
}

func TestCase3(t *testing.T) {
	roman := Solution(1875)
	if roman != "MDCCCLXXV" {
		t.Errorf("got %s, want %s", roman, "MDCCCLXXV")
	}
}

func TestCase4(t *testing.T) {
	roman := Solution(3)
	if roman != "III" {
		t.Errorf("got %s, want %s", roman, "III")
	}
}

func TestCase5(t *testing.T) {
	roman := Solution(4)
	if roman != "IV" {
		t.Errorf("got %s, want %s", roman, "IV")
	}
}

func TestCase6(t *testing.T) {
	roman := Solution(9)
	if roman != "IX" {
		t.Errorf("got %s, want %s", roman, "IX")
	}
}

func TestCase7(t *testing.T) {
	roman := Solution(58)
	if roman != "LVIII" {
		t.Errorf("got %s, want %s", roman, "LVIII")
	}
}

func TestCase8(t *testing.T) {
	t.Skip()
	roman := Solution(1994)
	if roman != "MCMXCIV" {
		t.Errorf("got %s, want %s", roman, "MCMXCIV")
	}
}
