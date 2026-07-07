package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("slice of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t * testing.T) {
	t.Run("2 slices", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func (t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("A slice with only one element", func(t *testing.T) {
		got := SumAllTails([]int{1})
		want := []int{0}

		checkSums(t, got, want)
	})
	
	t.Run("A slice with no elements", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Two slices with no elements", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		checkSums(t, got, want)
	})
}
