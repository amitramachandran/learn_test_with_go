package arrays

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestSumOfArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	sum := SumOfArray(arr)
	expected := 10

	if sum != expected {
		t.Errorf("expected sum was %d but got %d", expected, sum)
	}
}

func TestSumOfSlice(t *testing.T) {
	sli := []int{1, 2, 3, 4}
	sum := SumOfSlice(sli)
	expected := 10

	if sum != expected {
		t.Errorf("expected sum was %d but got %d", expected, sum)
	}
}

func TestSumAll(t *testing.T) {
	sli1 := []int{1, 2, 3, 4}
	sli2 := []int{1, 2, 6}
	got := SumAll(sli1, sli2)
	expected := []int{10, 9}

	// I have used slices instead of reflect.DeepEqual because its type safe
	if !slices.Equal(got, expected) {
		t.Errorf("expected sum was %d but got %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	sli1 := []int{1, 2, 3, 4}
	sli2 := []int{1, 2, 6}
	got := SumAllTails(sli1, sli2)
	expected := []int{9, 8}

	if !slices.Equal(got, expected) {
		t.Errorf("expected sum was %d but got %d", expected, got)
	}

}
