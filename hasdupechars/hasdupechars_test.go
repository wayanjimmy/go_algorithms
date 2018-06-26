package hasdupechars

import "testing"

func TestHasDupeChars(t *testing.T) {
	if HasDupeChars("aaaaaaaaaaaaaa") == false {
		t.Error("Expected true")
	}
	if HasDupeChars("aaaaaaaabaaaaa") == false {
		t.Error("Expected true")
	}
	if HasDupeChars("abc") == true {
		t.Error("Expected false")
	}
	if HasDupeChars("def") == true {
		t.Error("Expected false")
	}
}
