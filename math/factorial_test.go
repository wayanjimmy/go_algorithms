package math

import "testing"

func TestFactorialIterative(t *testing.T) {
	if FactorialIterative(1) != 1 {
		t.Error("Expected 1")
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

func TestFactorialRecursive(t *testing.T) {
	if FactorialRecursive(1) != 1 {
		t.Error("Expected 1")
	}
	if FactorialRecursive(2) != 2 {
		t.Error("Expected 2")
	}
	if FactorialRecursive(3) != 6 {
		t.Error("Expected 6")
	}
	if FactorialRecursive(4) != 24 {
		t.Error("Expected 24")
	}
	if FactorialRecursive(10) != 3628800 {
		t.Error("Expected 3628800")
	}
}
