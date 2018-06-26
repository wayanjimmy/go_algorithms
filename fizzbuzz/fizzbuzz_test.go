package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(15) != "FizzBuzz" {
		t.Error("Expected FizzBuzz")
	}

	if FizzBuzz(3) != "Fizz" {
		t.Error("Expected Fizz")
	}

	if FizzBuzz(5) != "Buzz" {
		t.Error("Expected Buzz")
	}

	if FizzBuzz(7) != "7" {
		t.Error("Expected 7")
	}
}
