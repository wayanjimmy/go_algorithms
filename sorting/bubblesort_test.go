package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{89, 34, 45, 3, 2, 34, 90, 88, 17, 20}
	fmt.Println("Unsorted array: ", a)
	BubbleSort(a)
	fmt.Println("Sorted array: ", a)

	if a[0] != 2 {
		t.Error("Expected 2")
	}

	if a[len(a)-1] != 90 {
		t.Error("Expected 90")
	}
}

func TestSwap(t *testing.T) {
	a := []int{1, 2}
	swap(a, 0, 1)

	if a[0] != 2 {
		t.Error("Expected 2")
	}

	if a[1] != 1 {
		t.Error("Expected 1")
	}
}
