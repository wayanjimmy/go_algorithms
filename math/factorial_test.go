package math

import "testing"

func TestFactorialIterative(t *testing.T) {
	if FactorialIterative(1) != 1 {
		t.Error("Expected 6")
	}

	if FactorialIterative(2) != 2 {
		t.Error("Expected 2")
	}

	if FactorialIterative(3) != 6 {
		t.Error("Expected 6")
	}

	if FactorialIterative(4) != 24 {
		t.Error("Expected 24")
	}

	if FactorialIterative(10) != 3628800 {
		t.Error("Expected 3628800")
	}
}
