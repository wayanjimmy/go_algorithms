package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	if FizzBuz(15) != "FizzBuzz" {
		t.Error("Expected FizzBuzz")
	}

	if FizzBuz(3) != "Fizz" {
		t.Error("Expected Fizz")
	}

	if FizzBuz(5) != "Buzz" {
		t.Error("Expected Buzz")
	}

	if FizzBuz(7) != "7" {
		t.Error("Expected 7")
	}
}
