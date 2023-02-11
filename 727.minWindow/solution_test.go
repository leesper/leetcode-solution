package leetcode727

import "testing"

func TestCase1(t *testing.T) {
	result := minWindow("abcdebdde", "bde")
	if result != "bcde" {
		t.Errorf("got %s, want %s", result, "bcde")
	}
}

func TestCase2(t *testing.T) {
	result := minWindow("abcdebdde", "bdf")
	if result != "" {
		t.Errorf("got %s, want %s", result, "")
	}
}

func TestCase3(t *testing.T) {
	result := minWindow("this is a test string", "tis")
	if result != "this" {
		t.Errorf("got %s, want %s", result, "this")
	}
}

func TestCase4(t *testing.T) {
	result := minWindow("asbfgedasfbdaaf", "bfd")
	if result != "bfged" {
		t.Errorf("got %s, want %s", result, "bfged")
	}
}

func TestCase5(t *testing.T) {
	result := minWindow("asbfgedasfbdaaf", "bfd")
	if result != "bfged" {
		t.Errorf("got %s, want %s", result, "bfged")
	}
}

func TestCase6(t *testing.T) {
	result := minWindow("Hello how are you", "ok")
	if result != "" {
		t.Errorf("got %s, want %s", result, "")
	}
}

func TestCase7(t *testing.T) {
	result := minWindow("fgrqsqsnodwmxzkzxwqegkndaa", "kzed")
	if result != "kzxwqegknd" {
		t.Errorf("got %s, want %s", result, "kzxwqegknd")
	}
}

func TestCase8(t *testing.T) {
	result := minWindow("cnhczmccqouqadqtmjjzl", "cm")
	if result != "czm" {
		t.Errorf("got %s, want %s", result, "czm")
	}
}

func TestCase9(t *testing.T) {
	result := minWindow("cnhczmccqouqadqtmjjzl", "mm")
	if result != "mccqouqadqtm" {
		t.Errorf("got %s, want %s", result, "mccqouqadqtm")
	}
}

func TestCase10(t *testing.T) {
	result := minWindow("ffynmlzesdshlvugsigobutgaetsnjlizvqjdpccdylclqcbghhixpjihximvhapymfkjxyyxfwvsfyctmhwmfjyjidnfryiyajmtakisaxwglwpqaxaicuprrvxybzdxunypzofhpclqiybgniqzsdeqwrdsfjyfkgmejxfqjkmukvgygafwokeoeglanevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaphktonqwwanapouqyjdbptqfowhemsnsl", "ntimcimzah")
	if result != "nevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaph" {
		t.Errorf("got %s, want %s", result, "nevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaph")
	}
}

func TestCase11(t *testing.T) {
	result := minWindow("cnhczmccqouqadqtcmjjzl", "cm")
	if result != "cm" {
		t.Errorf("got %s, want %s", result, "cm")
	}
}
