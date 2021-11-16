package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		actual := sum(numbers)
		if expected != actual {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})
	t.Run("collection of any items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		expected := 21
		actual := sum(numbers)
		if expected != actual {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})
}

func TestNumbersToSum(t *testing.T) {
	expected := []int{1,3}
	actual := numbersToSum([]int{1}, []int{1,2})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestNumbersToSumTail(t *testing.T) {
	expected := []int{1,2}
	actual := numbersToSumTail([]int{0, 1}, []int{1,2})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
