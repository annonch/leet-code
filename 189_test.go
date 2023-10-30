package ll189

import (
	"testing"
)

func TestMajorityElement1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	wantNums := []int{5, 6, 7, 1, 2, 3, 4}
	val := 3
	gotNums, err := Rotate(nums, val)

	if !equal(wantNums, gotNums) || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotNums, wantNums, err)
	}
}

func TestMajorityElement2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	wantNums := []int{3, 99, -1, -100}
	val := 2
	gotNums, err := Rotate(nums, val)

	if !equal(wantNums, gotNums) || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotNums, wantNums, err)
	}
}

func TestMajorityElement3(t *testing.T) {
	nums := []int{1}
	wantNums := []int{1}
	val := 0
	gotNums, err := Rotate(nums, val)

	if !equal(wantNums, gotNums) || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotNums, wantNums, err)
	}
}

func TestMajorityElement4(t *testing.T) {
	nums := []int{1, 2}
	wantNums := []int{2, 1}
	val := 3
	gotNums, err := Rotate(nums, val)

	if !equal(wantNums, gotNums) || err != nil {
		t.Fatalf("you fucked it up this time, got %v, wanted %v: Error: %q", gotNums, wantNums, err)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
