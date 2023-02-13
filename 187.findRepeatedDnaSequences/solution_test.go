package leetcode187

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", 10)
	answer1 := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	answer2 := []string{"CCCCCAAAAA", "AAAAACCCCC"}
	if !reflect.DeepEqual(repeatedDnas, answer1) && !reflect.DeepEqual(repeatedDnas, answer2) {
		t.Errorf("got %v, want %v or %v", repeatedDnas, answer1, answer2)
	}
}

func TestCase2(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AAAAAAAAAAAAA", 10)
	answer1 := []string{"AAAAAAAAAA"}
	if !reflect.DeepEqual(repeatedDnas, answer1) {
		t.Errorf("got %v, want %v", repeatedDnas, answer1)
	}
}

func TestCase3(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("CCATATGTATGGATAT", 4)
	answer1 := []string{"ATAT", "TATG"}
	answer2 := []string{"TATG", "ATAT"}
	if !reflect.DeepEqual(repeatedDnas, answer1) && !reflect.DeepEqual(repeatedDnas, answer2) {
		t.Errorf("got %v, want %v or %v", repeatedDnas, answer1, answer2)
	}
}

func TestCase4(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AAAACCCCTAAAACCCC", 8)
	answer1 := []string{"AAAACCCC"}
	if !reflect.DeepEqual(repeatedDnas, answer1) {
		t.Errorf("got %v, want %v", repeatedDnas, answer1)
	}
}

func TestCase5(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AGCTGAAAGCTTAGCTG", 5)
	answer1 := []string{"AGCTG"}
	if !reflect.DeepEqual(repeatedDnas, answer1) {
		t.Errorf("got %v, want %v", repeatedDnas, answer1)
	}
}

func TestCase6(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AGCTGAAAGCTTAGCTG", 5)
	answer1 := []string{"AGCTG"}
	if !reflect.DeepEqual(repeatedDnas, answer1) {
		t.Errorf("got %v, want %v", repeatedDnas, answer1)
	}
}

func TestCase7(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("AGAGCTCCAGAGCTTG", 6)
	answer1 := []string{"AGAGCT"}
	if !reflect.DeepEqual(repeatedDnas, answer1) {
		t.Errorf("got %v, want %v", repeatedDnas, answer1)
	}
}

func TestCase8(t *testing.T) {
	repeatedDnas := findRepeatedDnaSequences("ATATATATATATATAT", 6)
	answer1 := []string{"TATATA", "ATATAT"}
	answer2 := []string{"ATATAT", "TATATA"}
	if !reflect.DeepEqual(repeatedDnas, answer1) && !reflect.DeepEqual(repeatedDnas, answer2) {
		t.Errorf("got %v, want %v or %v", repeatedDnas, answer1, answer2)
	}
}
